package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/controllers/utils"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/services/services"
)

type CustomDocumentController struct {
	documents.UnimplementedCustomDocumentServerServer

	service services.NamespacedGvkService[*documents.CustomDocument]
}

func NewCustomDocumentController(service services.NamespacedGvkService[*documents.CustomDocument]) *CustomDocumentController {
	return &CustomDocumentController{service: service}
}

func (c *CustomDocumentController) GetCustomDocument(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetCustomDocumentRes, error) {
	if req.Name == "" {
		return &documents.GetCustomDocumentRes{CommonRes: utils.CommonResNotEmpty("name")}, nil
	}
	if req.Namespace == nil || *req.Namespace == "" {
		return &documents.GetCustomDocumentRes{CommonRes: utils.CommonResNotEmpty("namespace")}, nil
	}
	if req.Gvk == nil {
		return &documents.GetCustomDocumentRes{CommonRes: utils.CommonResNotEmpty("gvk")}, nil
	}

	document, err := c.service.Get(ctx, req.Gvk, *req.Namespace, req.Name)
	if err != nil {
		return &documents.GetCustomDocumentRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil
	}

	return &documents.GetCustomDocumentRes{
		CommonRes: &idl_common.CommonRes{Message: "Get Document Success"},
		Document:  document,
	}, nil
}

func (c *CustomDocumentController) ListCustomDocuments(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListCustomDocumentsRes, error) {
	if req.Gvk == nil {
		return &documents.ListCustomDocumentsRes{CommonRes: utils.CommonResNotEmpty("gvk")}, nil
	}

	namespaced := req.Namespace != ""

	var documentList []*documents.CustomDocument
	var err error
	if namespaced {
		documentList, err = c.service.ListNamespaced(ctx, req.Gvk, req.Namespace, req.LabelSelectors)
	} else {
		documentList, err = c.service.List(ctx, req.Gvk, req.LabelSelectors)
	}
	if err != nil {
		return &documents.ListCustomDocumentsRes{CommonRes: utils.CommonResWithErrorTypes(err)}, nil

	}

	return &documents.ListCustomDocumentsRes{
		CommonRes: &idl_common.CommonRes{Message: "List Documents Success"},
		Documents: documentList,
	}, nil
}

func (c *CustomDocumentController) CreateCustomDocument(ctx context.Context, req *documents.DesiredCustomDocumentReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.Document); err != nil {
		return utils.InvalidArguments(err), nil
	}
	if err := domain_helper.ValidationQualifiedName(req.Document); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if req.Document.Metadata.Namespace == "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}

	gvk, err := domain_helper.CreateGvkFromMetadatable(req.Document)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	err = c.service.Create(ctx, gvk, req.Document)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Create Document Success"}, nil
}

func (c *CustomDocumentController) UpdateCustomDocument(ctx context.Context, req *documents.DesiredCustomDocumentReq) (*idl_common.CommonRes, error) {
	if err := domain_helper.ValidationMetadatable(req.Document); err != nil {
		return utils.InvalidArguments(err), nil
	}
	if err := domain_helper.ValidationQualifiedName(req.Document); err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}
	if req.Document.Metadata.Namespace == "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}

	gvk, err := domain_helper.CreateGvkFromMetadatable(req.Document)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	err = c.service.Update(ctx, gvk, req.Document)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Update Document Success"}, nil
}

func (c *CustomDocumentController) DeleteCustomDocument(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	if req.Name == "" {
		return utils.CommonResNotEmpty("name"), nil
	}
	if req.Namespace == nil || *req.Namespace != "" {
		return utils.CommonResNotEmpty("namespace"), nil
	}
	if req.Gvk == nil {
		return utils.CommonResNotEmpty("gvk"), nil
	}

	err := c.service.Delete(ctx, req.Gvk, *req.Namespace, req.Name)
	if err != nil {
		return utils.CommonResWithErrorTypes(err), nil
	}

	return &idl_common.CommonRes{Message: "Delete Document Success"}, nil
}
