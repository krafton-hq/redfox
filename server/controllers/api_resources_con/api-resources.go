package api_resources_con

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/api_resources"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/repositories/repository_manager"
)

type Controller struct {
	api_resources.UnimplementedApiResourcesServerServer

	repoManager *repository_manager.Manager
}

func NewController(repoManager *repository_manager.Manager) *Controller {
	return &Controller{repoManager: repoManager}
}

func (c *Controller) ListApiResources(ctx context.Context, req *idl_common.CommonReq) (*api_resources.ListApiResourcesRes, error) {
	var apiResources []*idl_common.ApiResourceSpec
	for _, repositoryMetadata := range c.repoManager.GetNamespacedRepositoryMetadatas() {
		gvk := repositoryMetadata.Info()
		apiResources = append(apiResources, &idl_common.ApiResourceSpec{
			Name:  domain_helper.GetGvkName(gvk),
			Scope: idl_common.Scope_Namespaced,
			Gvk:   gvk,
		})
	}

	for _, gvk := range c.repoManager.GetClusterRepositoryMetadatas() {
		apiResources = append(apiResources, &idl_common.ApiResourceSpec{
			Name:  domain_helper.GetGvkName(gvk),
			Scope: idl_common.Scope_Cluster,
			Gvk:   gvk,
		})
	}
	return &api_resources.ListApiResourcesRes{
		CommonRes:    &idl_common.CommonRes{Message: "List ApiResources Success"},
		ApiResources: apiResources,
	}, nil
}
