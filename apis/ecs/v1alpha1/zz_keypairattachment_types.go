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

type KeyPairAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KeyPairAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// +kubebuilder:validation:Required
	InstanceIds []*string `json:"instanceIds" tf:"instance_ids,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`
}

// KeyPairAttachmentSpec defines the desired state of KeyPairAttachment
type KeyPairAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyPairAttachmentParameters `json:"forProvider"`
}

// KeyPairAttachmentStatus defines the observed state of KeyPairAttachment.
type KeyPairAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyPairAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPairAttachment is the Schema for the KeyPairAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type KeyPairAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyPairAttachmentSpec   `json:"spec"`
	Status            KeyPairAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPairAttachmentList contains a list of KeyPairAttachments
type KeyPairAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyPairAttachment `json:"items"`
}

// Repository type metadata.
var (
	KeyPairAttachment_Kind             = "KeyPairAttachment"
	KeyPairAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyPairAttachment_Kind}.String()
	KeyPairAttachment_KindAPIVersion   = KeyPairAttachment_Kind + "." + CRDGroupVersion.String()
	KeyPairAttachment_GroupVersionKind = CRDGroupVersion.WithKind(KeyPairAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyPairAttachment{}, &KeyPairAttachmentList{})
}