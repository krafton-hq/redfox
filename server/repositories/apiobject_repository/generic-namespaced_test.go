package apiobject_repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
)

func TestInmemoryNamespacedRepository_Create(t *testing.T) {
	repo := NewGenericNamespacedRepository[*documents.NatIp](domain_helper.NatIpGvk, NewMysqlClusterRepositoryFactory[*documents.NatIp](domain_helper.NewNatIpFactory(), tr))
	repo.EnableNamespace(context.TODO(), "default")
	repo.EnableNamespace(context.TODO(), "musong")

	ctx := tr.WithDatabaseContext(context.TODO())
	natIp, err := repo.Get(ctx, "default", "my-first-object")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", natIp)
}
