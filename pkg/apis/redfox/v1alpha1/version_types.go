package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Version struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VersionSpec   `json:"spec"`
	Status VersionStatus `json:"status,omitempty"`
}

type VersionSpec struct {
	VersionDetail VersionDetail `json:"versionDetail"`
	GitRef        VersionGitRef `json:"gitRef"`
}

type VersionDetail struct {
	DisplayVersion string `json:"displayVersion"`
	ProjectName    string `json:"projectName"`
	SubProjectName string `json:"subProjectName,omitempty"`
	BaseVersion    string `json:"baseVersion"`
	Revision       int    `json:"revision"`
}

type VersionGitRef struct {
	Branch     string `json:"branch"`
	Commit     string `json:"commit"`
	Repository string `json:"repository"`
}

type VersionStatus struct {
	Artifacts []VersionStatusArtifact `json:"artifacts,omitempty"`
}

type VersionStatusArtifact struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Uri      string `json:"uri"`
	Platform string `json:"platform"`

	// HumanFriendlyUri is optional field
	HumanFriendlyUri string `json:"humanFriendlyUri,omitempty"`
	// Description is optional field
	Description string `json:"description,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Version `json:"items"`
}
