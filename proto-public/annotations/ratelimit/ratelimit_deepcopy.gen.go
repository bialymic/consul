// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package ratelimit

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Spec within kubernetes types, where deepcopy-gen is used.
func (in *Spec) DeepCopyInto(out *Spec) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec. Required by controller-gen.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Spec. Required by controller-gen.
func (in *Spec) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}