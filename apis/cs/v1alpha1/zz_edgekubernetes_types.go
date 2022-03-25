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

type AddonsObservation struct {
}

type AddonsParameters struct {

	// +kubebuilder:validation:Optional
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EdgeKubernetesObservation struct {
	CertificateAuthority map[string]*string `json:"certificateAuthority,omitempty" tf:"certificate_authority,omitempty"`

	Connections map[string]*string `json:"connections,omitempty" tf:"connections,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	SlbInternet *string `json:"slbInternet,omitempty" tf:"slb_internet,omitempty"`

	SlbIntranet *string `json:"slbIntranet,omitempty" tf:"slb_intranet,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	WorkerNodes []WorkerNodesObservation `json:"workerNodes,omitempty" tf:"worker_nodes,omitempty"`
}

type EdgeKubernetesParameters struct {

	// +kubebuilder:validation:Optional
	Addons []AddonsParameters `json:"addons,omitempty" tf:"addons,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCert *string `json:"clientCert,omitempty" tf:"client_cert,omitempty"`

	// +kubebuilder:validation:Optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterCACert *string `json:"clusterCaCert,omitempty" tf:"cluster_ca_cert,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`

	// +kubebuilder:validation:Optional
	InstallCloudMonitor *bool `json:"installCloudMonitor,omitempty" tf:"install_cloud_monitor,omitempty"`

	// +kubebuilder:validation:Optional
	IsEnterpriseSecurityGroup *bool `json:"isEnterpriseSecurityGroup,omitempty" tf:"is_enterprise_security_group,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Optional
	KubeConfig *string `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`

	// +kubebuilder:validation:Optional
	LogConfig []LogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// +kubebuilder:validation:Optional
	NewNATGateway *bool `json:"newNatGateway,omitempty" tf:"new_nat_gateway,omitempty"`

	// +kubebuilder:validation:Optional
	NodeCidrMask *float64 `json:"nodeCidrMask,omitempty" tf:"node_cidr_mask,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PodCidr *string `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyMode *string `json:"proxyMode,omitempty" tf:"proxy_mode,omitempty"`

	// +kubebuilder:validation:Optional
	RDSInstances []*string `json:"rdsInstances,omitempty" tf:"rds_instances,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	RetainResources []*string `json:"retainResources,omitempty" tf:"retain_resources,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceCidr *string `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	SlbInternetEnabled *bool `json:"slbInternetEnabled,omitempty" tf:"slb_internet_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Optional
	WorkerDataDisks []WorkerDataDisksParameters `json:"workerDataDisks,omitempty" tf:"worker_data_disks,omitempty"`

	// +kubebuilder:validation:Optional
	WorkerDiskCategory *string `json:"workerDiskCategory,omitempty" tf:"worker_disk_category,omitempty"`

	// +kubebuilder:validation:Optional
	WorkerDiskPerformanceLevel *string `json:"workerDiskPerformanceLevel,omitempty" tf:"worker_disk_performance_level,omitempty"`

	// +kubebuilder:validation:Optional
	WorkerDiskSize *float64 `json:"workerDiskSize,omitempty" tf:"worker_disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	WorkerDiskSnapshotPolicyID *string `json:"workerDiskSnapshotPolicyId,omitempty" tf:"worker_disk_snapshot_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	WorkerInstanceChargeType *string `json:"workerInstanceChargeType,omitempty" tf:"worker_instance_charge_type,omitempty"`

	// +kubebuilder:validation:Required
	WorkerInstanceTypes []*string `json:"workerInstanceTypes" tf:"worker_instance_types,omitempty"`

	// +kubebuilder:validation:Required
	WorkerNumber *float64 `json:"workerNumber" tf:"worker_number,omitempty"`

	// +kubebuilder:validation:Required
	WorkerVswitchIds []*string `json:"workerVswitchIds" tf:"worker_vswitch_ids,omitempty"`
}

type LogConfigObservation struct {
}

type LogConfigParameters struct {

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type WorkerDataDisksObservation struct {
}

type WorkerDataDisksParameters struct {

	// +kubebuilder:validation:Optional
	AutoSnapshotPolicyID *string `json:"autoSnapshotPolicyId,omitempty" tf:"auto_snapshot_policy_id,omitempty"`

	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// +kubebuilder:validation:Optional
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PerformanceLevel *string `json:"performanceLevel,omitempty" tf:"performance_level,omitempty"`

	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`
}

type WorkerNodesObservation struct {
}

type WorkerNodesParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PrivateIP *string `json:"privateIp" tf:"private_ip,omitempty"`
}

// EdgeKubernetesSpec defines the desired state of EdgeKubernetes
type EdgeKubernetesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EdgeKubernetesParameters `json:"forProvider"`
}

// EdgeKubernetesStatus defines the observed state of EdgeKubernetes.
type EdgeKubernetesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EdgeKubernetesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeKubernetes is the Schema for the EdgeKubernetess API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alicloudjet}
type EdgeKubernetes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeKubernetesSpec   `json:"spec"`
	Status            EdgeKubernetesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeKubernetesList contains a list of EdgeKubernetess
type EdgeKubernetesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeKubernetes `json:"items"`
}

// Repository type metadata.
var (
	EdgeKubernetes_Kind             = "EdgeKubernetes"
	EdgeKubernetes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EdgeKubernetes_Kind}.String()
	EdgeKubernetes_KindAPIVersion   = EdgeKubernetes_Kind + "." + CRDGroupVersion.String()
	EdgeKubernetes_GroupVersionKind = CRDGroupVersion.WithKind(EdgeKubernetes_Kind)
)

func init() {
	SchemeBuilder.Register(&EdgeKubernetes{}, &EdgeKubernetesList{})
}