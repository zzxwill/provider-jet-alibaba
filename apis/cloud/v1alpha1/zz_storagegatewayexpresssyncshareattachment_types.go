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

type StorageGatewayExpressSyncShareAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StorageGatewayExpressSyncShareAttachmentParameters struct {

	// +kubebuilder:validation:Required
	ExpressSyncID *string `json:"expressSyncId" tf:"express_sync_id,omitempty"`

	// +kubebuilder:validation:Required
	GatewayID *string `json:"gatewayId" tf:"gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	ShareName *string `json:"shareName" tf:"share_name,omitempty"`
}

// StorageGatewayExpressSyncShareAttachmentSpec defines the desired state of StorageGatewayExpressSyncShareAttachment
type StorageGatewayExpressSyncShareAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageGatewayExpressSyncShareAttachmentParameters `json:"forProvider"`
}

// StorageGatewayExpressSyncShareAttachmentStatus defines the observed state of StorageGatewayExpressSyncShareAttachment.
type StorageGatewayExpressSyncShareAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageGatewayExpressSyncShareAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayExpressSyncShareAttachment is the Schema for the StorageGatewayExpressSyncShareAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type StorageGatewayExpressSyncShareAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageGatewayExpressSyncShareAttachmentSpec   `json:"spec"`
	Status            StorageGatewayExpressSyncShareAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayExpressSyncShareAttachmentList contains a list of StorageGatewayExpressSyncShareAttachments
type StorageGatewayExpressSyncShareAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageGatewayExpressSyncShareAttachment `json:"items"`
}

// Repository type metadata.
var (
	StorageGatewayExpressSyncShareAttachment_Kind             = "StorageGatewayExpressSyncShareAttachment"
	StorageGatewayExpressSyncShareAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageGatewayExpressSyncShareAttachment_Kind}.String()
	StorageGatewayExpressSyncShareAttachment_KindAPIVersion   = StorageGatewayExpressSyncShareAttachment_Kind + "." + CRDGroupVersion.String()
	StorageGatewayExpressSyncShareAttachment_GroupVersionKind = CRDGroupVersion.WithKind(StorageGatewayExpressSyncShareAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageGatewayExpressSyncShareAttachment{}, &StorageGatewayExpressSyncShareAttachmentList{})
}