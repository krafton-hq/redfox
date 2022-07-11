package domain_helper

import (
	"github.com/krafton-hq/red-fox/apis/crds"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"google.golang.org/protobuf/proto"
)

var NamespaceGvk = &idl_common.GroupVersionKindSpec{
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

var NatIpGvk = &idl_common.GroupVersionKindSpec{
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

var EndpointGvk = &idl_common.GroupVersionKindSpec{
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

var CrdGvk = &idl_common.GroupVersionKindSpec{
	Group:   "red-fox.sbx-central.io",
	Version: "v1alpha1",
	Kind:    "CustomResourceDefinition",
}

type gvkFactory struct {
}

func (f *gvkFactory) Create() *crds.CustomResourceDefinition {
	return &crds.CustomResourceDefinition{}
}

func NewCrdFactory() MetadatableFactory[*crds.CustomResourceDefinition] {
	return &gvkFactory{}
}

type customDocFactory struct {
}

func (f *customDocFactory) Create() *documents.CustomDocument {
	return &documents.CustomDocument{}
}

func NewCustomDocumentFactory() MetadatableFactory[*documents.CustomDocument] {
	return &customDocFactory{}
}

func GetSystemGvks() []*idl_common.GroupVersionKindSpec {
	return []*idl_common.GroupVersionKindSpec{
		proto.Clone(NatIpGvk).(*idl_common.GroupVersionKindSpec),
		proto.Clone(NamespaceGvk).(*idl_common.GroupVersionKindSpec),
		proto.Clone(EndpointGvk).(*idl_common.GroupVersionKindSpec),
		proto.Clone(CrdGvk).(*idl_common.GroupVersionKindSpec),
	}
}

func IsSystemGvk(gvk *idl_common.GroupVersionKindSpec) bool {
	for _, systemGvk := range GetSystemGvks() {
		if EqualsGvk(gvk, systemGvk) {
			return true
		}
	}
	return false
}
