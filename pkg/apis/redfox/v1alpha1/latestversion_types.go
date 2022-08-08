package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LatestVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LatestVersionSpec   `json:"spec"`
	Status LatestVersionStatus `json:"status"`
}

type LatestVersionSpec struct {
	GitRef LatestVersionGitRef `json:"gitRef"`
}

type LatestVersionGitRef struct {
	Branch     string `json:"branch"`
	Repository string `json:"repository"`
}

type LatestVersionStatus struct {
	VersionRef LatestVersionVersionRef `json:"versionRef"`
}

type LatestVersionVersionRef struct {
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LatestVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LatestVersion `json:"items"`
}
