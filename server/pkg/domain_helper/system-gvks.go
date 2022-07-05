package domain_helper

import (
	"github.com/krafton-hq/red-fox/apis/documents"
	gvks "github.com/krafton-hq/red-fox/apis/gvks"
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

var GvkGvk = &idl_common.GroupVersionKindSpec{
	Group:   "red-fox.sbx-central.io",
	Version: "v1alpha1",
	Kind:    "GroupVersionKind",
}

type gvkFactory struct {
}

func (f *gvkFactory) Create() *gvks.GroupVersionKind {
	return &gvks.GroupVersionKind{}
}

func NewGvkFactory() MetadatableFactory[*gvks.GroupVersionKind] {
	return &gvkFactory{}
}

func GetSystemGvks() []*idl_common.GroupVersionKindSpec {
	return []*idl_common.GroupVersionKindSpec{
		proto.Clone(NatIpGvk).(*idl_common.GroupVersionKindSpec),
		proto.Clone(NamespaceGvk).(*idl_common.GroupVersionKindSpec),
		proto.Clone(EndpointGvk).(*idl_common.GroupVersionKindSpec),
		proto.Clone(GvkGvk).(*idl_common.GroupVersionKindSpec),
	}
}
