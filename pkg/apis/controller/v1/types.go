package v1

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// YData is the specification of an ydata resource
type YData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   YDataSpec   `json:"spec"`
	Status YDataStatus `json:"status"`
}

// YDataSpec is the spec for an YData resource
type YDataSpec struct {
	MLModelName    string `json:"mlModel"`
	PipelineName   string `json:"pipeline"`
	ExperimentName string `json:"experiment"`
	ShouldRun      *bool  `json:"run"`
}

// YDataStatus is the status of an YData resource
type YDataStatus struct {
	PipelineStep string `json:"pipelineStep"`
	FooValue     int16  `json:"fooValue"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// YDataList is a list of YData resources
type YDataList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Items []YData `json:"items"`
}
