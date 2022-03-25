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

type ScalinggroupVserverGroupsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScalinggroupVserverGroupsParameters struct {

	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// +kubebuilder:validation:Required
	ScalingGroupID *string `json:"scalingGroupId" tf:"scaling_group_id,omitempty"`

	// +kubebuilder:validation:Required
	VserverGroups []VserverGroupsParameters `json:"vserverGroups" tf:"vserver_groups,omitempty"`
}

type VserverAttributesObservation struct {
}

type VserverAttributesParameters struct {

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	VserverGroupID *string `json:"vserverGroupId" tf:"vserver_group_id,omitempty"`

	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

type VserverGroupsObservation struct {
}

type VserverGroupsParameters struct {

	// +kubebuilder:validation:Required
	LoadbalancerID *string `json:"loadbalancerId" tf:"loadbalancer_id,omitempty"`

	// +kubebuilder:validation:Required
	VserverAttributes []VserverAttributesParameters `json:"vserverAttributes" tf:"vserver_attributes,omitempty"`
}

// ScalinggroupVserverGroupsSpec defines the desired state of ScalinggroupVserverGroups
type ScalinggroupVserverGroupsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScalinggroupVserverGroupsParameters `json:"forProvider"`
}

// ScalinggroupVserverGroupsStatus defines the observed state of ScalinggroupVserverGroups.
type ScalinggroupVserverGroupsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScalinggroupVserverGroupsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScalinggroupVserverGroups is the Schema for the ScalinggroupVserverGroupss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type ScalinggroupVserverGroups struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScalinggroupVserverGroupsSpec   `json:"spec"`
	Status            ScalinggroupVserverGroupsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScalinggroupVserverGroupsList contains a list of ScalinggroupVserverGroupss
type ScalinggroupVserverGroupsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalinggroupVserverGroups `json:"items"`
}

// Repository type metadata.
var (
	ScalinggroupVserverGroups_Kind             = "ScalinggroupVserverGroups"
	ScalinggroupVserverGroups_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScalinggroupVserverGroups_Kind}.String()
	ScalinggroupVserverGroups_KindAPIVersion   = ScalinggroupVserverGroups_Kind + "." + CRDGroupVersion.String()
	ScalinggroupVserverGroups_GroupVersionKind = CRDGroupVersion.WithKind(ScalinggroupVserverGroups_Kind)
)

func init() {
	SchemeBuilder.Register(&ScalinggroupVserverGroups{}, &ScalinggroupVserverGroupsList{})
}