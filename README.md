operator-sdk-v0.7.0-x86_64-apple-darwin new killer-operator

operator-sdk-v0.7.0-x86_64-apple-darwin add api --api-version=faisal.example.com/v1alpha1 --kind=PodKiller

operator-sdk-v0.7.0-x86_64-apple-darwin add controller --api-version=faisal.example.com/v1alpha1 --kind=PodKiller

//code
operator-sdk-v0.7.0-x86_64-apple-darwin generate k8s



operator-sdk-v0.7.0-x86_64-apple-darwin build killeroperator:1.0.0

docker tag killeroperator:1.0.0 quay.io/masood_faisal/publicrepo

#docker login -u masood_faisal quay.io 

docker push quay.io/masood_faisal/publicrepo


oc4 delete deployment killer-operator

oc4 create -f deploy/operator.yaml
