---
layout: "oci"
page_title: "OCI: oci_database_autonomous_data_warehouse"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouse"
description: |-
  Provides details about a specific AutonomousDataWarehouse
---

# Data Source: oci_database_autonomous_data_warehouse
The `oci_database_autonomous_data_warehouse` data source provides details about a specific AutonomousDataWarehouse

Gets the details of the specified Autonomous Data Warehouse.


## Example Usage

```hcl
data "oci_database_autonomous_data_warehouse" "test_autonomous_data_warehouse" {
	#Required
	autonomous_data_warehouse_id = "${var.autonomous_data_warehouse_autonomous_data_warehouse_id}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_warehouse_id` - (Required) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `connection_strings` - The connection string used to connect to the Data Warehouse. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Data Warehouse for the password value.
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
* `cpu_core_count` - The number of CPU cores to be made available to the database.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `db_name` - The database name.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Data Warehouse. The default is BRING_YOUR_OWN_LICENSE. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `service_console_url` - The URL of the Service Console for the Data Warehouse.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.

