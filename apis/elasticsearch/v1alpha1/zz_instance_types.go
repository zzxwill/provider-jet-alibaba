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

type InstanceObservation struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KibanaDomain *string `json:"kibanaDomain,omitempty" tf:"kibana_domain,omitempty"`

	KibanaPort *float64 `json:"kibanaPort,omitempty" tf:"kibana_port,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// +kubebuilder:validation:Optional
	ClientNodeAmount *float64 `json:"clientNodeAmount,omitempty" tf:"client_node_amount,omitempty"`

	// +kubebuilder:validation:Optional
	ClientNodeSpec *string `json:"clientNodeSpec,omitempty" tf:"client_node_spec,omitempty"`

	// +kubebuilder:validation:Required
	DataNodeAmount *float64 `json:"dataNodeAmount" tf:"data_node_amount,omitempty"`

	// +kubebuilder:validation:Optional
	DataNodeDiskEncrypted *bool `json:"dataNodeDiskEncrypted,omitempty" tf:"data_node_disk_encrypted,omitempty"`

	// +kubebuilder:validation:Required
	DataNodeDiskSize *float64 `json:"dataNodeDiskSize" tf:"data_node_disk_size,omitempty"`

	// +kubebuilder:validation:Required
	DataNodeDiskType *string `json:"dataNodeDiskType" tf:"data_node_disk_type,omitempty"`

	// +kubebuilder:validation:Required
	DataNodeSpec *string `json:"dataNodeSpec" tf:"data_node_spec,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EnableKibanaPrivateNetwork *bool `json:"enableKibanaPrivateNetwork,omitempty" tf:"enable_kibana_private_network,omitempty"`

	// +kubebuilder:validation:Optional
	EnableKibanaPublicNetwork *bool `json:"enableKibanaPublicNetwork,omitempty" tf:"enable_kibana_public_network,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePublic *bool `json:"enablePublic,omitempty" tf:"enable_public,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// +kubebuilder:validation:Optional
	KMSEncryptedPassword *string `json:"kmsEncryptedPassword,omitempty" tf:"kms_encrypted_password,omitempty"`

	// +kubebuilder:validation:Optional
	KMSEncryptionContext map[string]*string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context,omitempty"`

	// +kubebuilder:validation:Optional
	KibanaPrivateWhitelist []*string `json:"kibanaPrivateWhitelist,omitempty" tf:"kibana_private_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	KibanaWhitelist []*string `json:"kibanaWhitelist,omitempty" tf:"kibana_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	MasterNodeSpec *string `json:"masterNodeSpec,omitempty" tf:"master_node_spec,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateWhitelist []*string `json:"privateWhitelist,omitempty" tf:"private_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	PublicWhitelist []*string `json:"publicWhitelist,omitempty" tf:"public_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SettingConfig map[string]*string `json:"settingConfig,omitempty" tf:"setting_config,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`

	// +kubebuilder:validation:Required
	VswitchID *string `json:"vswitchId" tf:"vswitch_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneCount *float64 `json:"zoneCount,omitempty" tf:"zone_count,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}