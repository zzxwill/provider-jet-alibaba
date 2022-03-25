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

type SynchronizationInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SynchronizationInstanceParameters struct {

	// +kubebuilder:validation:Optional
	ComputeUnit *float64 `json:"computeUnit,omitempty" tf:"compute_unit,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseCount *float64 `json:"databaseCount,omitempty" tf:"database_count,omitempty"`

	// +kubebuilder:validation:Required
	DestinationEndpointEngineName *string `json:"destinationEndpointEngineName" tf:"destination_endpoint_engine_name,omitempty"`

	// +kubebuilder:validation:Required
	DestinationEndpointRegion *string `json:"destinationEndpointRegion" tf:"destination_endpoint_region,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceClass *string `json:"instanceClass,omitempty" tf:"instance_class,omitempty"`

	// +kubebuilder:validation:Optional
	PaymentDuration *float64 `json:"paymentDuration,omitempty" tf:"payment_duration,omitempty"`

	// +kubebuilder:validation:Optional
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" tf:"payment_duration_unit,omitempty"`

	// +kubebuilder:validation:Required
	PaymentType *string `json:"paymentType" tf:"payment_type,omitempty"`

	// +kubebuilder:validation:Optional
	Quantity *float64 `json:"quantity,omitempty" tf:"quantity,omitempty"`

	// +kubebuilder:validation:Required
	SourceEndpointEngineName *string `json:"sourceEndpointEngineName" tf:"source_endpoint_engine_name,omitempty"`

	// +kubebuilder:validation:Required
	SourceEndpointRegion *string `json:"sourceEndpointRegion" tf:"source_endpoint_region,omitempty"`

	// +kubebuilder:validation:Optional
	SyncArchitecture *string `json:"syncArchitecture,omitempty" tf:"sync_architecture,omitempty"`
}

// SynchronizationInstanceSpec defines the desired state of SynchronizationInstance
type SynchronizationInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SynchronizationInstanceParameters `json:"forProvider"`
}

// SynchronizationInstanceStatus defines the observed state of SynchronizationInstance.
type SynchronizationInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SynchronizationInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SynchronizationInstance is the Schema for the SynchronizationInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type SynchronizationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SynchronizationInstanceSpec   `json:"spec"`
	Status            SynchronizationInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SynchronizationInstanceList contains a list of SynchronizationInstances
type SynchronizationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SynchronizationInstance `json:"items"`
}

// Repository type metadata.
var (
	SynchronizationInstance_Kind             = "SynchronizationInstance"
	SynchronizationInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SynchronizationInstance_Kind}.String()
	SynchronizationInstance_KindAPIVersion   = SynchronizationInstance_Kind + "." + CRDGroupVersion.String()
	SynchronizationInstance_GroupVersionKind = CRDGroupVersion.WithKind(SynchronizationInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&SynchronizationInstance{}, &SynchronizationInstanceList{})
}