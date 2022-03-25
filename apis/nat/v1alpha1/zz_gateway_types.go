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

type BandwidthPackagesObservation struct {
	PublicIPAddresses *string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`
}

type BandwidthPackagesParameters struct {

	// +kubebuilder:validation:Required
	Bandwidth *float64 `json:"bandwidth" tf:"bandwidth,omitempty"`

	// +kubebuilder:validation:Required
	IPCount *float64 `json:"ipCount" tf:"ip_count,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type GatewayObservation struct {
	BandwidthPackageIds *string `json:"bandwidthPackageIds,omitempty" tf:"bandwidth_package_ids,omitempty"`

	ForwardTableIds *string `json:"forwardTableIds,omitempty" tf:"forward_table_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SnatTableIds *string `json:"snatTableIds,omitempty" tf:"snat_table_ids,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type GatewayParameters struct {

	// +kubebuilder:validation:Optional
	BandwidthPackages []BandwidthPackagesParameters `json:"bandwidthPackages,omitempty" tf:"bandwidth_packages,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	NATGatewayName *string `json:"natGatewayName,omitempty" tf:"nat_gateway_name,omitempty"`

	// +kubebuilder:validation:Optional
	NATType *string `json:"natType,omitempty" tf:"nat_type,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// +kubebuilder:validation:Optional
	PaymentType *string `json:"paymentType,omitempty" tf:"payment_type,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}