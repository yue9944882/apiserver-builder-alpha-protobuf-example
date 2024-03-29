// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DeepOnes returns a DeepOneInformer.
	DeepOnes() DeepOneInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DeepOnes returns a DeepOneInformer.
func (v *version) DeepOnes() DeepOneInformer {
	return &deepOneInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
