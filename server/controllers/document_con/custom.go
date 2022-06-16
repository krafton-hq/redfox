package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type CustomDocumentController struct {
	documents.UnimplementedCustomDocumentServerServer
}

func (c *CustomDocumentController) GetCustomDocument(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetCustomDocumentRes, error) {
	panic("implement me")
}

func (c *CustomDocumentController) ListCustomDocuments(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListCustomDocumentsRes, error) {
	panic("implement me")
}

func (c *CustomDocumentController) CreateCustomDocument(ctx context.Context, req *documents.DesiredCustomDocumentReq) (*idl_common.CommonRes, error) {
	panic("implement me")
}

func (c *CustomDocumentController) UpdateCustomDocument(ctx context.Context, req *documents.DesiredCustomDocumentReq) (*idl_common.CommonRes, error) {
	panic("implement me")
}

func (c *CustomDocumentController) DeleteCustomDocument(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	panic("implement me")
}
