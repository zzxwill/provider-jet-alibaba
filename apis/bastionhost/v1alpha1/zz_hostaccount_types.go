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

type HostAccountObservation struct {
	HostAccountID *string `json:"hostAccountId,omitempty" tf:"host_account_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostAccountParameters struct {

	// +kubebuilder:validation:Required
	HostAccountName *string `json:"hostAccountName" tf:"host_account_name,omitempty"`

	// +kubebuilder:validation:Required
	HostID *string `json:"hostId" tf:"host_id,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	PassPhraseSecretRef *v1.SecretKeySelector `json:"passPhraseSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ProtocolName *string `json:"protocolName" tf:"protocol_name,omitempty"`
}

// HostAccountSpec defines the desired state of HostAccount
type HostAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostAccountParameters `json:"forProvider"`
}

// HostAccountStatus defines the observed state of HostAccount.
type HostAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostAccount is the Schema for the HostAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type HostAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostAccountSpec   `json:"spec"`
	Status            HostAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostAccountList contains a list of HostAccounts
type HostAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostAccount `json:"items"`
}

// Repository type metadata.
var (
	HostAccount_Kind             = "HostAccount"
	HostAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostAccount_Kind}.String()
	HostAccount_KindAPIVersion   = HostAccount_Kind + "." + CRDGroupVersion.String()
	HostAccount_GroupVersionKind = CRDGroupVersion.WithKind(HostAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&HostAccount{}, &HostAccountList{})
}