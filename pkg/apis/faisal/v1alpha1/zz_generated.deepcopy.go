// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodKiller) DeepCopyInto(out *PodKiller) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodKiller.
func (in *PodKiller) DeepCopy() *PodKiller {
	if in == nil {
		return nil
	}
	out := new(PodKiller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodKiller) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodKillerList) DeepCopyInto(out *PodKillerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodKiller, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodKillerList.
func (in *PodKillerList) DeepCopy() *PodKillerList {
	if in == nil {
		return nil
	}
	out := new(PodKillerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodKillerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodKillerSpec) DeepCopyInto(out *PodKillerSpec) {
	*out = *in
	if in.NamespacesToAvoid != nil {
		in, out := &in.NamespacesToAvoid, &out.NamespacesToAvoid
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodKillerSpec.
func (in *PodKillerSpec) DeepCopy() *PodKillerSpec {
	if in == nil {
		return nil
	}
	out := new(PodKillerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodKillerStatus) DeepCopyInto(out *PodKillerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodKillerStatus.
func (in *PodKillerStatus) DeepCopy() *PodKillerStatus {
	if in == nil {
		return nil
	}
	out := new(PodKillerStatus)
	in.DeepCopyInto(out)
	return out
}
