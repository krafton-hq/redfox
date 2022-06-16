package app_lifecycle_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/app_lifecycle"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type GrpcController struct {
	app_lifecycle.UnimplementedApplicationLifecycleServer
}

func NewAppLifecycle() *GrpcController {
	return &GrpcController{}
}

func (c *GrpcController) Version(ctx context.Context, req *idl_common.CommonReq) (*idl_common.CommonRes, error) {
	return &idl_common.CommonRes{
		Message: "bar",
	}, nil
}

func (c *GrpcController) Livez(ctx context.Context, req *idl_common.CommonReq) (*idl_common.CommonRes, error) {
	return &idl_common.CommonRes{
		Message: "bar",
	}, nil
}

func (c *GrpcController) Readyz(ctx context.Context, req *idl_common.CommonReq) (*idl_common.CommonRes, error) {
	return &idl_common.CommonRes{
		Message: "bar",
	}, nil
}
