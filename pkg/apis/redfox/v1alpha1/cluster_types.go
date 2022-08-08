package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status"`
}

type ClusterSpec struct {
	ClusterName    string `json:"clusterName"`
	ClusterRegion  string `json:"clusterRegion"`
	ClusterGroup   string `json:"clusterGroup"`
	ServicePhase   string `json:"servicePhase"`
	ServiceTag     string `json:"serviceTag"`
	ClusterType    string `json:"clusterType"`
	InfraVendor    string `json:"infraVendor"`
	InfraAccountId string `json:"infraAccountId"`
	ClusterEngine  string `json:"clusterEngine"`
}

type ClusterStatus struct {
	Apiserver            ClusterStatusApiserver `json:"apiserver"`
	ServiceAccountIssuer string                 `json:"serviceAccountIssuer"`
	AwsIamIdps           map[string]string      `json:"awsIamIdps"`
}

type ClusterStatusApiserver struct {
	Endpoint string `json:"endpoint"`
	CaCert   string `json:"caCert"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Cluster `json:"items"`
}
