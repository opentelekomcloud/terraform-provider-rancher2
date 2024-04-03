package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func machineConfigV2OpentelekomcloudFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"access_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Access key used for AK/SK auth",
		},
		"secret_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Access key used for AK/SK auth",
		},
		"auth_url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Authentication URL",
		},
		"availability_zone": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "OpenTelekomCloud Availability zone",
		},
		"ca_cert": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "CA certificate bundle to verify against",
		},
		"domain_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.domain_id",
				"opentelekomcloud_config.0.domain_name"},
			Description: "OpenTelekomCloud Domain ID",
		},
		"domain_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.domain_id",
				"opentelekomcloud_config.0.domain_name"},
			Description: "OpenTelekomCloud Domain Name",
		},
		"eip": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Elastic IP to use",
		},
		"eip_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "5_bgp",
			Description: "Elastic ip type (either `5_bgp` or `5_mailbgp`)",
		},
		"endpoint_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "public",
			Description: "Endpoint type",
		},
		"flavor_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Flavor id to use for the instance",
		},
		"flavor_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "s3.xlarge.2",
			Description: "Flavor name to use for the instance",
		},
		"bandwidth_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "100",
			Description: "Bandwidth size",
		},
		"bandwidth_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "PER",
			Description: "Bandwidth share type",
		},
		"image_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Image ID to use for the instance",
		},
		"image_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "Standard_Ubuntu_20.04_latest",
			Description: "Image name to use for the instance",
		},
		"ip_version": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     4,
			Description: "Version of IP address assigned for the machine (only 4 is supported by OTC for now)",
		},
		"keypair_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Key pair to use to SSH to the instance",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Computed:    true,
			Description: "OpenTelekomCloud Password",
		},
		"private_key_file": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Private key file to use for SSH (absolute path)",
		},
		"project_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.project_id",
				"opentelekomcloud_config.0.project_name"},
			Description: "OpenTelekomCloud Project ID",
		},
		"project_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.project_id",
				"opentelekomcloud_config.0.project_name"},
			Description: "OpenTelekomCloud Project name",
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Region name",
		},
		"root_volume_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     40,
			Description: "Set volume size of root partition (in GB)",
		},
		"root_volume_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "SSD",
			Description: "Set volume type of root partition (one of `SATA`, `SAS`, `SSD`)",
		},
		"security_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed:    true,
			Description: "A list of security groups to use",
		},
		"server_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Define server group where server will be created",
		},
		"server_group_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Define server group where server will be created by ID",
		},
		"skip_default_sg": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Parameter to skip default security group creation",
		},
		"skip_eip": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "If set machine IP will be set to instance local IP instead of elastic IP",
		},
		"ssh_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "22",
			Description: "Machine SSH port",
		},
		"ssh_user": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "ubuntu",
			Description: "SSH user",
		},
		"subnet_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.subnet_id",
				"opentelekomcloud_config.0.subnet_name"},
			Description: "Subnet ID of the machine",
		},
		"subnet_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.subnet_id",
				"opentelekomcloud_config.0.subnet_name"},
			Description: "Subnet name of the machine",
		},
		"token": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Authorization token that can be used instead of AK/SK or user/password auth",
		},
		"tags": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "OTC Tags (e.g. key1.value1,key2.value2,key3)",
		},
		"user_data_file": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "File containing an userdata script",
		},
		"user_data_raw": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Contents of user data file as a string",
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "OpenTelekomCloud username",
		},
		"vpc_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.vpc_id",
				"opentelekomcloud_config.0.vpc_name"},
			Description: "VPC ID of the machine",
		},
		"vpc_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ExactlyOneOf: []string{"opentelekomcloud_config.0.vpc_id",
				"opentelekomcloud_config.0.vpc_name"},
			Description: "VPC Name of the machine",
		},
	}
	return s
}
