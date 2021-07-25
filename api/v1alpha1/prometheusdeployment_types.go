/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PrometheusDeploymentSpec defines the desired state of PrometheusDeployment
type PrometheusDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name string `json:"name"`
	//Defines the Prometheus chart details
	Prometheus Prometheus `json:"prometheus,omitempty"`
	//Last modifier of the custom resource
	LastModifier string `json:"last_modifier"`
	//Metadata of the prometheusdeployment
	Metadata string `json:"metadata,omitempty"`
}

type Prometheus struct {
	//Specifies if the prometheus chart is to be installed. Assumed to be true if not specified
	Enabled bool `json:"enabled",omitempty"`
	//Version of the chart to be installed
	Version string `json:"version,omitempty"`
	//Specifies the repo from where prometheus chart is to be used
	Repo string `json:"repo"`
}

// PrometheusDeploymentStatus defines the observed state of PrometheusDeployment
type PrometheusDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PrometheusDeployment is the Schema for the prometheusdeployments API
type PrometheusDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrometheusDeploymentSpec   `json:"spec,omitempty"`
	Status PrometheusDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PrometheusDeploymentList contains a list of PrometheusDeployment
type PrometheusDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrometheusDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrometheusDeployment{}, &PrometheusDeploymentList{})
}
