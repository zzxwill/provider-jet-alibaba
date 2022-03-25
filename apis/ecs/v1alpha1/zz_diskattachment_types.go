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

type DiskAttachmentObservation struct {
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiskAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	Bootable *bool `json:"bootable,omitempty" tf:"bootable,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty" tf:"delete_with_instance,omitempty"`

	// +kubebuilder:validation:Required
	DiskID *string `json:"diskId" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`
}

// DiskAttachmentSpec defines the desired state of DiskAttachment
type DiskAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskAttachmentParameters `json:"forProvider"`
}

// DiskAttachmentStatus defines the observed state of DiskAttachment.
type DiskAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskAttachment is the Schema for the DiskAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DiskAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskAttachmentSpec   `json:"spec"`
	Status            DiskAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskAttachmentList contains a list of DiskAttachments
type DiskAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskAttachment `json:"items"`
}

// Repository type metadata.
var (
	DiskAttachment_Kind             = "DiskAttachment"
	DiskAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskAttachment_Kind}.String()
	DiskAttachment_KindAPIVersion   = DiskAttachment_Kind + "." + CRDGroupVersion.String()
	DiskAttachment_GroupVersionKind = CRDGroupVersion.WithKind(DiskAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskAttachment{}, &DiskAttachmentList{})
}