package document_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type DocumentController struct {
	documents.UnimplementedCustomDocumentServerServer
}

func (c *DocumentController) GetCustomDocument(ctx context.Context, req *idl_common.SingleObjectReq) (*documents.GetCustomDocumentRes, error) {
	panic("implement me")
}

func (c *DocumentController) ListCustomDocuments(ctx context.Context, req *idl_common.ListObjectReq) (*documents.ListCustomDocumentsRes, error) {
	panic("implement me")
}

func (c *DocumentController) CreateCustomDocument(ctx context.Context, req *documents.DesiredCustomDocumentReq) (*idl_common.CommonRes, error) {
	panic("implement me")
}

func (c *DocumentController) UpdateCustomDocument(ctx context.Context, req *documents.DesiredCustomDocumentReq) (*idl_common.CommonRes, error) {
	panic("implement me")
}

func (c *DocumentController) DeleteCustomDocument(ctx context.Context, req *idl_common.SingleObjectReq) (*idl_common.CommonRes, error) {
	panic("implement me")
}
