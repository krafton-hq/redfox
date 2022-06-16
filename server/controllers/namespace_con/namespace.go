package namespace_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/controllers/utils"
	"github.com/krafton-hq/red-fox/server/pkg/validation"
	"github.com/krafton-hq/red-fox/server/services/namespace_service"
)

type Controller struct {
	namespaces.UnimplementedNamespaceServerServer

	service *namespace_service.Service
}

func NewController(service *namespace_service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetNamespace(ctx context.Context, req *idl_common.SingleObjectReq) (*namespaces.GetNamespaceRes, error) {
	if req.Name == "" {
		return &namespaces.GetNamespaceRes{CommonRes: utils.CommonResNotEmpty("name")}, nil
	}

	namespace, err := c.service.Get(ctx, req.Name)
	if err != nil {
		return &namespaces.GetNamespaceRes{CommonRes: utils.CommonResInternalError(err)}, nil
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
		return &namespaces.ListNamespacesRes{CommonRes: utils.CommonResInternalError(err)}, nil
	}

	return &namespaces.ListNamespacesRes{
		CommonRes:  &idl_common.CommonRes{Message: "List Namespaces Success"},
		Namespaces: nss,
	}, nil
}

func (c *Controller) CreateNamespace(ctx context.Context, req *namespaces.CreateNamespaceReq) (*idl_common.CommonRes, error) {
	if errors := validation.IsDiscoveryName(req.Namespace.Metadata.Name); len(errors) > 0 {
		return utils.CommonResDnsLabel("name", errors), nil
	}
	if req.Namespace.Metadata.Namespace != "" {
		return utils.CommonResEmpty("Namespace", "namespace"), nil
	}
	for _, objMeta := range req.Namespace.Spec.ApiObjects {
		if errors := validation.IsGroup(objMeta.Group); len(errors) > 0 {
			return utils.CommonResDnsLabel("group", errors), nil
		}
		if rawErr := validation.IsVersion(objMeta.Version); rawErr != "" {
			return utils.CommonResFieldValid("version", []string{rawErr}), nil
		}
		if rawErr := validation.IsKind(objMeta.Kind); rawErr != "" {
			return utils.CommonResFieldValid("kind", []string{rawErr}), nil
		}
	}

	err := c.service.Create(ctx, req.Namespace)
	if err != nil {
		return utils.CommonResInternalError(err), nil
	}

	return &idl_common.CommonRes{Message: "Create Namespace Success"}, nil
}

func (c *Controller) DeleteNamespaces(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	if req.Name == "" {
		return utils.CommonResNotEmpty("name"), nil
	}

	err := c.service.Delete(ctx, req.Name)
	if err != nil {
		return utils.CommonResInternalError(err), nil
	}

	return &idl_common.CommonRes{Message: "Delete Namespace Success"}, nil
}
