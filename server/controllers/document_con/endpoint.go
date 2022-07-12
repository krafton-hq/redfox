package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/controllers/utils"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/services/services"
)

type EndpointController struct {
	documents.UnimplementedEndpointServerServer

	service services.NamespacedService[*documents.Endpoint]
}

func NewEndpointController(service services.NamespacedService[*documents.Endpoint]) *EndpointController {
	return &EndpointController{service: service}
}

func (c *EndpointController) GetEndpoint(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetEndpointRes, error) {
	if req.Name == "" {
		return &documents.GetEndpointRes{CommonRes: utils.CommonResNotEmpty("name")}, nil
	}
	if req.Namespace == nil || *req.Namespace == "" {
		return &documents.GetEndpointRes{CommonRes: utils.CommonResNotEmpty("namespace")}, nil
	}

	endpoint, err := c.service.Get(ctx, *req.Namespace, req.Name)
	if err != nil {
		return &documents.GetEndpointRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil
	}

	return &documents.GetEndpointRes{
		CommonRes: &idl_common.CommonRes{Message: "Get Endpoint Success"},
		Endpoint:  endpoint,
	}, nil
}

func (c *EndpointController) ListEndpoints(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListEndpointsRes, error) {
	namespaced := req.Namespace != ""

	var endpoints []*documents.Endpoint
	var err error
	if namespaced {
		endpoints, err = c.service.ListNamespaced(ctx, req.Namespace, req.LabelSelectors)
	} else {
		endpoints, err = c.service.List(ctx, req.LabelSelectors)
	}
	if err != nil {
		return &documents.ListEndpointsRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil

	}

	return &documents.ListEndpointsRes{
		CommonRes: &idl_common.CommonRes{Message: "List Endpoints Success"},
		Endpoints: endpoints,
	}, nil
}

func (c *EndpointController) CreateEndpoint(ctx context.Context, req *documents.DesiredEndpointReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.Endpoint); err != nil {
		return utils.InvalidArguments(err), nil
	}
	if err := domain_helper.ValidationDiscoverableName(req.Endpoint); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if req.Endpoint.Metadata.Namespace == "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}

	err := c.service.Create(ctx, req.Endpoint)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Create Endpoint Success"}, nil
}

func (c *EndpointController) UpdateEndpoint(ctx context.Context, req *documents.DesiredEndpointReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.Endpoint); err != nil {
		return utils.InvalidArguments(err), nil
	}
	if err := domain_helper.ValidationDiscoverableName(req.Endpoint); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if req.Endpoint.Metadata.Namespace == "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}

	err := c.service.Update(ctx, req.Endpoint)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Update Endpoint Success"}, nil
}

func (c *EndpointController) DeleteEndpoint(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	if req.Name == "" {
		return utils.CommonResNotEmpty("name"), nil
	}
	if req.Namespace == nil || *req.Namespace != "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}

	err := c.service.Delete(ctx, *req.Namespace, req.Name)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Delete Endpoint Success"}, nil
}
