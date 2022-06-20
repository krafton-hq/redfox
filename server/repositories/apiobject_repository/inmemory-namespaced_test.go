package apiobject_repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
)

func TestInmemoryNamespacedRepository_Create(t *testing.T) {
	repo := NewInmemoryNamespacedRepository[*documents.NatIp](domain_helper.NatIpGvk, newMysqlClusterRepositoryFactory[*documents.NatIp](domain_helper.NewNatIpFactory()))
	repo.EnableNamespace(context.TODO(), "default")
	repo.EnableNamespace(context.TODO(), "musong")

	ctx := context.WithValue(context.TODO(), databaseKey{}, db)
	natIp, err := repo.Get(ctx, "default", "my-first-object")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", natIp)
}
