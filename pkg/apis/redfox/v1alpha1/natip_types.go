package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NatIp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NatIpSpec `json:"spec"`
}

// IpType string describes ip version of cidrs
// +enum
type IpType string

const (
	Ipv4 IpType = "Ipv4"
	Ipv6 IpType = "Ipv6"
)

type NatIpSpec struct {
	IpType IpType   `json:"ipType"`
	Cidrs  []string `json:"cidrs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NatIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NatIp `json:"items"`
}
