// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKECluster) DeepCopyInto(out *GKECluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKECluster.
func (in *GKECluster) DeepCopy() *GKECluster {
	if in == nil {
		return nil
	}
	out := new(GKECluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GKECluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEClusterList) DeepCopyInto(out *GKEClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GKECluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEClusterList.
func (in *GKEClusterList) DeepCopy() *GKEClusterList {
	if in == nil {
		return nil
	}
	out := new(GKEClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GKEClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEClusterSpec) DeepCopyInto(out *GKEClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeLocations != nil {
		in, out := &in.NodeLocations, &out.NodeLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeTaints != nil {
		in, out := &in.NodeTaints, &out.NodeTaints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeVersion != nil {
		in, out := &in.NodeVersion, &out.NodeVersion
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEClusterSpec.
func (in *GKEClusterSpec) DeepCopy() *GKEClusterSpec {
	if in == nil {
		return nil
	}
	out := new(GKEClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEClusterStatus) DeepCopyInto(out *GKEClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEClusterStatus.
func (in *GKEClusterStatus) DeepCopy() *GKEClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GKEClusterStatus)
	in.DeepCopyInto(out)
	return out
}