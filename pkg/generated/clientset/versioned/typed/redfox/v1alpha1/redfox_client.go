// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/krafton-hq/redfox/pkg/apis/redfox/v1alpha1"
	"github.com/krafton-hq/redfox/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type MetadataV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
	LatestVersionsGetter
	NatIpsGetter
	RegionMetadatasGetter
	VersionsGetter
	VersionCountsGetter
}

// MetadataV1alpha1Client is used to interact with features provided by the metadata.sbx-central.io group.
type MetadataV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MetadataV1alpha1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

func (c *MetadataV1alpha1Client) LatestVersions(namespace string) LatestVersionInterface {
	return newLatestVersions(c, namespace)
}

func (c *MetadataV1alpha1Client) NatIps(namespace string) NatIpInterface {
	return newNatIps(c, namespace)
}

func (c *MetadataV1alpha1Client) RegionMetadatas(namespace string) RegionMetadataInterface {
	return newRegionMetadatas(c, namespace)
}

func (c *MetadataV1alpha1Client) Versions(namespace string) VersionInterface {
	return newVersions(c, namespace)
}

func (c *MetadataV1alpha1Client) VersionCounts(namespace string) VersionCountInterface {
	return newVersionCounts(c, namespace)
}

// NewForConfig creates a new MetadataV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*MetadataV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new MetadataV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*MetadataV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &MetadataV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MetadataV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MetadataV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MetadataV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MetadataV1alpha1Client {
	return &MetadataV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MetadataV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
