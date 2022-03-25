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

type HostGroupAccountUserGroupAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostGroupAccountUserGroupAttachmentParameters struct {

	// +kubebuilder:validation:Required
	HostAccountNames []*string `json:"hostAccountNames" tf:"host_account_names,omitempty"`

	// +kubebuilder:validation:Required
	HostGroupID *string `json:"hostGroupId" tf:"host_group_id,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Required
	UserGroupID *string `json:"userGroupId" tf:"user_group_id,omitempty"`
}

// HostGroupAccountUserGroupAttachmentSpec defines the desired state of HostGroupAccountUserGroupAttachment
type HostGroupAccountUserGroupAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostGroupAccountUserGroupAttachmentParameters `json:"forProvider"`
}

// HostGroupAccountUserGroupAttachmentStatus defines the observed state of HostGroupAccountUserGroupAttachment.
type HostGroupAccountUserGroupAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostGroupAccountUserGroupAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostGroupAccountUserGroupAttachment is the Schema for the HostGroupAccountUserGroupAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type HostGroupAccountUserGroupAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostGroupAccountUserGroupAttachmentSpec   `json:"spec"`
	Status            HostGroupAccountUserGroupAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostGroupAccountUserGroupAttachmentList contains a list of HostGroupAccountUserGroupAttachments
type HostGroupAccountUserGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostGroupAccountUserGroupAttachment `json:"items"`
}

// Repository type metadata.
var (
	HostGroupAccountUserGroupAttachment_Kind             = "HostGroupAccountUserGroupAttachment"
	HostGroupAccountUserGroupAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostGroupAccountUserGroupAttachment_Kind}.String()
	HostGroupAccountUserGroupAttachment_KindAPIVersion   = HostGroupAccountUserGroupAttachment_Kind + "." + CRDGroupVersion.String()
	HostGroupAccountUserGroupAttachment_GroupVersionKind = CRDGroupVersion.WithKind(HostGroupAccountUserGroupAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&HostGroupAccountUserGroupAttachment{}, &HostGroupAccountUserGroupAttachmentList{})
}