// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NOPaxosProtocol) DeepCopyInto(out *NOPaxosProtocol) {
	*out = *in
	out.Sequencer = in.Sequencer
	out.Protocol = in.Protocol
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NOPaxosProtocol.
func (in *NOPaxosProtocol) DeepCopy() *NOPaxosProtocol {
	if in == nil {
		return nil
	}
	out := new(NOPaxosProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NOPaxosProtocolSpec) DeepCopyInto(out *NOPaxosProtocolSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NOPaxosProtocolSpec.
func (in *NOPaxosProtocolSpec) DeepCopy() *NOPaxosProtocolSpec {
	if in == nil {
		return nil
	}
	out := new(NOPaxosProtocolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NOPaxosSequencerSpec) DeepCopyInto(out *NOPaxosSequencerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NOPaxosSequencerSpec.
func (in *NOPaxosSequencerSpec) DeepCopy() *NOPaxosSequencerSpec {
	if in == nil {
		return nil
	}
	out := new(NOPaxosSequencerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Partition) DeepCopyInto(out *Partition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Partition.
func (in *Partition) DeepCopy() *Partition {
	if in == nil {
		return nil
	}
	out := new(Partition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Partition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionList) DeepCopyInto(out *PartitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Partition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionList.
func (in *PartitionList) DeepCopy() *PartitionList {
	if in == nil {
		return nil
	}
	out := new(PartitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSet) DeepCopyInto(out *PartitionSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSet.
func (in *PartitionSet) DeepCopy() *PartitionSet {
	if in == nil {
		return nil
	}
	out := new(PartitionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartitionSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSetList) DeepCopyInto(out *PartitionSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PartitionSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSetList.
func (in *PartitionSetList) DeepCopy() *PartitionSetList {
	if in == nil {
		return nil
	}
	out := new(PartitionSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartitionSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSetSpec) DeepCopyInto(out *PartitionSetSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSetSpec.
func (in *PartitionSetSpec) DeepCopy() *PartitionSetSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSetStatus) DeepCopyInto(out *PartitionSetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSetStatus.
func (in *PartitionSetStatus) DeepCopy() *PartitionSetStatus {
	if in == nil {
		return nil
	}
	out := new(PartitionSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSpec) DeepCopyInto(out *PartitionSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Raft != nil {
		in, out := &in.Raft, &out.Raft
		*out = new(RaftProtocol)
		**out = **in
	}
	if in.NOPaxos != nil {
		in, out := &in.NOPaxos, &out.NOPaxos
		*out = new(NOPaxosProtocol)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSpec.
func (in *PartitionSpec) DeepCopy() *PartitionSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionStatus) DeepCopyInto(out *PartitionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionStatus.
func (in *PartitionStatus) DeepCopy() *PartitionStatus {
	if in == nil {
		return nil
	}
	out := new(PartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionTemplateSpec) DeepCopyInto(out *PartitionTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionTemplateSpec.
func (in *PartitionTemplateSpec) DeepCopy() *PartitionTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftProtocol) DeepCopyInto(out *RaftProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftProtocol.
func (in *RaftProtocol) DeepCopy() *RaftProtocol {
	if in == nil {
		return nil
	}
	out := new(RaftProtocol)
	in.DeepCopyInto(out)
	return out
}
