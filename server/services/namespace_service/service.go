package namespace_service

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	repository      apiobject_repository.ClusterRepository[*namespaces.Namespace]
	namespacedRepos []apiobject_repository.NamespacedRepositoryMetadata
}

func NewService(repository apiobject_repository.ClusterRepository[*namespaces.Namespace], namespacedRepos []apiobject_repository.NamespacedRepositoryMetadata) *Service {
	return &Service{repository: repository, namespacedRepos: namespacedRepos}
}

func (s *Service) Init(ctx context.Context) error {
	nss, err := s.repository.List(ctx, nil)
	if err != nil {
		return err
	}
	for _, namespace := range nss {
		s.updateNamespacedRepositories(ctx, namespace)
	}
	return nil
}

func (s *Service) Get(ctx context.Context, name string) (*namespaces.Namespace, error) {
	return s.repository.Get(ctx, name)
}

func (s *Service) List(ctx context.Context, labelSelectors map[string]string) ([]*namespaces.Namespace, error) {
	return s.repository.List(ctx, labelSelectors)
}

func (s *Service) Create(ctx context.Context, namespace *namespaces.Namespace) error {
	err := s.repository.Create(ctx, namespace)
	if err != nil {
		return err
	}

	s.updateNamespacedRepositories(ctx, namespace)
	return nil
}

func (s *Service) Update(ctx context.Context, namespace *namespaces.Namespace) error {
	err := s.repository.Update(ctx, namespace)
	if err != nil {
		return err
	}

	s.updateNamespacedRepositories(ctx, namespace)
	return nil
}

func (s *Service) updateNamespacedRepositories(ctx context.Context, namespace *namespaces.Namespace) {
	nsEnableTargets := map[string]apiobject_repository.NamespacedRepositoryMetadata{}
	for _, objMeta := range namespace.Spec.ApiObjects {
		for _, repoMetadata := range s.namespacedRepos {
			if proto.Equal(objMeta, repoMetadata.Info()) {
				if objMeta.Enabled {
					nsEnableTargets[objMeta.Kind] = repoMetadata
				}
			}
		}
	}

	for _, repoMetadata := range s.namespacedRepos {
		gvk := repoMetadata.Info()
		if _, exist := nsEnableTargets[gvk.Kind]; exist {
			result := repoMetadata.EnableNamespace(ctx, namespace.Metadata.Name)
			zap.S().Infow("Enabled Namespace", "result", result, "name", namespace.Metadata.Name, "gvk", gvk.String())
		} else {
			result := repoMetadata.DisableNamespace(ctx, namespace.Metadata.Name)
			zap.S().Infow("Disabled Namespace", "result", result, "name", namespace.Metadata.Name, "gvk", gvk.String())
		}
	}
}

func (s *Service) Delete(ctx context.Context, name string) error {
	err := s.repository.Delete(ctx, name)
	if err != nil {
		return err
	}

	for _, metadata := range s.namespacedRepos {
		result := metadata.DisableNamespace(ctx, name)
		zap.S().Infow("Disabled Namespace", "result", result, "name", name, "gvk", metadata.Info().String())
	}
	return nil
}
