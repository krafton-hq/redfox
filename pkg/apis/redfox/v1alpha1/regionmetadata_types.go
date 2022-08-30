package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RegionMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RegionMetadataSpec `json:"spec"`
}

type RegionMetadataSpec struct {
	Argocd              RegionMetadataArgocd   `json:"argocd"`
	ContainerRegistries map[string]string      `json:"containerRegistries,omitempty"`
	Identity            RegionMetadataIdentity `json:"identity"`
}

type RegionMetadataArgocd struct {
	Endpoint      string `json:"endpoint"`
	ClusterRegion string `json:"clusterRegion"`
	ClusterName   string `json:"clusterName"`
}

type RegionMetadataIdentity struct {
	Name           string `json:"name"`
	InfraAccountId string `json:"infraAccountId"`
	InfraVendor    string `json:"infraVendor"`
	RegionGroup    string `json:"regionGroup"`
}
