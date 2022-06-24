package domain_helper

import (
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
)

var NamespaceGvk = &idl_common.GroupVersionKind{
	Group:   "core",
	Version: "v1",
	Kind:    "Namespace",
}

type namespaceFactory struct {
}

func (f *namespaceFactory) Create() *namespaces.Namespace {
	return &namespaces.Namespace{}
}

func NewNamespaceFactory() MetadatableFactory[*namespaces.Namespace] {
	return &namespaceFactory{}
}

var NatIpGvk = &idl_common.GroupVersionKind{
	Group:   "red-fox.sbx-central.io",
	Version: "v1alpha1",
	Kind:    "NatIp",
}

type natIpFactory struct {
}

func (f *natIpFactory) Create() *documents.NatIp {
	return &documents.NatIp{}
}

func NewNatIpFactory() MetadatableFactory[*documents.NatIp] {
	return &natIpFactory{}
}

var EndpointGvk = &idl_common.GroupVersionKind{
	Group:   "red-fox.sbx-central.io",
	Version: "v1alpha1",
	Kind:    "Endpoint",
}

type endpointFactory struct {
}

func (f *endpointFactory) Create() *documents.Endpoint {
	return &documents.Endpoint{}
}

func NewEndpointFactory() MetadatableFactory[*documents.Endpoint] {
	return &endpointFactory{}
}
