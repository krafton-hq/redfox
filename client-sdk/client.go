package client_sdk

import (
	"github.com/krafton-hq/red-fox/apis/api_resources"
	"github.com/krafton-hq/red-fox/apis/app_lifecycle"
	"github.com/krafton-hq/red-fox/apis/crds"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/namespaces"
)

type RedFoxClient struct {
	documents.NatIpServerClient
	documents.EndpointServerClient
	documents.CustomDocumentServerClient
	namespaces.NamespaceServerClient
	crds.CustomDocumentDefinitionServerClient
	api_resources.ApiResourcesServerClient
	app_lifecycle.ApplicationLifecycleClient
}
