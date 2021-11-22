---
subcategory: "Server Group"
layout: "zpa"
page_title: "ZPA: server_group"
description: |-
  Creates a ZPA Server Group details.
  
---

# zpa_server_group

The **zpa_server_group** resource creates a server group in the Zscaler Private Access cloud. This resource can then be referenced in an application segment resource.

## Example Usage

```hcl
# ZPA Server Group resource - Dynamic Discover Enabled
resource "zpa_server_group" "example" {
  name = "Example"
  description = "Example"
  enabled = false
  dynamic_discovery = false
  app_connector_groups {
    id = [data.zpa_app_connector_group.example.id]
  }
}
```

```hcl
# ZPA Server Group resource - Dynamic Discover Disabled
resource "zpa_server_group" "example" {
  name = "Example"
  description = "Example"
  enabled = false
  dynamic_discovery = false
  app_connector_groups {
    id = [data.zpa_app_connector_group.example.id]
  }
  servers {
    id = [zpa_application_server.example.id]
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) This field defines the name of the server group.
* `description` - (Optional) This field is the description of the server group.
* `dynamic_discovery` - (Optional) This field controls dynamic discovery of the servers.
* `app_connector_groups` - (Required) This field is an array of app-connector-id only.
* `config_space` - (Optional)
* `enabled` - (Optional) This field defines if the server group is enabled or disabled.
* `ip_anchored` - (Optional)
* `servers` - (Optional) This field becomes required when `dynamic_discovery` is disabled.