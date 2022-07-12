package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/controllers/utils"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/services/services"
)

type NatIpController struct {
	documents.UnimplementedNatIpServerServer

	service services.NamespacedService[*documents.NatIp]
}

func NewNatIpDocController(service services.NamespacedService[*documents.NatIp]) *NatIpController {
	return &NatIpController{service: service}
}

func (c *NatIpController) GetNatIp(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetNatIpRes, error) {
	if req.Name == "" {
		return &documents.GetNatIpRes{CommonRes: utils.CommonResNotEmpty("name")}, nil
	}
	if req.Namespace == nil || *req.Namespace == "" {
		return &documents.GetNatIpRes{CommonRes: utils.CommonResNotEmpty("namespace")}, nil
	}

	natIp, err := c.service.Get(ctx, *req.Namespace, req.Name)
	if err != nil {
		return &documents.GetNatIpRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil
	}

	return &documents.GetNatIpRes{
		CommonRes: &idl_common.CommonRes{Message: "Get NatIp Success"},
		NatIp:     natIp,
	}, nil
}

func (c *NatIpController) ListNatIps(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListNatIpsRes, error) {
	namespaced := req.Namespace != ""

	var natIps []*documents.NatIp
	var err error
	if namespaced {
		natIps, err = c.service.ListNamespaced(ctx, req.Namespace, req.LabelSelectors)
	} else {
		natIps, err = c.service.List(ctx, req.LabelSelectors)
	}
	if err != nil {
		return &documents.ListNatIpsRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil
	}

	return &documents.ListNatIpsRes{
		CommonRes: &idl_common.CommonRes{Message: "List NatIps Success"},
		NatIps:    natIps,
	}, nil
}

func (c *NatIpController) CreateNatIp(ctx context.Context, req *documents.DesiredNatIpReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.NatIp); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if err := domain_helper.ValidationDiscoverableName(req.NatIp); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if req.NatIp.Metadata.Namespace == "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}
	if err := domain_helper.ValidationNatIpSpec(req.NatIp.Spec); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	err := c.service.Create(ctx, req.NatIp)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Create NatIp Success"}, nil
}

func (c *NatIpController) UpdateNatIp(ctx context.Context, req *documents.DesiredNatIpReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.NatIp); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if err := domain_helper.ValidationDiscoverableName(req.NatIp); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if req.NatIp.Metadata.Namespace == "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}
	if err := domain_helper.ValidationNatIpSpec(req.NatIp.Spec); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	err := c.service.Update(ctx, req.NatIp)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Update NatIp Success"}, nil
}

func (c *NatIpController) DeleteNatIp(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
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

	return &idl_common.CommonRes{Message: "Delete NatIp Success"}, nil
}
