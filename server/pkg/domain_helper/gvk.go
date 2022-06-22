package domain_helper

import (
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/namespaces"
)

var NamespaceGvk = &namespaces.GroupVersionKind{
	Group:   "core",
	Version: "v1",
	Kind:    "Namespace",
	Enabled: true,
}

type namespaceFactory struct {
}

func (f *namespaceFactory) Create() *namespaces.Namespace {
	return &namespaces.Namespace{}
}

func NewNamespaceFactory() MetadatableFactory[*namespaces.Namespace] {
	return &namespaceFactory{}
}

var NatIpGvk = &namespaces.GroupVersionKind{
	Group:   "red-fox.sbx-central.io",
	Version: "v1alpha1",
	Kind:    "NatIp",
	Enabled: true,
}

type natIpFactory struct {
}

func (f *natIpFactory) Create() *documents.NatIp {
	return &documents.NatIp{}
}

func NewNatIpFactory() MetadatableFactory[*documents.NatIp] {
	return &natIpFactory{}
}

var EndpointGvk = &namespaces.GroupVersionKind{
	Group:   "red-fox.sbx-central.io",
	Version: "v1alpha1",
	Kind:    "Endpoint",
	Enabled: true,
}

type endpointFactory struct {
}

func (f *endpointFactory) Create() *documents.Endpoint {
	return &documents.Endpoint{}
}

func NewEndpointFactory() MetadatableFactory[*documents.Endpoint] {
	return &endpointFactory{}
}
