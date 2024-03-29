// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/krafton-hq/redfox/pkg/generated/clientset/versioned/typed/redfox/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMetadataV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMetadataV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeMetadataV1alpha1) IngressAddresses(namespace string) v1alpha1.IngressAddressInterface {
	return &FakeIngressAddresses{c, namespace}
}

func (c *FakeMetadataV1alpha1) LatestCommits(namespace string) v1alpha1.LatestCommitInterface {
	return &FakeLatestCommits{c, namespace}
}

func (c *FakeMetadataV1alpha1) LatestVersions(namespace string) v1alpha1.LatestVersionInterface {
	return &FakeLatestVersions{c, namespace}
}

func (c *FakeMetadataV1alpha1) NatIps(namespace string) v1alpha1.NatIpInterface {
	return &FakeNatIps{c, namespace}
}

func (c *FakeMetadataV1alpha1) RegionMetadatas(namespace string) v1alpha1.RegionMetadataInterface {
	return &FakeRegionMetadatas{c, namespace}
}

func (c *FakeMetadataV1alpha1) Versions(namespace string) v1alpha1.VersionInterface {
	return &FakeVersions{c, namespace}
}

func (c *FakeMetadataV1alpha1) VersionCounts(namespace string) v1alpha1.VersionCountInterface {
	return &FakeVersionCounts{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMetadataV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
