package rancher2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const (
	testAccRancher2MachineConfigV2 = "rancher2_machine_config_v2"
)

var (
	testAccRancher2MachineConfigV2Basic  string
	testAccRancher2MachineConfigV2Update string
)

func init() {
	testAccRancher2MachineConfigV2Basic = `
resource "` + testAccRancher2MachineConfigV2 + `" "foo" {
  generate_name = "test-foo"
  opentelekomcloud_config {
    auth_url          = "https://iam.eu-de.otc.t-systems.com/v3"
    availability_zone = "eu-de-01"
    username          = "test-user"
    password          = "dummy-password"
    region            = "eu-de"
    project_name      = "eu-de"
    domain_name       = "OTC-EU-DE-00000000001-TEST-DOMAIN"
    flavor_name       = "s3.medium.1"
    subnet_name       = "subnet-test"
    vpc_name          = "vpc-test"
    tags              = "key1.value1,key2.value2,key3"
    security_groups   = ["default", "sg-test"]
  }
}`
	testAccRancher2MachineConfigV2Update = `
resource "` + testAccRancher2MachineConfigV2 + `" "foo" {
  generate_name = "test-foo"
  opentelekomcloud_config {
    auth_url          = "https://iam.eu-de.otc.t-systems.com/v3"
    availability_zone = "eu-de-01"
    username          = "test-user"
    password          = "dummy-password-2"
    region            = "eu-nl"
    project_name      = "eu-de_main"
    domain_name       = "OTC-EU-DE-00000000001-TEST-DOMAIN"
    flavor_name       = "s3x.2.medium.1"
    subnet_name       = "subnet-prod"
    vpc_name          = "vpc-test"
    security_groups   = ["default"]
  }
}`
}

func TestAccRancher2MachineConfigV2Basic(t *testing.T) {
	rsName := testAccRancher2MachineConfigV2 + ".foo"
	var machineConfig *MachineConfigV2Opentelekomcloud
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRancher2MachineConfigV2Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRancher2MachineConfigV2Basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2MachineConfigV2Exists(rsName, machineConfig),
					resource.TestCheckResourceAttrSet(rsName, "name"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.region", "eu-de"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.project_name", "eu-de"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.flavor_name", "s3.medium.1"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.subnet_name", "subnet-test"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.security_groups.1", "sg-test"),
				),
			},
			{
				Config: testAccRancher2MachineConfigV2Update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2MachineConfigV2Exists(rsName, machineConfig),
					resource.TestCheckResourceAttrSet(rsName, "name"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.region", "eu-nl"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.project_name", "eu-de_main"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.flavor_name", "s3x.2.medium.1"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.subnet_name", "subnet-prod"),
					resource.TestCheckResourceAttr(rsName, "opentelekomcloud_config.0.security_groups.0", "default"),
				),
			},
		},
	})
}

func testAccCheckRancher2MachineConfigV2Exists(n string, machineConfig *MachineConfigV2Opentelekomcloud) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Machine Config ID is set")
		}

		client, _ := testAccProvider.Meta().(*Config)

		machineConfig = &MachineConfigV2Opentelekomcloud{}

		err := client.getObjectV2ByID("local", rs.Primary.ID, machineConfigV2OpentelekomcloudAPIType, machineConfig)
		if err != nil {
			if IsNotFound(err) {
				return fmt.Errorf("Machine Config V2 not found")
			}
			return err
		}

		return nil
	}
}

func testAccCheckRancher2MachineConfigV2Destroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != testAccRancher2MachineConfigV2 {
			continue
		}
		client, _ := testAccProvider.Meta().(*Config)

		resp := &MachineConfigV2Opentelekomcloud{}

		err := client.getObjectV2ByID("local", rs.Primary.ID, machineConfigV2OpentelekomcloudAPIType, resp)
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return err
		}
		return fmt.Errorf("Machine Config V2 still exists")
	}
	return nil
}
