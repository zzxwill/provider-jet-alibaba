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

type ExchangeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ExchangeParameters struct {

	// +kubebuilder:validation:Optional
	AlternateExchange *string `json:"alternateExchange,omitempty" tf:"alternate_exchange,omitempty"`

	// +kubebuilder:validation:Required
	AutoDeleteState *bool `json:"autoDeleteState" tf:"auto_delete_state,omitempty"`

	// +kubebuilder:validation:Required
	ExchangeName *string `json:"exchangeName" tf:"exchange_name,omitempty"`

	// +kubebuilder:validation:Required
	ExchangeType *string `json:"exchangeType" tf:"exchange_type,omitempty"`

	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Required
	Internal *bool `json:"internal" tf:"internal,omitempty"`

	// +kubebuilder:validation:Required
	VirtualHostName *string `json:"virtualHostName" tf:"virtual_host_name,omitempty"`
}

// ExchangeSpec defines the desired state of Exchange
type ExchangeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExchangeParameters `json:"forProvider"`
}

// ExchangeStatus defines the observed state of Exchange.
type ExchangeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExchangeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Exchange is the Schema for the Exchanges API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Exchange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExchangeSpec   `json:"spec"`
	Status            ExchangeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExchangeList contains a list of Exchanges
type ExchangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Exchange `json:"items"`
}

// Repository type metadata.
var (
	Exchange_Kind             = "Exchange"
	Exchange_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Exchange_Kind}.String()
	Exchange_KindAPIVersion   = Exchange_Kind + "." + CRDGroupVersion.String()
	Exchange_GroupVersionKind = CRDGroupVersion.WithKind(Exchange_Kind)
)

func init() {
	SchemeBuilder.Register(&Exchange{}, &ExchangeList{})
}