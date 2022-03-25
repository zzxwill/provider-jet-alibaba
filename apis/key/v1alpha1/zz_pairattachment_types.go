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

type PairAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PairAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// +kubebuilder:validation:Required
	InstanceIds []*string `json:"instanceIds" tf:"instance_ids,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`
}

// PairAttachmentSpec defines the desired state of PairAttachment
type PairAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PairAttachmentParameters `json:"forProvider"`
}

// PairAttachmentStatus defines the observed state of PairAttachment.
type PairAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PairAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PairAttachment is the Schema for the PairAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type PairAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PairAttachmentSpec   `json:"spec"`
	Status            PairAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PairAttachmentList contains a list of PairAttachments
type PairAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PairAttachment `json:"items"`
}

// Repository type metadata.
var (
	PairAttachment_Kind             = "PairAttachment"
	PairAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PairAttachment_Kind}.String()
	PairAttachment_KindAPIVersion   = PairAttachment_Kind + "." + CRDGroupVersion.String()
	PairAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PairAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PairAttachment{}, &PairAttachmentList{})
}