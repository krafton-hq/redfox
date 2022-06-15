package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type NatIpDocController struct {
	documents.UnimplementedNatIpServerServer
}

func (c *NatIpDocController) GetNatIp(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetNatIpRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *NatIpDocController) ListNatIps(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListNatIpsRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *NatIpDocController) CreateNatIp(ctx context.Context, req *documents.DesiredNatIpReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *NatIpDocController) UpdateNatIp(ctx context.Context, req *documents.DesiredNatIpReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *NatIpDocController) DeleteNatIp(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}
