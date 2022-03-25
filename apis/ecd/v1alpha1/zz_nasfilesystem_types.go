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

type NasFileSystemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NasFileSystemParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// +kubebuilder:validation:Optional
	MountTargetDomain *string `json:"mountTargetDomain,omitempty" tf:"mount_target_domain,omitempty"`

	// +kubebuilder:validation:Optional
	NasFileSystemName *string `json:"nasFileSystemName,omitempty" tf:"nas_file_system_name,omitempty"`

	// +kubebuilder:validation:Required
	OfficeSiteID *string `json:"officeSiteId" tf:"office_site_id,omitempty"`

	// +kubebuilder:validation:Optional
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`
}

// NasFileSystemSpec defines the desired state of NasFileSystem
type NasFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NasFileSystemParameters `json:"forProvider"`
}

// NasFileSystemStatus defines the observed state of NasFileSystem.
type NasFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NasFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NasFileSystem is the Schema for the NasFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type NasFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NasFileSystemSpec   `json:"spec"`
	Status            NasFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NasFileSystemList contains a list of NasFileSystems
type NasFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NasFileSystem `json:"items"`
}

// Repository type metadata.
var (
	NasFileSystem_Kind             = "NasFileSystem"
	NasFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NasFileSystem_Kind}.String()
	NasFileSystem_KindAPIVersion   = NasFileSystem_Kind + "." + CRDGroupVersion.String()
	NasFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(NasFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&NasFileSystem{}, &NasFileSystemList{})
}