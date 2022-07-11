package namespace_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/controllers/utils"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/services/service_helper"
)

type Controller struct {
	namespaces.UnimplementedNamespaceServerServer

	service service_helper.ClusterService[*namespaces.Namespace]
}

func NewController(service service_helper.ClusterService[*namespaces.Namespace]) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetNamespace(ctx context.Context, req *idl_common.SingleObjectReq) (*namespaces.GetNamespaceRes, error) {
	if req.Name == "" {
		return &namespaces.GetNamespaceRes{CommonRes: utils.CommonResNotEmpty("name")}, nil
	}

	namespace, err := c.service.Get(ctx, req.Name)
	if err != nil {
		return &namespaces.GetNamespaceRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil
	}

	return &namespaces.GetNamespaceRes{
		CommonRes: &idl_common.CommonRes{Message: "Get Namespace Success"},
		Namespace: namespace,
	}, nil
}

func (c *Controller) ListNamespaces(ctx context.Context, req *idl_common.ListObjectReq) (*namespaces.ListNamespacesRes, error) {
	if req.Namespace != "" {
		return &namespaces.ListNamespacesRes{CommonRes: utils.CommonResEmpty("Namespace", "namespace")}, nil
	}

	nss, err := c.service.List(ctx, req.LabelSelectors)
	if err != nil {
		return &namespaces.ListNamespacesRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil
	}

	return &namespaces.ListNamespacesRes{
		CommonRes:  &idl_common.CommonRes{Message: "List Namespaces Success"},
		Namespaces: nss,
	}, nil
}

func (c *Controller) CreateNamespace(ctx context.Context, req *namespaces.CreateNamespaceReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.Namespace); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if err := domain_helper.ValidationDiscoverableName(req.Namespace); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if err := domain_helper.ValidationNamespaceSpec(req.Namespace.Spec); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	err := c.service.Create(ctx, req.Namespace)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Create Namespace Success"}, nil
}

func (c *Controller) UpdateNamespace(ctx context.Context, req *namespaces.UpdateNamespaceReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.Namespace); err != nil {
		return utils.InvalidArguments(err), nil
	}
	if err := domain_helper.ValidationDiscoverableName(req.Namespace); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if err := domain_helper.ValidationNamespaceSpec(req.Namespace.Spec); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	err := c.service.Update(ctx, req.Namespace)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Update Namespace Success"}, nil
}

func (c *Controller) DeleteNamespaces(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	if req.Name == "" {
		return utils.CommonResNotEmpty("name"), nil
	}

	err := c.service.Delete(ctx, req.Name)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Delete Namespace Success"}, nil
}
