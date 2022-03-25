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

type K8SApplicationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type K8SApplicationParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationDescriotion *string `json:"applicationDescriotion,omitempty" tf:"application_descriotion,omitempty"`

	// +kubebuilder:validation:Required
	ApplicationName *string `json:"applicationName" tf:"application_name,omitempty"`

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	Command *string `json:"command,omitempty" tf:"command,omitempty"`

	// +kubebuilder:validation:Optional
	CommandArgs []*string `json:"commandArgs,omitempty" tf:"command_args,omitempty"`

	// +kubebuilder:validation:Optional
	EdasContainerVersion *string `json:"edasContainerVersion,omitempty" tf:"edas_container_version,omitempty"`

	// +kubebuilder:validation:Optional
	Envs map[string]*string `json:"envs,omitempty" tf:"envs,omitempty"`

	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// +kubebuilder:validation:Optional
	InternetSlbID *string `json:"internetSlbId,omitempty" tf:"internet_slb_id,omitempty"`

	// +kubebuilder:validation:Optional
	InternetSlbPort *float64 `json:"internetSlbPort,omitempty" tf:"internet_slb_port,omitempty"`

	// +kubebuilder:validation:Optional
	InternetSlbProtocol *string `json:"internetSlbProtocol,omitempty" tf:"internet_slb_protocol,omitempty"`

	// +kubebuilder:validation:Optional
	InternetTargetPort *float64 `json:"internetTargetPort,omitempty" tf:"internet_target_port,omitempty"`

	// +kubebuilder:validation:Optional
	Jdk *string `json:"jdk,omitempty" tf:"jdk,omitempty"`

	// +kubebuilder:validation:Optional
	LimitMCPU *float64 `json:"limitMCpu,omitempty" tf:"limit_m_cpu,omitempty"`

	// +kubebuilder:validation:Optional
	LimitMem *float64 `json:"limitMem,omitempty" tf:"limit_mem,omitempty"`

	// +kubebuilder:validation:Optional
	Liveness *string `json:"liveness,omitempty" tf:"liveness,omitempty"`

	// +kubebuilder:validation:Optional
	LocalVolume *string `json:"localVolume,omitempty" tf:"local_volume,omitempty"`

	// +kubebuilder:validation:Optional
	LogicalRegionID *string `json:"logicalRegionId,omitempty" tf:"logical_region_id,omitempty"`

	// +kubebuilder:validation:Optional
	MountDescs *string `json:"mountDescs,omitempty" tf:"mount_descs,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	NasID *string `json:"nasId,omitempty" tf:"nas_id,omitempty"`

	// +kubebuilder:validation:Optional
	PackageType *string `json:"packageType,omitempty" tf:"package_type,omitempty"`

	// +kubebuilder:validation:Optional
	PackageURL *string `json:"packageUrl,omitempty" tf:"package_url,omitempty"`

	// +kubebuilder:validation:Optional
	PackageVersion *string `json:"packageVersion,omitempty" tf:"package_version,omitempty"`

	// +kubebuilder:validation:Optional
	PostStart *string `json:"postStart,omitempty" tf:"post_start,omitempty"`

	// +kubebuilder:validation:Optional
	PreStop *string `json:"preStop,omitempty" tf:"pre_stop,omitempty"`

	// +kubebuilder:validation:Optional
	Readiness *string `json:"readiness,omitempty" tf:"readiness,omitempty"`

	// +kubebuilder:validation:Optional
	Replicas *float64 `json:"replicas,omitempty" tf:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	RequestsMCPU *float64 `json:"requestsMCpu,omitempty" tf:"requests_m_cpu,omitempty"`

	// +kubebuilder:validation:Optional
	RequestsMem *float64 `json:"requestsMem,omitempty" tf:"requests_mem,omitempty"`

	// +kubebuilder:validation:Optional
	WebContainer *string `json:"webContainer,omitempty" tf:"web_container,omitempty"`
}

// K8SApplicationSpec defines the desired state of K8SApplication
type K8SApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     K8SApplicationParameters `json:"forProvider"`
}

// K8SApplicationStatus defines the observed state of K8SApplication.
type K8SApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        K8SApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// K8SApplication is the Schema for the K8SApplications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type K8SApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              K8SApplicationSpec   `json:"spec"`
	Status            K8SApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8SApplicationList contains a list of K8SApplications
type K8SApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8SApplication `json:"items"`
}

// Repository type metadata.
var (
	K8SApplication_Kind             = "K8SApplication"
	K8SApplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: K8SApplication_Kind}.String()
	K8SApplication_KindAPIVersion   = K8SApplication_Kind + "." + CRDGroupVersion.String()
	K8SApplication_GroupVersionKind = CRDGroupVersion.WithKind(K8SApplication_Kind)
)

func init() {
	SchemeBuilder.Register(&K8SApplication{}, &K8SApplicationList{})
}