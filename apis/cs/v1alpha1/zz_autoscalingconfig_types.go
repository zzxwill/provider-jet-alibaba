/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoscalingConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AutoscalingConfigParameters struct {

	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	CoolDownDuration *string `json:"coolDownDuration,omitempty" tf:"cool_down_duration,omitempty"`

	// +kubebuilder:validation:Optional
	GpuUtilizationThreshold *string `json:"gpuUtilizationThreshold,omitempty" tf:"gpu_utilization_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ScanInterval *string `json:"scanInterval,omitempty" tf:"scan_interval,omitempty"`

	// +kubebuilder:validation:Optional
	UnneededDuration *string `json:"unneededDuration,omitempty" tf:"unneeded_duration,omitempty"`

	// +kubebuilder:validation:Optional
	UtilizationThreshold *string `json:"utilizationThreshold,omitempty" tf:"utilization_threshold,omitempty"`
}

// AutoscalingConfigSpec defines the desired state of AutoscalingConfig
type AutoscalingConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoscalingConfigParameters `json:"forProvider"`
}

// AutoscalingConfigStatus defines the observed state of AutoscalingConfig.
type AutoscalingConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoscalingConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingConfig is the Schema for the AutoscalingConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type AutoscalingConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingConfigSpec   `json:"spec"`
	Status            AutoscalingConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingConfigList contains a list of AutoscalingConfigs
type AutoscalingConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingConfig `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingConfig_Kind             = "AutoscalingConfig"
	AutoscalingConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoscalingConfig_Kind}.String()
	AutoscalingConfig_KindAPIVersion   = AutoscalingConfig_Kind + "." + CRDGroupVersion.String()
	AutoscalingConfig_GroupVersionKind = CRDGroupVersion.WithKind(AutoscalingConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingConfig{}, &AutoscalingConfigList{})
}