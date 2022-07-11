package crd_service

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/crds"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
	"github.com/krafton-hq/red-fox/server/repositories/repository_manager"
)

type Service struct {
	repository apiobject_repository.ClusterRepository[*crds.CustomResourceDefinition]

	repoManager *repository_manager.Manager
}

func NewService(repository apiobject_repository.ClusterRepository[*crds.CustomResourceDefinition], repoManager *repository_manager.Manager) *Service {
	return &Service{
		repository:  repository,
		repoManager: repoManager,
	}
}

func (s *Service) Init(ctx context.Context) error {
	crdList, err := s.repository.List(ctx, nil)
	if err != nil {
		return err
	}

	for _, crd := range crdList {
		err = s.repoManager.AddCustomRepository(ctx, crd.Spec.Gvk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Get(ctx context.Context, name string) (*crds.CustomResourceDefinition, error) {
	_, _, err := domain_helper.ParseGvkName(name)
	if err != nil {
		return nil, err
	}

	return s.repository.Get(ctx, name)
}

func (s *Service) List(ctx context.Context, labelSelectors map[string]string) ([]*crds.CustomResourceDefinition, error) {
	return s.repository.List(ctx, labelSelectors)
}

func (s *Service) Create(ctx context.Context, obj *crds.CustomResourceDefinition) error {
	err := s.repository.Create(ctx, obj)
	if err != nil {
		return err
	}

	err = s.repoManager.AddCustomRepository(ctx, obj.Spec.Gvk)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, obj *crds.CustomResourceDefinition) error {
	err := s.repository.Update(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, name string) error {
	err := s.repository.Delete(ctx, name)
	if err != nil {
		return err
	}

	err = s.repoManager.DeleteCustomRepository(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
