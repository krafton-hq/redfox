package custom_document_service

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/repositories/repository_manager"
)

type Service struct {
	repoManager *repository_manager.Manager
}

func NewService(repoManager *repository_manager.Manager) *Service {
	return &Service{repoManager: repoManager}
}

func (s *Service) Get(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, name string) (*documents.CustomDocument, error) {
	repository, found := s.repoManager.GetCustomRepository(ctx, domain_helper.GetGvkName(gvk))
	if !found {
		return nil, errors.NewNotFound(fmt.Sprintf("gvk{%s}", domain_helper.GetGvkName(gvk)))
	}
	return repository.Get(ctx, namespace, name)
}

func (s *Service) List(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, labelSelectors map[string]string) ([]*documents.CustomDocument, error) {
	repository, found := s.repoManager.GetCustomRepository(ctx, domain_helper.GetGvkName(gvk))
	if !found {
		return nil, errors.NewNotFound(fmt.Sprintf("gvk{%s}", domain_helper.GetGvkName(gvk)))
	}
	return repository.List(ctx, labelSelectors)
}

func (s *Service) ListNamespaced(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, labelSelectors map[string]string) ([]*documents.CustomDocument, error) {
	repository, found := s.repoManager.GetCustomRepository(ctx, domain_helper.GetGvkName(gvk))
	if !found {
		return nil, errors.NewNotFound(fmt.Sprintf("gvk{%s}", domain_helper.GetGvkName(gvk)))
	}
	return repository.ListNamespaced(ctx, namespace, labelSelectors)
}

func (s *Service) Create(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, Document *documents.CustomDocument) error {
	repository, found := s.repoManager.GetCustomRepository(ctx, domain_helper.GetGvkName(gvk))
	if !found {
		return errors.NewNotFound(fmt.Sprintf("gvk{%s}", domain_helper.GetGvkName(gvk)))
	}
	return repository.Create(ctx, Document)
}

func (s *Service) Update(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, Document *documents.CustomDocument) error {
	repository, found := s.repoManager.GetCustomRepository(ctx, domain_helper.GetGvkName(gvk))
	if !found {
		return errors.NewNotFound(fmt.Sprintf("gvk{%s}", domain_helper.GetGvkName(gvk)))
	}
	return repository.Update(ctx, Document)
}

func (s *Service) Delete(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, name string) error {
	repository, found := s.repoManager.GetCustomRepository(ctx, domain_helper.GetGvkName(gvk))
	if !found {
		return errors.NewNotFound(fmt.Sprintf("gvk{%s}", domain_helper.GetGvkName(gvk)))
	}
	return repository.Delete(ctx, namespace, name)

}
