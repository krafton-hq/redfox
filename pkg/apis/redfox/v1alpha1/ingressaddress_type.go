package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec IngressAddressSpec `json:"spec"`
}

type IngressAddressSpec struct {
	Default IngressAddressSpecDefault `json:"default"`
	Specs   IngressAddressSpecs       `json:"specs"`
}

type IngressAddressSpecDefault struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type IngressAddressSpecs struct {
	Hosts     []string `json:"hosts"`
	Ports     []int    `json:"ports"`
	Protocols []string `json:"protocols"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []IngressAddress `json:"items"`
}
