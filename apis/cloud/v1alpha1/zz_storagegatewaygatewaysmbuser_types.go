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

type StorageGatewayGatewaySMBUserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StorageGatewayGatewaySMBUserParameters struct {

	// +kubebuilder:validation:Required
	GatewayID *string `json:"gatewayId" tf:"gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// StorageGatewayGatewaySMBUserSpec defines the desired state of StorageGatewayGatewaySMBUser
type StorageGatewayGatewaySMBUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageGatewayGatewaySMBUserParameters `json:"forProvider"`
}

// StorageGatewayGatewaySMBUserStatus defines the observed state of StorageGatewayGatewaySMBUser.
type StorageGatewayGatewaySMBUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageGatewayGatewaySMBUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayGatewaySMBUser is the Schema for the StorageGatewayGatewaySMBUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type StorageGatewayGatewaySMBUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageGatewayGatewaySMBUserSpec   `json:"spec"`
	Status            StorageGatewayGatewaySMBUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageGatewayGatewaySMBUserList contains a list of StorageGatewayGatewaySMBUsers
type StorageGatewayGatewaySMBUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageGatewayGatewaySMBUser `json:"items"`
}

// Repository type metadata.
var (
	StorageGatewayGatewaySMBUser_Kind             = "StorageGatewayGatewaySMBUser"
	StorageGatewayGatewaySMBUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageGatewayGatewaySMBUser_Kind}.String()
	StorageGatewayGatewaySMBUser_KindAPIVersion   = StorageGatewayGatewaySMBUser_Kind + "." + CRDGroupVersion.String()
	StorageGatewayGatewaySMBUser_GroupVersionKind = CRDGroupVersion.WithKind(StorageGatewayGatewaySMBUser_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageGatewayGatewaySMBUser{}, &StorageGatewayGatewaySMBUserList{})
}