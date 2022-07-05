package gvk_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/gvks"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type Controller struct {
	gvks.UnimplementedGroupVersionKindServerServer
}

func (c *Controller) GetGroupVersionKind(ctx context.Context, req *idl_common.SingleObjectReq) (*gvks.GetGroupVersionKindRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) ListGroupVersionKinds(ctx context.Context, req *idl_common.ListObjectReq) (*gvks.ListGroupVersionKindsRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) CreateGroupVersionKind(ctx context.Context, req *gvks.CreateGroupVersionKindReq) (*gvks.CreateGroupVersionKindRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) DeleteGroupVersionKind(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	//TODO implement me
	panic("implement me")
}
