package app_lifecycle_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type AppLifecycle struct {
	app_life.UnimplementedApplicationLifecycleServer
}

func NewAppLifecycle() *AppLifecycle {
	return &AppLifecycle{}
}

func (c *AppLifecycle) Version(ctx context.Context, req *idl_common.CommonReq) (*idl_common.CommonRes, error) {
	return &idl_common.CommonRes{
		Message: "bar",
	}, nil
}

func (c *AppLifecycle) Livez(ctx context.Context, req *idl_common.CommonReq) (*idl_common.CommonRes, error) {
	return &idl_common.CommonRes{
		Message: "bar",
	}, nil
}

func (c *AppLifecycle) Readyz(ctx context.Context, req *idl_common.CommonReq) (*idl_common.CommonRes, error) {
	return &idl_common.CommonRes{
		Message: "bar",
	}, nil
}
