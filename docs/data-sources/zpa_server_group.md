---
subcategory: "Server Groups"
layout: "zpa"
page_title: "ZPA: server_group"
description: |-
  Gets a ZPA Server Group details.

---

# zpa_server_group

The **zpa_server_group` data source provides details about a specific Server Group created in the Zscaler Private Access.
This data source is required when creating:

1. Application Segment
2. Application Servers

## Example Usage

```hcl
# ZPA Server Group Data Source
data "zpa_server_group" "example" {
 name = "server_group_name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) This field defines the name of the server group.

### Read-Only

* `config_space` - (String)
* `description` - (String) This field is the description of the server group.
* `dynamic_discovery` - (Boolean) This field controls dynamic discovery of the servers.
* `enabled` - (Boolean) This field defines if the server group is enabled or disabled.
* `ip_anchored` - (Boolean)

`app_connector_groups` (List of Objects) This field is a json array of app-connector-id only.

* `id`  - (String) Accepts values in the format: `"id":"<appConnectorGrpId>"`