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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package apis

import (
	"github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple"
	_ "github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple/install"
	simplev1 "github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		simplev1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetSimpleAPIBuilder(),
	}
}

var simpleApiGroup = builders.NewApiGroupBuilder(
	"simple.io.example",
	"github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple").
	WithUnVersionedApi(simple.ApiVersion).
	WithVersionedApis(
		simplev1.ApiVersion,
	).
	WithRootScopedKinds()

func GetSimpleAPIBuilder() *builders.APIGroupBuilder {
	return simpleApiGroup
}
