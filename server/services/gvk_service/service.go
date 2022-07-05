package gvk_service

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/gvks"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
)

type Service struct {
	repository apiobject_repository.ClusterRepository[*gvks.GroupVersionKind]

	gvkCache map[string]*gvks.GroupVersionKind
}

func NewService(repository apiobject_repository.ClusterRepository[*gvks.GroupVersionKind]) *Service {
	return &Service{
		repository: repository,
		gvkCache:   map[string]*gvks.GroupVersionKind{},
	}
}

func (s *Service) Init(ctx context.Context) error {
	gvkList, err := s.repository.List(ctx, nil)
	if err != nil {
		return err
	}

	s.gvkCache = map[string]*gvks.GroupVersionKind{}
	for _, gvk := range gvkList {
		s.gvkCache[domain_helper.GetGvkName(gvk.Spec.Detail)] = gvk
	}
	return nil
}

func (s *Service) Get(ctx context.Context, name string) (*gvks.GroupVersionKind, error) {
	_, _, err := domain_helper.ParseGvkName(name)
	if err != nil {
		return nil, err
	}

	if gvk, found := s.gvkCache[name]; found {
		return gvk, nil
	}
	return nil, errors.NewNotFound(name)
}

func (s *Service) List(ctx context.Context, labelSelectors map[string]string) ([]*gvks.GroupVersionKind, error) {
	return s.repository.List(ctx, labelSelectors)
}

func (s *Service) Create(ctx context.Context, obj *gvks.GroupVersionKind) error {
	return nil
}

func (s *Service) Update(ctx context.Context, obj *gvks.GroupVersionKind) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
