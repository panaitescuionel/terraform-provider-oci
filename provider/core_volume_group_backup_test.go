// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	VolumeGroupBackupRequiredOnlyResource = VolumeGroupBackupResourceDependencies + `
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
	#Required
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"
}
`

	VolumeGroupBackupResourceConfig = VolumeGroupBackupResourceDependencies + `
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
	#Required
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"

	#Optional
	compartment_id = "${var.compartment_id}"
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_group_backup_defined_tags_value}")}"
	display_name = "${var.volume_group_backup_display_name}"
	freeform_tags = "${var.volume_group_backup_freeform_tags}"
	type = "${var.volume_group_backup_type}"
}
`
	VolumeGroupBackupPropertyVariables = `
variable "volume_group_backup_defined_tags_value" { default = "value" }
variable "volume_group_backup_display_name" { default = "displayName" }
variable "volume_group_backup_freeform_tags" { default = {"Department"= "Finance"} }
variable "volume_group_backup_type" { default = "INCREMENTAL" }

`
	VolumeGroupBackupResourceDependencies = VolumeGroupPropertyVariables + VolumeGroupResourceConfig
)

func TestCoreVolumeGroupBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_group_backup.test_volume_group_backup"
	datasourceName := "data.oci_core_volume_group_backups.test_volume_group_backups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeGroupBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + VolumeGroupBackupPropertyVariables + compartmentIdVariableStr + VolumeGroupBackupRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VolumeGroupBackupPropertyVariables + compartmentIdVariableStr + VolumeGroupBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
					resource.TestCheckResourceAttr(resourceName, "volume_backup_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "volume_group_backup_defined_tags_value" { default = "updatedValue" }
variable "volume_group_backup_display_name" { default = "displayName2" }
variable "volume_group_backup_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_group_backup_type" { default = "INCREMENTAL" }

                ` + compartmentIdVariableStr + VolumeGroupBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
					resource.TestCheckResourceAttr(resourceName, "volume_backup_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "volume_group_backup_defined_tags_value" { default = "updatedValue" }
variable "volume_group_backup_display_name" { default = "displayName2" }
variable "volume_group_backup_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_group_backup_type" { default = "INCREMENTAL" }

data "oci_core_volume_group_backups" "test_volume_group_backups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.volume_group_backup_display_name}"
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_volume_group_backup.test_volume_group_backup.id}"]
    }
}
                ` + compartmentIdVariableStr + VolumeGroupBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.type", "INCREMENTAL"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.volume_backup_ids.#", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.volume_group_id"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreVolumeGroupBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_group_backup" {
			noResourceFound = false
			request := oci_core.GetVolumeGroupBackupRequest{}

			tmp := rs.Primary.ID
			request.VolumeGroupBackupId = &tmp

			response, err := client.GetVolumeGroupBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeGroupBackupLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
