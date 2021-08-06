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

// MonitoringConfigurationSpec defines the desired state of MonitoringConfiguration
type MonitoringConfigurationSpec struct {
	//Defines the apps to be installed
	Apps []App `json:"spec,apps"`
	//Name of the monitoring configuration
	Name string `json:"name"`
	//Last modifier of the custom resource
	LastModifier string `json:"last_modifier"`
	//Metadata of the MonitoringConfiguration
	Metadata string `json:"metadata,omitempty"`
}

type App struct {
	//Name of the application
	Name string `json:"name"`
	//Repo from where chart needs to be downloaded
	Repo string `json:"repo"`
	//Chart name which needs to be installed
	ChartName string `json:"chartName"`
	//Version of the chart
	ChartVersion string `json:"chartVersion"`
	//Namespace of installation
	Namespace string `json:"namespace"`
}

// MonitoringConfigurationStatus defines the observed state of MonitoringConfiguration
type MonitoringConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MonitoringConfiguration is the Schema for the monitoringconfigurations API
type MonitoringConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringConfigurationSpec   `json:"spec,omitempty"`
	Status MonitoringConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MonitoringConfigurationList contains a list of MonitoringConfiguration
type MonitoringConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringConfiguration{}, &MonitoringConfigurationList{})
}
