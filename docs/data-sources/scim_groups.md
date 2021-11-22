---
subcategory: "SCIM Groups"
layout: "zpa"
page_title: "ZPA: scim_groups"
description: |-
  Gets a ZPA SCIM Group details.
  
---

# zpa_scim_groups

The **zpa_scim_groups** data source provides details about a specific SCIM Group imported into Zscaler Private Access cloud by an Identity Provider (IdP).
This data source is required when creating:

1. Access policy Rule
2. Access policy timeout rule
3. Access policy forwarding rule

## Example Usage

```hcl
# ZPA SCIM Groups Data Source
data "zpa_scim_groups" "engineering" {
    name = "Engineering"
    idp_name = "idp_name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name. The name of the scim group to be exported.
* `idp_name` - (Required) Name. The name of the IdP where the scim group must be exported from.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `idp_group_id`(Optional)
* `idp_id` (Optional)