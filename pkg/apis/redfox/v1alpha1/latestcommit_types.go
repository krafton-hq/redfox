package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LatestCommit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LatestCommitSpec   `json:"spec"`
	Status LatestCommitStatus `json:"status"`
}

type LatestCommitSpec struct {
	GitRef LatestCommitGitRef `json:"gitRef"`
}

type LatestCommitGitRef struct {
	Branch     string `json:"branch"`
	Repository string `json:"repository"`
}

type LatestCommitStatus struct {
	Commit          string      `json:"commit"`
	CommitTimestamp metav1.Time `json:"commitTimestamp"`
	SyncTimestamp   metav1.Time `json:"syncTimestamp"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LatestCommitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LatestCommit `json:"items"`
}
