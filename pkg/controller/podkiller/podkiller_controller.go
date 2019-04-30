package podkiller

import (
	"context"
	faisalv1alpha1 "github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_podkiller")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new PodKiller Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePodKiller{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("podkiller-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PodKiller
	err = c.Watch(&source.Kind{Type: &faisalv1alpha1.PodKiller{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner PodKiller
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &faisalv1alpha1.PodKiller{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcilePodKiller{}

// ReconcilePodKiller reconciles a PodKiller object
type ReconcilePodKiller struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a PodKiller object and makes changes based on the state read
// and what is in the PodKiller.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcilePodKiller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling PodKiller")

	// Fetch the PodKiller instance
	instance := &faisalv1alpha1.PodKiller{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	allNamespaces, err := getAllNamespace(r)
	if err != nil {
		reqLogger.Error(err, "failed to list all namespaces")
		return reconcile.Result{}, err
	}

	//allowedNamespaces := make([]string, 0)

	safeNamespaceMap := getSafeNamespaces(instance)

	for _, namespace := range allNamespaces.Items {

		isInMap := safeNamespaceMap[namespace.Name] || strings.HasPrefix(namespace.Name, "openshift-")
		//log.Info("Looping namespace tp avoid list ", "NS", namespace.Name, "IsInMap", isInMap)

		if isInMap == false {
			//log.Info("This namesoace is one of the target ", "NS", namespace.Name, "IsInMap", isInMap)
			//allowedNamespaces = append(allowedNamespaces, namespace.Name)

			existingPods := &corev1.PodList{}

			//TODO move this to a go routine
			err = r.client.List(context.TODO(),
				&client.ListOptions{
					Namespace: namespace.Name,
				},
				existingPods)

			if err != nil {
				reqLogger.Error(err, "failed to list existing pods")
				return reconcile.Result{}, err
			}

			for i := range existingPods.Items {
				log.Info("=============== ", "PN", existingPods.Items[i].Name)
				go possiblyDeletePods(&existingPods.Items[i], r)
			}
		}

	}

	/*
		allNamespaces := &corev1.NamespaceList{}
		err = r.client.List(context.TODO(),
					  &client.ListOptions{},
					  allNamespaces,
		)
		if err != nil {
			reqLogger.Error(err, "failed to list all namespaces")
			return reconcile.Result{}, err
		}

		allowedNamespaces := make([]string, 0)
		safeNamespaceMap := make(map[string]bool, 0)
		for _, ele := range instance.Spec.NamespacesToAvoid {
			log.Info("Adding namespace tp avoid list ", "NS", ele)
			safeNamespaceMap[ele] = true
		}

		for _, namespace := range allNamespaces.Items{

			isInMap := safeNamespaceMap[namespace.Name] || strings.HasPrefix(namespace.Name, "openshift-")
			//log.Info("Looping namespace tp avoid list ", "NS", namespace.Name, "IsInMap", isInMap)

			if isInMap == false {
				log.Info("This namesoace is one of the target ", "NS", namespace.Name, "IsInMap", isInMap)
				allowedNamespaces = append(allowedNamespaces, namespace.Name)

				existingPods := &corev1.PodList{}

				err = r.client.List(context.TODO(),
					&client.ListOptions{
						Namespace:     namespace.Name,
					},
					existingPods)

				if err != nil {
					reqLogger.Error(err, "failed to list existing pods")
					return reconcile.Result{}, err
				}


				for _, pod := range existingPods.Items {
					if pod.Status.Phase != corev1.PodRunning {
						continue
					}

					log.Info("Working on Pod ", "NS", namespace.Name, "Pod", pod.Name)
					podStartTime := pod.Status.StartTime
					startTime := time.Unix(podStartTime.Unix(), 0) // convert to standard go time
					runningDuration := time.Now().Sub(startTime)


					//if runningDuration > 111111 {
					//
					//}


					for _, or := range pod.OwnerReferences {
						//log.Info("Pod OR 11111111", "Pod", pod)
						log.Info("Pod OR ", "Pod", pod.Name, "OR", or.Name, "ORK", or.Kind, "ORC", or.Controller)
						if *or.Controller == true {
							controllerType := or.Kind
							controllerName := or.Name

							controllerNamed := &types.NamespacedName{
								Namespace: pod.Namespace,
								Name: controllerName, }

							//make sure to add check to avoid killer-operator pods

							if controllerType == "ReplicaSet" {
								targetReplicaSet := &v1.ReplicaSet{}
								err = r.client.Get(context.TODO(), *controllerNamed, targetReplicaSet)
								reqLogger.Info("About to DELETE DEpoyment for Replicaset ", "RunDuration", runningDuration, "RSName", targetReplicaSet)
								//r.client.Delete(context.TODO(), targetReplicaSet)
							} else if controllerType == "ReplicationController" {
								targetReplicationController := &corev1.ReplicationController{}
								err = r.client.Get(context.TODO(), *controllerNamed, targetReplicationController)
								reqLogger.Info("About to DELETE DEpoyment for Replication Controller ", "RunDuration", runningDuration, "RSName", targetReplicationController)
								//r.client.Delete(context.TODO(), targetReplicaSet)
							}
						}
					}
				}
			}

		}



		findPodsToKill := &corev1.PodList{}
		err = r.client.List(context.TODO(),
					  &client.ListOptions{
					  	Namespace: request.Namespace,
					  },
					  findPodsToKill,
			)

		if err != nil {
			reqLogger.Error(err, "failed to list existing pods in the podSet")
			return reconcile.Result{}, err
		}*/

	return reconcile.Result{Requeue: true}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *faisalv1alpha1.PodKiller) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}

func getAllNamespace(r *ReconcilePodKiller) (*corev1.NamespaceList, error) {
	allNamespaces := &corev1.NamespaceList{}
	err := r.client.List(context.TODO(),
		&client.ListOptions{},
		allNamespaces,
	)
	if err != nil {
		log.Error(err, "failed to list all namespaces")
		return nil, err
	}

	return allNamespaces, nil

}

func getSafeNamespaces(instance *faisalv1alpha1.PodKiller) map[string]bool {
	safeNamespaceMap := make(map[string]bool, 0)
	for _, ele := range instance.Spec.NamespacesToAvoid {
		log.Info("Adding namespace tp avoid list ", "NS", ele)
		safeNamespaceMap[ele] = true
	}

	return safeNamespaceMap
}

func possiblyDeletePods(pod *corev1.Pod, r *ReconcilePodKiller) {
	log.Info("Working on Pod ", "NS", pod.Namespace, "Pod", pod.Name, "Statud", pod.Status.Phase)

	if pod.Status.Phase != corev1.PodRunning {
		return
	}

	podStartTime := pod.Status.StartTime
	startTime := time.Unix(podStartTime.Unix(), 0) // convert to standard go time
	runningDuration := time.Now().Sub(startTime)

	//if runningDuration > 111111 {
	//
	//}

	for _, or := range pod.OwnerReferences {
		//log.Info("Pod OR 11111111", "Pod", pod)
		log.Info("Pod OR ", "Pod", pod.Name, "OR", or.Name, "ORK", or.Kind, "ORC", or.Controller)
		if *or.Controller == true {
			controllerType := or.Kind
			controllerName := or.Name

			controllerNamed := &types.NamespacedName{
				Namespace: pod.Namespace,
				Name:      controllerName}

			//make sure to add check to avoid killer-operator pods

			if controllerType == "ReplicaSet" {
				targetReplicaSet := &v1.ReplicaSet{}
				err := r.client.Get(context.TODO(), *controllerNamed, targetReplicaSet)
				if err != nil {
					log.Error(err, "CAnnot fetch Replica Set")
				}
				log.Info("About to DELETE DEpoyment for Replicaset ", "RunDuration", runningDuration, "RSName", targetReplicaSet)
				//r.client.Delete(context.TODO(), targetReplicaSet)
			} else if controllerType == "ReplicationController" {
				targetReplicationController := &corev1.ReplicationController{}
				err := r.client.Get(context.TODO(), *controllerNamed, targetReplicationController)
				if err != nil {
					log.Error(err, "CAnnot fetch Replication Controller")
				}
				log.Info("About to DELETE DEpoyment for Replication Controller ", "RunDuration", runningDuration, "RSName", targetReplicationController)
				//r.client.Delete(context.TODO(), targetReplicaSet)
			}
		}
	}
}

/*func (r *ReconcilePodKiller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling PodKiller")

	// Fetch the PodKiller instance
	instance := &faisalv1alpha1.PodKiller{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	pod := newPodForCR(instance)

	// Set PodKiller instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}
*/
