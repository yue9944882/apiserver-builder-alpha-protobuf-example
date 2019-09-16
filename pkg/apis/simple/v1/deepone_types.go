/*
Copyright 2017 The Kubernetes Authors.

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

package v1

import (
	"github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
// +resource:path=deepones
// DeepOne defines a resident of innsmouth
type DeepOne struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   DeepOneSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status DeepOneStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type SamplePrimitiveAlias int64

// DeepOnesSpec defines the desired state of DeepOne
type DeepOneSpec struct {
	// fish_required defines the number of fish required by the DeepOne.
	// NOTE: the type has to be int64 instead of ambiguous int
	FishRequired int64 `json:"fish_required,omitempty" protobuf:"varint,1,opt,name=fish_required"`

	Sample            SampleElem               `json:"sample,omitempty" protobuf:"bytes,2,opt,name=sample"`
	SamplePointer     *SamplePointerElem       `json:"sample_pointer,omitempty" protobuf:"bytes,3,opt,name=sample_pointer"`
	SampleList        []SampleListElem         `json:"sample_list,omitempty" protobuf:"bytes,4,rep,name=sample_list"`
	SamplePointerList []*SampleListPointerElem `json:"sample_pointer_list,omitempty" protobuf:"bytes,5,rep,name=sample_pointer_list"`
	SampleMap         map[string]SampleMapElem `json:"sample_map,omitempty" protobuf:"bytes,6,rep,name=sample_map"`
	// NOTE: Maps using pointer as value type is not supported protobuf serialization
	//SamplePointerMap     map[string]*SampleMapPointerElem `json:"sample_pointer_map,omitempty" protobuf:"bytes,7,rep,name=sample_pointer_map"`
	SamplePrimitiveAlias SamplePrimitiveAlias `json:"sample_primitive_alias,omitempty" protobuf:"varint,8,opt,name=sample_primitive_alias"`

	// Example of using a constant
	Const      common.CustomType            `json:"const,omitempty" protobuf:"bytes,9,opt,name=const"`
	ConstPtr   *common.CustomType           `json:"constPtr,omitempty" protobuf:"bytes,10,opt,name=constPtr"`
	ConstSlice []common.CustomType          `json:"constSlice,omitempty" protobuf:"bytes,11,rep,name=constSlice"`
	ConstMap   map[string]common.CustomType `json:"constMap,omitempty" protobuf:"bytes,12,rep,name=constMap"`

	// TODO: Fix issues with deep copy to make these work
	//ConstSlicePtr []*common.CustomType          `json:"constSlicePtr,omitempty"`
	//ConstMapPtr map[string]*common.CustomType `json:"constMapPtr,omitempty"`
}

type SampleListElem struct {
	Sub []SampleListSubElem `json:"sub,omitempty" protobuf:"bytes,1,rep,name=sub"`
}

type SampleListSubElem struct {
	Foo string `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo"`
}

type SampleListPointerElem struct {
	Sub []*SampleListPointerSubElem `json:"sub,omitempty" protobuf:"bytes,1,rep,name=sub"`
}

type SampleListPointerSubElem struct {
	Foo string `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo"`
}

type SampleMapElem struct {
	Sub map[string]SampleMapSubElem `json:"sub,omitempty" protobuf:"bytes,1,rep,name=sub"`
}

type SampleMapSubElem struct {
	Foo string `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo"`
}

type SampleMapPointerElem struct {
	// NOTE: Maps using pointer as value type is not supported protobuf serialization
	//Sub map[string]*SampleMapPointerSubElem `json:"sub,omitempty" protobuf:"bytes,1,rep,name=sub"`
}

type SampleMapPointerSubElem struct {
	Foo string `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo"`
}

type SamplePointerElem struct {
	Sub *SamplePointerSubElem `json:"sub,omitempty" protobuf:"bytes,1,rep,name=sub"`
}

type SamplePointerSubElem struct {
	Foo string `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo"`
}

type SampleElem struct {
	Sub SampleSubElem `json:"sub,omitempty" protobuf:"bytes,1,rep,name=sub"`
}

type SampleSubElem struct {
	Foo string `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo"`
}

// DeepOneStatus defines the observed state of DeepOne
type DeepOneStatus struct {
	// actual_fish defines the number of fish caught by the DeepOne.
	// NOTE: the type has to be int64 instead of ambiguous int
	ActualFish int64 `json:"actual_fish,omitempty" protobuf:"varint,1,opt,name=actual_fish"`
}
