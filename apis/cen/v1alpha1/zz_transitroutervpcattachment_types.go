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

type TransitRouterVPCAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TransitRouterAttachmentID *string `json:"transitRouterAttachmentId,omitempty" tf:"transit_router_attachment_id,omitempty"`
}

type TransitRouterVPCAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	AutoCreateVPCRoute *bool `json:"autoCreateVpcRoute,omitempty" tf:"auto_create_vpc_route,omitempty"`

	// +kubebuilder:validation:Required
	CenID *string `json:"cenId" tf:"cen_id,omitempty"`

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTableAssociationEnabled *bool `json:"routeTableAssociationEnabled,omitempty" tf:"route_table_association_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTablePropagationEnabled *bool `json:"routeTablePropagationEnabled,omitempty" tf:"route_table_propagation_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TransitRouterAttachmentDescription *string `json:"transitRouterAttachmentDescription,omitempty" tf:"transit_router_attachment_description,omitempty"`

	// +kubebuilder:validation:Optional
	TransitRouterAttachmentName *string `json:"transitRouterAttachmentName,omitempty" tf:"transit_router_attachment_name,omitempty"`

	// +kubebuilder:validation:Optional
	TransitRouterID *string `json:"transitRouterId,omitempty" tf:"transit_router_id,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCOwnerID *string `json:"vpcOwnerId,omitempty" tf:"vpc_owner_id,omitempty"`

	// +kubebuilder:validation:Required
	ZoneMappings []ZoneMappingsParameters `json:"zoneMappings" tf:"zone_mappings,omitempty"`
}

type ZoneMappingsObservation struct {
}

type ZoneMappingsParameters struct {

	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// TransitRouterVPCAttachmentSpec defines the desired state of TransitRouterVPCAttachment
type TransitRouterVPCAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitRouterVPCAttachmentParameters `json:"forProvider"`
}

// TransitRouterVPCAttachmentStatus defines the observed state of TransitRouterVPCAttachment.
type TransitRouterVPCAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitRouterVPCAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitRouterVPCAttachment is the Schema for the TransitRouterVPCAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type TransitRouterVPCAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitRouterVPCAttachmentSpec   `json:"spec"`
	Status            TransitRouterVPCAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitRouterVPCAttachmentList contains a list of TransitRouterVPCAttachments
type TransitRouterVPCAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitRouterVPCAttachment `json:"items"`
}

// Repository type metadata.
var (
	TransitRouterVPCAttachment_Kind             = "TransitRouterVPCAttachment"
	TransitRouterVPCAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitRouterVPCAttachment_Kind}.String()
	TransitRouterVPCAttachment_KindAPIVersion   = TransitRouterVPCAttachment_Kind + "." + CRDGroupVersion.String()
	TransitRouterVPCAttachment_GroupVersionKind = CRDGroupVersion.WithKind(TransitRouterVPCAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitRouterVPCAttachment{}, &TransitRouterVPCAttachmentList{})
}