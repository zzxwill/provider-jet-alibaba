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

type VPCEndpointServiceResourceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCEndpointServiceResourceParameters struct {

	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Required
	ServiceID *string `json:"serviceId" tf:"service_id,omitempty"`
}

// VPCEndpointServiceResourceSpec defines the desired state of VPCEndpointServiceResource
type VPCEndpointServiceResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointServiceResourceParameters `json:"forProvider"`
}

// VPCEndpointServiceResourceStatus defines the observed state of VPCEndpointServiceResource.
type VPCEndpointServiceResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointServiceResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceResource is the Schema for the VPCEndpointServiceResources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type VPCEndpointServiceResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointServiceResourceSpec   `json:"spec"`
	Status            VPCEndpointServiceResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceResourceList contains a list of VPCEndpointServiceResources
type VPCEndpointServiceResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointServiceResource `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointServiceResource_Kind             = "VPCEndpointServiceResource"
	VPCEndpointServiceResource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointServiceResource_Kind}.String()
	VPCEndpointServiceResource_KindAPIVersion   = VPCEndpointServiceResource_Kind + "." + CRDGroupVersion.String()
	VPCEndpointServiceResource_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointServiceResource_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointServiceResource{}, &VPCEndpointServiceResourceList{})
}