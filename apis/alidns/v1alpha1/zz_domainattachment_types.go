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

type DomainAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainAttachmentParameters struct {

	// +kubebuilder:validation:Required
	DomainNames []*string `json:"domainNames" tf:"domain_names,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`
}

// DomainAttachmentSpec defines the desired state of DomainAttachment
type DomainAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainAttachmentParameters `json:"forProvider"`
}

// DomainAttachmentStatus defines the observed state of DomainAttachment.
type DomainAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainAttachment is the Schema for the DomainAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type DomainAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainAttachmentSpec   `json:"spec"`
	Status            DomainAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainAttachmentList contains a list of DomainAttachments
type DomainAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainAttachment `json:"items"`
}

// Repository type metadata.
var (
	DomainAttachment_Kind             = "DomainAttachment"
	DomainAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainAttachment_Kind}.String()
	DomainAttachment_KindAPIVersion   = DomainAttachment_Kind + "." + CRDGroupVersion.String()
	DomainAttachment_GroupVersionKind = CRDGroupVersion.WithKind(DomainAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainAttachment{}, &DomainAttachmentList{})
}