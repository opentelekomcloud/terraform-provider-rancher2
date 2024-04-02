package rancher2

import (
	norman "github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	machineConfigV2OpentelekomcloudKind       = "OtcConfig"
	machineConfigV2OpentelekomcloudAPIVersion = "rke-machine-config.cattle.io/v1"
	machineConfigV2OpentelekomcloudAPIType    = "rke-machine-config.cattle.io.otcconfig"
)

//Types

type machineConfigV2Opentelekomcloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AccessKey         string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	SecretKey         string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
	AuthURL           string `json:"authUrl,omitempty" yaml:"authUrl,omitempty"`
	AvailabilityZone  string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`
	CaCert            string `json:"cacert,omitempty" yaml:"cacert,omitempty"`
	DomainID          string `json:"domainId,omitempty" yaml:"domainId,omitempty"`
	DomainName        string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
	Eip               string `json:"eip,omitempty" yaml:"eip,omitempty"`
	EipType           string `json:"eipType,omitempty" yaml:"eipType,omitempty"`
	EndpointType      string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
	FlavorID          string `json:"flavorId,omitempty" yaml:"flavorId,omitempty"`
	FlavorName        string `json:"flavorName,omitempty" yaml:"flavorName,omitempty"`
	BandwidthSize     string `json:"bandwidthSize,omitempty" yaml:"bandwidthSize,omitempty"`
	BandwidthType     string `json:"bandwidthType,omitempty" yaml:"bandwidthType,omitempty"`
	ImageID           string `json:"imageId,omitempty" yaml:"imageId,omitempty"`
	ImageName         string `json:"imageName,omitempty" yaml:"imageName,omitempty"`
	IPVersion         string `json:"ipVersion,omitempty" yaml:"ipVersion,omitempty"`
	KeypairName       string `json:"keypairName,omitempty" yaml:"keypairName,omitempty"`
	Password          string `json:"password,omitempty" yaml:"password,omitempty"`
	PrivateKeyFile    string `json:"privateKeyFile,omitempty" yaml:"privateKeyFile,omitempty"`
	ProjectID         string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	ProjectName       string `json:"projectName,omitempty" yaml:"projectName,omitempty"`
	Region            string `json:"region,omitempty" yaml:"region,omitempty"`
	RootVolumeSize    string `json:"rootVolumeSize,omitempty" yaml:"rootVolumeSize,omitempty"`
	RootVolumeType    string `json:"rootVolumeType,omitempty" yaml:"rootVolumeType,omitempty"`
	SecGroups         string `json:"secGroups,omitempty" yaml:"secGroups,omitempty"`
	ServerGroup       string `json:"serverGroup,omitempty" yaml:"serverGroup,omitempty"`
	ServerGroupID     string `json:"serverGroupID,omitempty" yaml:"serverGroupID,omitempty"`
	SkipDefaultSg     bool   `json:"skipDefaultSg,omitempty"  yaml:"skipDefaultSg,omitempty"`
	SkipEip           bool   `json:"skipEip,omitempty"  yaml:"skipEip,omitempty"`
	SSHPort           string `json:"sshPort,omitempty" yaml:"sshPort,omitempty"`
	SSHUser           string `json:"sshUser,omitempty" yaml:"sshUser,omitempty"`
	SubnetID          string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
	SubnetName        string `json:"subnetName,omitempty" yaml:"subnetName,omitempty"`
	Token             string `json:"token,omitempty"  yaml:"token,omitempty"`
	Tags              string `json:"tags,omitempty"  yaml:"tags,omitempty"`
	UserDataFile      string `json:"userDataFile,omitempty"  yaml:"userDataFile,omitempty"`
	UserDataRaw       string `json:"userDataRaw,omitempty"  yaml:"userDataRaw,omitempty"`
	Username          string `json:"username,omitempty" yaml:"username,omitempty"`
	VpcID             string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
	VpcName           string `json:"vpcName,omitempty" yaml:"vpcName,omitempty"`
}

type MachineConfigV2Opentelekomcloud struct {
	norman.Resource
	machineConfigV2Opentelekomcloud
}

// Flatteners

func flattenMachineConfigV2Opentelekomcloud(in *MachineConfigV2Opentelekomcloud) []interface{} {
	if in == nil {
		return nil
	}

	obj := make(map[string]interface{})

	obj["access_key"] = in.AccessKey
	obj["secret_key"] = in.SecretKey
	obj["auth_url"] = in.AuthURL
	obj["availability_zone"] = in.AvailabilityZone
	obj["ca_cert"] = in.CaCert
	obj["domain_id"] = in.DomainID
	obj["domain_name"] = in.DomainName
	obj["eip"] = in.Eip
	obj["eip_type"] = in.EipType
	obj["endpoint_type"] = in.EndpointType
	obj["flavor_id"] = in.FlavorID
	obj["flavor_name"] = in.FlavorName
	obj["bandwidth_size"] = in.BandwidthSize
	obj["bandwidth_type"] = in.BandwidthType
	obj["image_id"] = in.ImageID
	obj["image_name"] = in.ImageName
	obj["ip_version"] = in.IPVersion
	obj["keypair_name"] = in.KeypairName
	obj["password"] = in.Password
	obj["private_key_file"] = in.PrivateKeyFile
	obj["project_id"] = in.ProjectID
	obj["project_name"] = in.ProjectName
	obj["region"] = in.Region
	obj["root_volume_size"] = in.RootVolumeSize
	obj["root_volume_type"] = in.RootVolumeType
	obj["security_groups"] = toArraySplitComma(in.SecGroups)
	obj["server_group"] = in.ServerGroup
	obj["server_group_id"] = in.ServerGroupID
	obj["skip_default_sg"] = in.SkipDefaultSg
	obj["skip_eip"] = in.SkipEip
	obj["ssh_port"] = in.SSHPort
	obj["ssh_user"] = in.SSHUser
	obj["subnet_id"] = in.SubnetID
	obj["subnet_name"] = in.SubnetName
	obj["token"] = in.Token
	obj["tags"] = in.Tags
	obj["user_data_file"] = in.UserDataFile
	obj["user_data_raw"] = in.UserDataRaw
	obj["username"] = in.Username
	obj["vpc_id"] = in.VpcID
	obj["vpc_name"] = in.VpcName
	return []interface{}{obj}
}

// Expanders

func expandMachineConfigV2Opentelekomcloud(p []interface{}, source *MachineConfigV2) *MachineConfigV2Opentelekomcloud {
	if p == nil || len(p) == 0 || p[0] == nil {
		return nil
	}
	obj := &MachineConfigV2Opentelekomcloud{}

	if len(source.ID) > 0 {
		obj.ID = source.ID
	}
	in := p[0].(map[string]interface{})

	obj.TypeMeta.Kind = machineConfigV2OpentelekomcloudKind
	obj.TypeMeta.APIVersion = machineConfigV2OpentelekomcloudAPIVersion
	source.TypeMeta = obj.TypeMeta
	obj.ObjectMeta = source.ObjectMeta

	if v, ok := in["access_key"].(string); ok && len(v) > 0 {
		obj.AccessKey = v
	}
	if v, ok := in["secret_key"].(string); ok && len(v) > 0 {
		obj.SecretKey = v
	}
	if v, ok := in["auth_url"].(string); ok && len(v) > 0 {
		obj.AuthURL = v
	}
	if v, ok := in["availability_zone"].(string); ok && len(v) > 0 {
		obj.AvailabilityZone = v
	}
	if v, ok := in["cacert"].(string); ok && len(v) > 0 {
		obj.CaCert = v
	}
	if v, ok := in["domain_id"].(string); ok && len(v) > 0 {
		obj.DomainID = v
	}
	if v, ok := in["domain_name"].(string); ok && len(v) > 0 {
		obj.DomainName = v
	}
	if v, ok := in["eip"].(string); ok && len(v) > 0 {
		obj.Eip = v
	}
	if v, ok := in["eip_type"].(string); ok && len(v) > 0 {
		obj.EipType = v
	}
	if v, ok := in["endpoint_type"].(string); ok && len(v) > 0 {
		obj.EndpointType = v
	}
	if v, ok := in["flavor_id"].(string); ok && len(v) > 0 {
		obj.FlavorID = v
	}
	if v, ok := in["flavor_name"].(string); ok && len(v) > 0 {
		obj.FlavorName = v
	}
	if v, ok := in["bandwidth_size"].(string); ok && len(v) > 0 {
		obj.BandwidthSize = v
	}
	if v, ok := in["bandwidth_type"].(string); ok && len(v) > 0 {
		obj.BandwidthType = v
	}
	if v, ok := in["image_id"].(string); ok && len(v) > 0 {
		obj.ImageID = v
	}
	if v, ok := in["image_name"].(string); ok && len(v) > 0 {
		obj.ImageName = v
	}
	if v, ok := in["ip_version"].(string); ok && len(v) > 0 {
		obj.IPVersion = v
	}
	if v, ok := in["keypair_name"].(string); ok && len(v) > 0 {
		obj.KeypairName = v
	}
	if v, ok := in["password"].(string); ok && len(v) > 0 {
		obj.Password = v
	}
	if v, ok := in["private_key_file"].(string); ok && len(v) > 0 {
		obj.PrivateKeyFile = v
	}
	if v, ok := in["project_Id"].(string); ok && len(v) > 0 {
		obj.ProjectID = v
	}
	if v, ok := in["project_name"].(string); ok && len(v) > 0 {
		obj.ProjectName = v
	}
	if v, ok := in["region"].(string); ok && len(v) > 0 {
		obj.Region = v
	}
	if v, ok := in["root_volume_size"].(string); ok && len(v) > 0 {
		obj.RootVolumeSize = v
	}
	if v, ok := in["root_volume_type"].(string); ok && len(v) > 0 {
		obj.RootVolumeType = v
	}
	if v, ok := in["security_groups"].([]interface{}); ok && len(v) > 0 {
		obj.SecGroups = toArrayJoinComma(v)
	}
	if v, ok := in["server_group"].(string); ok && len(v) > 0 {
		obj.ServerGroup = v
	}
	if v, ok := in["server_group_id"].(string); ok && len(v) > 0 {
		obj.ServerGroupID = v
	}
	if v, ok := in["skip_default_sg"].(bool); ok {
		obj.SkipDefaultSg = v
	}
	if v, ok := in["skip_eip"].(bool); ok {
		obj.SkipEip = v
	}
	if v, ok := in["ssh_port"].(string); ok && len(v) > 0 {
		obj.SSHPort = v
	}
	if v, ok := in["ssh_user"].(string); ok && len(v) > 0 {
		obj.SSHUser = v
	}
	if v, ok := in["subnet_id"].(string); ok && len(v) > 0 {
		obj.SubnetID = v
	}
	if v, ok := in["subnet_name"].(string); ok && len(v) > 0 {
		obj.SubnetName = v
	}
	if v, ok := in["token"].(string); ok && len(v) > 0 {
		obj.Token = v
	}
	if v, ok := in["tags"].(string); ok && len(v) > 0 {
		obj.Tags = v
	}
	if v, ok := in["user_data_file"].(string); ok && len(v) > 0 {
		obj.UserDataFile = v
	}
	if v, ok := in["user_data_raw"].(string); ok && len(v) > 0 {
		obj.UserDataRaw = v
	}
	if v, ok := in["username"].(string); ok && len(v) > 0 {
		obj.Username = v
	}
	if v, ok := in["vpc_id"].(string); ok && len(v) > 0 {
		obj.VpcID = v
	}
	if v, ok := in["vpc_name"].(string); ok && len(v) > 0 {
		obj.VpcName = v
	}
	return obj
}
