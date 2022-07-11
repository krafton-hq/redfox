package repository_manager

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/crds"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/repositories/apiobject_repository"
	"go.uber.org/zap"
)

type Manager struct {
	customClusterRepoFactory apiobject_repository.ClusterRepositoryFactory[*documents.CustomDocument]

	customRepositories map[string]apiobject_repository.NamespacedRepository[*documents.CustomDocument]
	natIpRepository    apiobject_repository.NamespacedRepository[*documents.NatIp]
	endpointRepository apiobject_repository.NamespacedRepository[*documents.Endpoint]

	namespaceRepository apiobject_repository.ClusterRepository[*namespaces.Namespace]
	crdRepository       apiobject_repository.ClusterRepository[*crds.CustomResourceDefinition]
}

func NewManager(customClusterRepoFactory apiobject_repository.ClusterRepositoryFactory[*documents.CustomDocument], natIpRepository apiobject_repository.NamespacedRepository[*documents.NatIp], endpointRepository apiobject_repository.NamespacedRepository[*documents.Endpoint], namespaceRepository apiobject_repository.ClusterRepository[*namespaces.Namespace], crdRepository apiobject_repository.ClusterRepository[*crds.CustomResourceDefinition]) *Manager {
	return &Manager{
		customClusterRepoFactory: customClusterRepoFactory,
		customRepositories:       map[string]apiobject_repository.NamespacedRepository[*documents.CustomDocument]{},
		natIpRepository:          natIpRepository,
		endpointRepository:       endpointRepository,
		namespaceRepository:      namespaceRepository,
		crdRepository:            crdRepository,
	}
}

func (m *Manager) GetNamespacedRepositoryMetadatas() []apiobject_repository.NamespacedRepositoryMetadata {
	var metadatas []apiobject_repository.NamespacedRepositoryMetadata
	for _, repo := range m.customRepositories {
		metadatas = append(metadatas, repo)
	}
	metadatas = append(metadatas, m.natIpRepository)
	metadatas = append(metadatas, m.endpointRepository)
	return metadatas
}

func (m *Manager) GetCustomRepository(ctx context.Context, gvkName string) (apiobject_repository.NamespacedRepository[*documents.CustomDocument], bool) {
	val, ok := m.customRepositories[gvkName]
	return val, ok
}

func (m *Manager) AddCustomRepository(ctx context.Context, gvk *idl_common.GroupVersionKindSpec) error {
	gvkName := domain_helper.GetGvkName(gvk)
	if _, found := m.customRepositories[gvkName]; found {
		return errors.NewInvalidArguments(fmt.Sprintf("Already Exists Gvk name:'%s'", gvkName))
	}

	repository := apiobject_repository.NewGenericNamespacedRepository[*documents.CustomDocument](gvk, m.customClusterRepoFactory)
	zap.S().Infow("CustomDoc Repository is Created", "gvk", gvk)
	m.customRepositories[gvkName] = repository
	return nil
}

func (m *Manager) DeleteCustomRepository(ctx context.Context, gvkName string) error {
	if _, found := m.customRepositories[gvkName]; !found {
		return errors.NewInvalidArguments(fmt.Sprintf("Requested Gvk Is Not Exists name:'%s'", gvkName))
	}

	err := m.customRepositories[gvkName].Truncate(ctx)
	if err != nil {
		zap.S().Infow("CustomDoc Repository Delete Failed while Truncate Repository", "gvkName", gvkName, "error", err)
		return err
	}
	zap.S().Infow("CustomDoc Repository is Deleted", "gvkName", gvkName)

	delete(m.customRepositories, gvkName)
	return nil
}
