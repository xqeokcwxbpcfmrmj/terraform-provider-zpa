---
subcategory: "Inspection"
layout: "zscaler"
page_title: "ZPA: zpa_inspection"
description: |-
  Creates and manages Inspection Profile in Zscaler Private Access cloud.
---

# Resource: zpa_inspection_profile

The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

## Example Usage

```hcl
data "zpa_inspection_profile" "this" {
  name = "Example"
}

resource "zpa_inspection_profile" "this" {
  name                          = "Example"
  description                   = "Example"
  paranoia_level                = "2"
  predefined_controls_version   = "OWASP_CRS/3.3.0"
  incarnation_number            = "6"
  custom_controls {
      id = [ "216196257331305413" ]
  }
  predefined_controls {
      id = [ "72057594037930388"]
  }
  controls_info {
    control_type = "PREDEFINED"
  }
  global_control_actions = [
          "PREDEFINED:PASS",
          "CUSTOM:NONE",
          "OVERRIDE_ACTION:COMMON"
  ]
  common_global_override_actions_config = {
          "PREDEF_CNTRL_GLOBAL_ACTION": "PASS",
          "IS_OVERRIDE_ACTION_COMMON": "TRUE"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the inspection profile.
* `description` - (Optional) Description of the inspection profile.
* `paranoia_level` - (Required) OWASP Predefined Paranoia Level. Range: [1-4], inclusive
* `predefined_controls` - (Required) The predefined controls
  * `id` - (Required) ID of the predefined control
  * `action` - (Required) The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
  * `action_value` - (Required) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
  * `attachment` (Optional) Control attachment
  * `control_group` (Optional) Control group

* `custom_controls` - (Optional) Types for custom controls
  * `type` (Optional) Types for custom controls
  * `control_rule_json` (Optional) Custom controls string in JSON format
  * `rules` - (Optional) Rules of the custom controls applied as conditions `JSON`
    * `conditions` - (Optional)
      * `lhs` - (Optional) Signifies the key for the object type Supported values: `SIZE`, `VALUE`
      * `op` - (Optional) If lhs is set to SIZE, then the user may pass one of the following: `EQ: Equals`, `LE: Less than or equal to`, `GE: Greater than or equal to`. If the lhs is set to `VALUE`, then the user may pass one of the following: `CONTAINS`, `STARTS_WITH`, `ENDS_WITH`, `RX`.
      * `rhs` - (Optional) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: `GET`,`HEAD`, `POST`, `OPTIONS`, `PUT`, `DELETE`, `TRACE`
    * `names` - (Optional) Name of the rules. If rules.type is set to `REQUEST_HEADERS`, `REQUEST_COOKIES`, or `RESPONSE_HEADERS`, the rules.name field is required.
    * `type` - (Optional) Type value for the rules
    * `version` - (Optional) The version of the predefined control, the default is: `OWASP_CRS/3.3.0`

* `associated_inspection_profile_names` - (Optional) Name of the inspection profile
  * `id`- (Optional)
  * `name`- (Optional)

* `common_global_override_actions_config` - (Optional)
* `controls_info` - (Optional) Types for custom controls
  * `control_type` - (Optional) Control types. Supported Values: `CUSTOM`, `PREDEFINED`, `ZSCALER`
  * `count` - (Optional) Control information counts `Long`

## Attributes Reference
