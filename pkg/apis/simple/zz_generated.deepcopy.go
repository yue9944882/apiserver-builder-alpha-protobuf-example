// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package simple

import (
	common "github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple/common"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeepOne) DeepCopyInto(out *DeepOne) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeepOne.
func (in *DeepOne) DeepCopy() *DeepOne {
	if in == nil {
		return nil
	}
	out := new(DeepOne)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeepOne) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeepOneList) DeepCopyInto(out *DeepOneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeepOne, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeepOneList.
func (in *DeepOneList) DeepCopy() *DeepOneList {
	if in == nil {
		return nil
	}
	out := new(DeepOneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeepOneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeepOneSpec) DeepCopyInto(out *DeepOneSpec) {
	*out = *in
	out.Sample = in.Sample
	if in.SamplePointer != nil {
		in, out := &in.SamplePointer, &out.SamplePointer
		*out = new(SamplePointerElem)
		(*in).DeepCopyInto(*out)
	}
	if in.SampleList != nil {
		in, out := &in.SampleList, &out.SampleList
		*out = make([]SampleListElem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SamplePointerList != nil {
		in, out := &in.SamplePointerList, &out.SamplePointerList
		*out = make([]*SampleListPointerElem, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SampleListPointerElem)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SampleMap != nil {
		in, out := &in.SampleMap, &out.SampleMap
		*out = make(map[string]SampleMapElem, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SamplePointerMap != nil {
		in, out := &in.SamplePointerMap, &out.SamplePointerMap
		*out = make(map[string]*SampleMapPointerElem, len(*in))
		for key, val := range *in {
			var outVal *SampleMapPointerElem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SampleMapPointerElem)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ConstPtr != nil {
		in, out := &in.ConstPtr, &out.ConstPtr
		*out = new(common.CustomType)
		**out = **in
	}
	if in.ConstSlice != nil {
		in, out := &in.ConstSlice, &out.ConstSlice
		*out = make([]common.CustomType, len(*in))
		copy(*out, *in)
	}
	if in.ConstMap != nil {
		in, out := &in.ConstMap, &out.ConstMap
		*out = make(map[string]common.CustomType, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeepOneSpec.
func (in *DeepOneSpec) DeepCopy() *DeepOneSpec {
	if in == nil {
		return nil
	}
	out := new(DeepOneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeepOneStatus) DeepCopyInto(out *DeepOneStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeepOneStatus.
func (in *DeepOneStatus) DeepCopy() *DeepOneStatus {
	if in == nil {
		return nil
	}
	out := new(DeepOneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleElem) DeepCopyInto(out *SampleElem) {
	*out = *in
	out.Sub = in.Sub
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleElem.
func (in *SampleElem) DeepCopy() *SampleElem {
	if in == nil {
		return nil
	}
	out := new(SampleElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleListElem) DeepCopyInto(out *SampleListElem) {
	*out = *in
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = make([]SampleListSubElem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleListElem.
func (in *SampleListElem) DeepCopy() *SampleListElem {
	if in == nil {
		return nil
	}
	out := new(SampleListElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleListPointerElem) DeepCopyInto(out *SampleListPointerElem) {
	*out = *in
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = make([]*SampleListPointerSubElem, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SampleListPointerSubElem)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleListPointerElem.
func (in *SampleListPointerElem) DeepCopy() *SampleListPointerElem {
	if in == nil {
		return nil
	}
	out := new(SampleListPointerElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleListPointerSubElem) DeepCopyInto(out *SampleListPointerSubElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleListPointerSubElem.
func (in *SampleListPointerSubElem) DeepCopy() *SampleListPointerSubElem {
	if in == nil {
		return nil
	}
	out := new(SampleListPointerSubElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleListSubElem) DeepCopyInto(out *SampleListSubElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleListSubElem.
func (in *SampleListSubElem) DeepCopy() *SampleListSubElem {
	if in == nil {
		return nil
	}
	out := new(SampleListSubElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleMapElem) DeepCopyInto(out *SampleMapElem) {
	*out = *in
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = make(map[string]SampleMapSubElem, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleMapElem.
func (in *SampleMapElem) DeepCopy() *SampleMapElem {
	if in == nil {
		return nil
	}
	out := new(SampleMapElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleMapPointerElem) DeepCopyInto(out *SampleMapPointerElem) {
	*out = *in
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = make(map[string]*SampleMapPointerSubElem, len(*in))
		for key, val := range *in {
			var outVal *SampleMapPointerSubElem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SampleMapPointerSubElem)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleMapPointerElem.
func (in *SampleMapPointerElem) DeepCopy() *SampleMapPointerElem {
	if in == nil {
		return nil
	}
	out := new(SampleMapPointerElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleMapPointerSubElem) DeepCopyInto(out *SampleMapPointerSubElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleMapPointerSubElem.
func (in *SampleMapPointerSubElem) DeepCopy() *SampleMapPointerSubElem {
	if in == nil {
		return nil
	}
	out := new(SampleMapPointerSubElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleMapSubElem) DeepCopyInto(out *SampleMapSubElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleMapSubElem.
func (in *SampleMapSubElem) DeepCopy() *SampleMapSubElem {
	if in == nil {
		return nil
	}
	out := new(SampleMapSubElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamplePointerElem) DeepCopyInto(out *SamplePointerElem) {
	*out = *in
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = new(SamplePointerSubElem)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamplePointerElem.
func (in *SamplePointerElem) DeepCopy() *SamplePointerElem {
	if in == nil {
		return nil
	}
	out := new(SamplePointerElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamplePointerSubElem) DeepCopyInto(out *SamplePointerSubElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamplePointerSubElem.
func (in *SamplePointerSubElem) DeepCopy() *SamplePointerSubElem {
	if in == nil {
		return nil
	}
	out := new(SamplePointerSubElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleSubElem) DeepCopyInto(out *SampleSubElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleSubElem.
func (in *SampleSubElem) DeepCopy() *SampleSubElem {
	if in == nil {
		return nil
	}
	out := new(SampleSubElem)
	in.DeepCopyInto(out)
	return out
}
