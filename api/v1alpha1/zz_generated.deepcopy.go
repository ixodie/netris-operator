// +build !ignore_autogenerated

/*
Copyright 2021. Netris, Inc.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allocation) DeepCopyInto(out *Allocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allocation.
func (in *Allocation) DeepCopy() *Allocation {
	if in == nil {
		return nil
	}
	out := new(Allocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Allocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationList) DeepCopyInto(out *AllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Allocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationList.
func (in *AllocationList) DeepCopy() *AllocationList {
	if in == nil {
		return nil
	}
	out := new(AllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationMeta) DeepCopyInto(out *AllocationMeta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationMeta.
func (in *AllocationMeta) DeepCopy() *AllocationMeta {
	if in == nil {
		return nil
	}
	out := new(AllocationMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AllocationMeta) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationMetaList) DeepCopyInto(out *AllocationMetaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AllocationMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationMetaList.
func (in *AllocationMetaList) DeepCopy() *AllocationMetaList {
	if in == nil {
		return nil
	}
	out := new(AllocationMetaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AllocationMetaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationMetaSpec) DeepCopyInto(out *AllocationMetaSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationMetaSpec.
func (in *AllocationMetaSpec) DeepCopy() *AllocationMetaSpec {
	if in == nil {
		return nil
	}
	out := new(AllocationMetaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationMetaStatus) DeepCopyInto(out *AllocationMetaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationMetaStatus.
func (in *AllocationMetaStatus) DeepCopy() *AllocationMetaStatus {
	if in == nil {
		return nil
	}
	out := new(AllocationMetaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationSpec) DeepCopyInto(out *AllocationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationSpec.
func (in *AllocationSpec) DeepCopy() *AllocationSpec {
	if in == nil {
		return nil
	}
	out := new(AllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationStatus) DeepCopyInto(out *AllocationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationStatus.
func (in *AllocationStatus) DeepCopy() *AllocationStatus {
	if in == nil {
		return nil
	}
	out := new(AllocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGP) DeepCopyInto(out *BGP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGP.
func (in *BGP) DeepCopy() *BGP {
	if in == nil {
		return nil
	}
	out := new(BGP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPList) DeepCopyInto(out *BGPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPList.
func (in *BGPList) DeepCopy() *BGPList {
	if in == nil {
		return nil
	}
	out := new(BGPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPMeta) DeepCopyInto(out *BGPMeta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPMeta.
func (in *BGPMeta) DeepCopy() *BGPMeta {
	if in == nil {
		return nil
	}
	out := new(BGPMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPMeta) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPMetaList) DeepCopyInto(out *BGPMetaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGPMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPMetaList.
func (in *BGPMetaList) DeepCopy() *BGPMetaList {
	if in == nil {
		return nil
	}
	out := new(BGPMetaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPMetaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPMetaSpec) DeepCopyInto(out *BGPMetaSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPMetaSpec.
func (in *BGPMetaSpec) DeepCopy() *BGPMetaSpec {
	if in == nil {
		return nil
	}
	out := new(BGPMetaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPMetaStatus) DeepCopyInto(out *BGPMetaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPMetaStatus.
func (in *BGPMetaStatus) DeepCopy() *BGPMetaStatus {
	if in == nil {
		return nil
	}
	out := new(BGPMetaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPMultihop) DeepCopyInto(out *BGPMultihop) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPMultihop.
func (in *BGPMultihop) DeepCopy() *BGPMultihop {
	if in == nil {
		return nil
	}
	out := new(BGPMultihop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPSpec) DeepCopyInto(out *BGPSpec) {
	*out = *in
	out.Transport = in.Transport
	out.Multihop = in.Multihop
	if in.PrefixListInbound != nil {
		in, out := &in.PrefixListInbound, &out.PrefixListInbound
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrefixListOutbound != nil {
		in, out := &in.PrefixListOutbound, &out.PrefixListOutbound
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SendBGPCommunity != nil {
		in, out := &in.SendBGPCommunity, &out.SendBGPCommunity
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPSpec.
func (in *BGPSpec) DeepCopy() *BGPSpec {
	if in == nil {
		return nil
	}
	out := new(BGPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPStatus) DeepCopyInto(out *BGPStatus) {
	*out = *in
	in.ModifiedDate.DeepCopyInto(&out.ModifiedDate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPStatus.
func (in *BGPStatus) DeepCopy() *BGPStatus {
	if in == nil {
		return nil
	}
	out := new(BGPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPTransport) DeepCopyInto(out *BGPTransport) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPTransport.
func (in *BGPTransport) DeepCopy() *BGPTransport {
	if in == nil {
		return nil
	}
	out := new(BGPTransport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LB) DeepCopyInto(out *L4LB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LB.
func (in *L4LB) DeepCopy() *L4LB {
	if in == nil {
		return nil
	}
	out := new(L4LB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L4LB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBCheck) DeepCopyInto(out *L4LBCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBCheck.
func (in *L4LBCheck) DeepCopy() *L4LBCheck {
	if in == nil {
		return nil
	}
	out := new(L4LBCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBFrontend) DeepCopyInto(out *L4LBFrontend) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBFrontend.
func (in *L4LBFrontend) DeepCopy() *L4LBFrontend {
	if in == nil {
		return nil
	}
	out := new(L4LBFrontend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBList) DeepCopyInto(out *L4LBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]L4LB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBList.
func (in *L4LBList) DeepCopy() *L4LBList {
	if in == nil {
		return nil
	}
	out := new(L4LBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L4LBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMeta) DeepCopyInto(out *L4LBMeta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMeta.
func (in *L4LBMeta) DeepCopy() *L4LBMeta {
	if in == nil {
		return nil
	}
	out := new(L4LBMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L4LBMeta) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaBackend) DeepCopyInto(out *L4LBMetaBackend) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaBackend.
func (in *L4LBMetaBackend) DeepCopy() *L4LBMetaBackend {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaHealthCheck) DeepCopyInto(out *L4LBMetaHealthCheck) {
	*out = *in
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(L4LBMetaHealthCheckTCP)
		**out = **in
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(L4LBMetaHealthCheckHTTP)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaHealthCheck.
func (in *L4LBMetaHealthCheck) DeepCopy() *L4LBMetaHealthCheck {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaHealthCheckHTTP) DeepCopyInto(out *L4LBMetaHealthCheckHTTP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaHealthCheckHTTP.
func (in *L4LBMetaHealthCheckHTTP) DeepCopy() *L4LBMetaHealthCheckHTTP {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaHealthCheckHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaHealthCheckTCP) DeepCopyInto(out *L4LBMetaHealthCheckTCP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaHealthCheckTCP.
func (in *L4LBMetaHealthCheckTCP) DeepCopy() *L4LBMetaHealthCheckTCP {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaHealthCheckTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaList) DeepCopyInto(out *L4LBMetaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]L4LBMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaList.
func (in *L4LBMetaList) DeepCopy() *L4LBMetaList {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L4LBMetaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaSpec) DeepCopyInto(out *L4LBMetaSpec) {
	*out = *in
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(L4LBMetaHealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = make([]L4LBMetaBackend, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaSpec.
func (in *L4LBMetaSpec) DeepCopy() *L4LBMetaSpec {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBMetaStatus) DeepCopyInto(out *L4LBMetaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBMetaStatus.
func (in *L4LBMetaStatus) DeepCopy() *L4LBMetaStatus {
	if in == nil {
		return nil
	}
	out := new(L4LBMetaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBSpec) DeepCopyInto(out *L4LBSpec) {
	*out = *in
	out.Check = in.Check
	out.Frontend = in.Frontend
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = make([]L4LBBackend, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBSpec.
func (in *L4LBSpec) DeepCopy() *L4LBSpec {
	if in == nil {
		return nil
	}
	out := new(L4LBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4LBStatus) DeepCopyInto(out *L4LBStatus) {
	*out = *in
	in.ModifiedDate.DeepCopyInto(&out.ModifiedDate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4LBStatus.
func (in *L4LBStatus) DeepCopy() *L4LBStatus {
	if in == nil {
		return nil
	}
	out := new(L4LBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Site) DeepCopyInto(out *Site) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Site.
func (in *Site) DeepCopy() *Site {
	if in == nil {
		return nil
	}
	out := new(Site)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Site) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteList) DeepCopyInto(out *SiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Site, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteList.
func (in *SiteList) DeepCopy() *SiteList {
	if in == nil {
		return nil
	}
	out := new(SiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteMeta) DeepCopyInto(out *SiteMeta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteMeta.
func (in *SiteMeta) DeepCopy() *SiteMeta {
	if in == nil {
		return nil
	}
	out := new(SiteMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteMeta) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteMetaList) DeepCopyInto(out *SiteMetaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SiteMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteMetaList.
func (in *SiteMetaList) DeepCopy() *SiteMetaList {
	if in == nil {
		return nil
	}
	out := new(SiteMetaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteMetaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteMetaSpec) DeepCopyInto(out *SiteMetaSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteMetaSpec.
func (in *SiteMetaSpec) DeepCopy() *SiteMetaSpec {
	if in == nil {
		return nil
	}
	out := new(SiteMetaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteMetaStatus) DeepCopyInto(out *SiteMetaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteMetaStatus.
func (in *SiteMetaStatus) DeepCopy() *SiteMetaStatus {
	if in == nil {
		return nil
	}
	out := new(SiteMetaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteSpec) DeepCopyInto(out *SiteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteSpec.
func (in *SiteSpec) DeepCopy() *SiteSpec {
	if in == nil {
		return nil
	}
	out := new(SiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteStatus) DeepCopyInto(out *SiteStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteStatus.
func (in *SiteStatus) DeepCopy() *SiteStatus {
	if in == nil {
		return nil
	}
	out := new(SiteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNet) DeepCopyInto(out *VNet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNet.
func (in *VNet) DeepCopy() *VNet {
	if in == nil {
		return nil
	}
	out := new(VNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VNet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetList) DeepCopyInto(out *VNetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VNet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetList.
func (in *VNetList) DeepCopy() *VNetList {
	if in == nil {
		return nil
	}
	out := new(VNetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VNetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMeta) DeepCopyInto(out *VNetMeta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMeta.
func (in *VNetMeta) DeepCopy() *VNetMeta {
	if in == nil {
		return nil
	}
	out := new(VNetMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VNetMeta) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMetaGateway) DeepCopyInto(out *VNetMetaGateway) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMetaGateway.
func (in *VNetMetaGateway) DeepCopy() *VNetMetaGateway {
	if in == nil {
		return nil
	}
	out := new(VNetMetaGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMetaList) DeepCopyInto(out *VNetMetaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VNetMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMetaList.
func (in *VNetMetaList) DeepCopy() *VNetMetaList {
	if in == nil {
		return nil
	}
	out := new(VNetMetaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VNetMetaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMetaMember) DeepCopyInto(out *VNetMetaMember) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMetaMember.
func (in *VNetMetaMember) DeepCopy() *VNetMetaMember {
	if in == nil {
		return nil
	}
	out := new(VNetMetaMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMetaSite) DeepCopyInto(out *VNetMetaSite) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMetaSite.
func (in *VNetMetaSite) DeepCopy() *VNetMetaSite {
	if in == nil {
		return nil
	}
	out := new(VNetMetaSite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMetaSpec) DeepCopyInto(out *VNetMetaSpec) {
	*out = *in
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]VNetMetaGateway, len(*in))
		copy(*out, *in)
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]VNetMetaMember, len(*in))
		copy(*out, *in)
	}
	if in.Sites != nil {
		in, out := &in.Sites, &out.Sites
		*out = make([]VNetMetaSite, len(*in))
		copy(*out, *in)
	}
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMetaSpec.
func (in *VNetMetaSpec) DeepCopy() *VNetMetaSpec {
	if in == nil {
		return nil
	}
	out := new(VNetMetaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetMetaStatus) DeepCopyInto(out *VNetMetaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetMetaStatus.
func (in *VNetMetaStatus) DeepCopy() *VNetMetaStatus {
	if in == nil {
		return nil
	}
	out := new(VNetMetaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetSite) DeepCopyInto(out *VNetSite) {
	*out = *in
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]VNetGateway, len(*in))
		copy(*out, *in)
	}
	if in.SwitchPorts != nil {
		in, out := &in.SwitchPorts, &out.SwitchPorts
		*out = make([]VNetSwitchPort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetSite.
func (in *VNetSite) DeepCopy() *VNetSite {
	if in == nil {
		return nil
	}
	out := new(VNetSite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetSpec) DeepCopyInto(out *VNetSpec) {
	*out = *in
	if in.GuestTenants != nil {
		in, out := &in.GuestTenants, &out.GuestTenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Sites != nil {
		in, out := &in.Sites, &out.Sites
		*out = make([]VNetSite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetSpec.
func (in *VNetSpec) DeepCopy() *VNetSpec {
	if in == nil {
		return nil
	}
	out := new(VNetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetStatus) DeepCopyInto(out *VNetStatus) {
	*out = *in
	in.ModifiedDate.DeepCopyInto(&out.ModifiedDate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetStatus.
func (in *VNetStatus) DeepCopy() *VNetStatus {
	if in == nil {
		return nil
	}
	out := new(VNetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetSwitchPort) DeepCopyInto(out *VNetSwitchPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetSwitchPort.
func (in *VNetSwitchPort) DeepCopy() *VNetSwitchPort {
	if in == nil {
		return nil
	}
	out := new(VNetSwitchPort)
	in.DeepCopyInto(out)
	return out
}
