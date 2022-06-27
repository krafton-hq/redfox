package natip_service

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
)

type Service struct {
	repository apiobject_repository.NamespacedRepository[*documents.NatIp]
}

func NewService(repository apiobject_repository.NamespacedRepository[*documents.NatIp]) *Service {
	return &Service{repository: repository}
}

func (s *Service) Get(ctx context.Context, namespace string, name string) (*documents.NatIp, error) {
	return s.repository.Get(ctx, namespace, name)
}

func (s *Service) List(ctx context.Context, labelSelectors map[string]string) ([]*documents.NatIp, error) {
	return s.repository.List(ctx, labelSelectors)
}

func (s *Service) ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]*documents.NatIp, error) {
	return s.repository.ListNamespaced(ctx, namespace, labelSelectors)
}

func (s *Service) Create(ctx context.Context, natIp *documents.NatIp) error {

	return s.repository.Create(ctx, natIp)
}

func (s *Service) Update(ctx context.Context, natIp *documents.NatIp) error {
	return s.repository.Update(ctx, natIp)
}

func (s *Service) Delete(ctx context.Context, namespace string, name string) error {
	return s.repository.Delete(ctx, namespace, name)
}
