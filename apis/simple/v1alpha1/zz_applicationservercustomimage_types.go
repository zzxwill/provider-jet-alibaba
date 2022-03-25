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

type ApplicationServerCustomImageObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApplicationServerCustomImageParameters struct {

	// +kubebuilder:validation:Required
	CustomImageName *string `json:"customImageName" tf:"custom_image_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Required
	SystemSnapshotID *string `json:"systemSnapshotId" tf:"system_snapshot_id,omitempty"`
}

// ApplicationServerCustomImageSpec defines the desired state of ApplicationServerCustomImage
type ApplicationServerCustomImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationServerCustomImageParameters `json:"forProvider"`
}

// ApplicationServerCustomImageStatus defines the observed state of ApplicationServerCustomImage.
type ApplicationServerCustomImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationServerCustomImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationServerCustomImage is the Schema for the ApplicationServerCustomImages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ApplicationServerCustomImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationServerCustomImageSpec   `json:"spec"`
	Status            ApplicationServerCustomImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationServerCustomImageList contains a list of ApplicationServerCustomImages
type ApplicationServerCustomImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationServerCustomImage `json:"items"`
}

// Repository type metadata.
var (
	ApplicationServerCustomImage_Kind             = "ApplicationServerCustomImage"
	ApplicationServerCustomImage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationServerCustomImage_Kind}.String()
	ApplicationServerCustomImage_KindAPIVersion   = ApplicationServerCustomImage_Kind + "." + CRDGroupVersion.String()
	ApplicationServerCustomImage_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationServerCustomImage_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationServerCustomImage{}, &ApplicationServerCustomImageList{})
}