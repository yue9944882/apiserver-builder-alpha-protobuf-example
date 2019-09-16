// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple
// +k8s:defaulter-gen=TypeMeta
// +groupName=simple.io.example
package v1 // import "github.com/yue9944882/apiserver-builder-alpha-protobuf-example/pkg/apis/simple/v1"
