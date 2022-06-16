package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type EndpointController struct {
	documents.UnimplementedEndpointServerServer
}

func (c *EndpointController) GetEndpoint(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetEndpointRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *EndpointController) ListEndpoints(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListEndpointsRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *EndpointController) CreateEndpoint(ctx context.Context, req *documents.DesiredEndpointReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *EndpointController) UpdateEndpoint(ctx context.Context, req *documents.DesiredEndpointReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *EndpointController) DeleteEndpoint(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}
