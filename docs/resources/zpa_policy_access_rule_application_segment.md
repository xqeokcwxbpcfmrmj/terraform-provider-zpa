---
subcategory: "Policy Set Controller"
layout: "zscaler"
page_title: "ZPA: policy_access_rule"
description: |-
  Creates and manages ZPA Policy Access Rule with Application Segment conditions.
---

# Resource: zpa_policy_access_rule

The **zpa_policy_access_rule** resource creates and manages a policy access rule with Application Segment conditions in the Zscaler Private Access cloud.

## Example Usage

```hcl
data "zpa_policy_type" "access_policy" {
  policy_type = "ACCESS_POLICY"
}

resource "zpa_policy_access_rule" "this" {
  name                          = "Example"
  description                   = "Example"
  action                        = "ALLOW"
  operator                      = "AND"
  policy_set_id                 = data.zpa_policy_type.access_policy.id

  conditions {
    negated  = false
    operator = "OR"
    operands {
      name =  "Example"
      object_type = "APP"
      lhs = "id"
      rhs = data.zpa_application_segment.this.id
    }
  }
}

resource "zpa_application_segment" "this" {
  name             = "Example"
  description      = "Example"
  enabled          = true
  health_reporting = "ON_ACCESS"
  bypass_type      = "NEVER"
  is_cname_enabled = true
  tcp_port_ranges  = ["80", "80"]
  domain_names     = ["crm.example.com"]
  segment_group_id = zpa_segment_group.this.id
  server_groups {
    id = [zpa_server_group.this.id]
  }
}

resource "zpa_segment_group" "this" {
  name            = "Example"
  description     = "Example"
  enabled         = true
}

// Retrieve App Connector Group
data "zpa_app_connector_group" "this" {
  name = "Example"
}

// Create Server Group
resource "zpa_server_group" "this" {
  name              = "Example"
  description       = "Example"
  enabled           = true
  dynamic_discovery = true
  app_connector_groups {
    id = [data.zpa_app_connector_group.this.id]
  }
}
```

### Required

* `name` - (Required) This is the name of the policy rule.
* `policy_set_id` - (Required) Use [zpa_policy_type](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_policy_type) data source to retrieve the necessary policy Set ID ``policy_set_id``

## Attributes Reference

* `action` (Optional) This is for providing the rule action. Supported values: ``ALLOW``, ``DENY``
* `custom_msg` (String) This is for providing a customer message for the user.
* `description` (String) This is the description of the access policy rule.
* `operator` (Optional) Supported values: ``AND``, ``OR``
* `policy_type` (Optional) Supported values: ``ACCESS_POLICY`` or ``GLOBAL_POLICY``
* `rule_order` (Optional)

* `conditions` - (Optional)
  * `negated` - (Optional) Supported values: ``true`` or ``false``
  * `operator` (Optional) Supported values: ``AND``, and ``OR``
  * `operands` (Optional) - Operands block must be repeated if multiple per `object_type` conditions are to be added to the rule.
    * `name` (Optional)
    * `object_type` (Optional) This is for specifying the policy critiera. Supported values: `APP`. Use [zpa_application_segment](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) resource or data source to associate or retrieve the Application Segment ``id`` attribute .
    * `lhs` (Optional) LHS must always carry the string value ``id``.
    * `rhs` (Optional) This is the ``id`` value of the application segment resource.

* `app_connector_groups`
  * `id` - (Optional) The ID of an app connector group resource

* `app_server_groups`
  * `id` - (Optional) The ID of a server group resource

## Import

Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
[Visit](https://github.com/zscaler/zscaler-terraformer)

Policy Access Rule for Browser Access can be imported by using`<POLICY ACCESS RULE ID>` as the import ID.

For example:

```shell
terraform import zpa_policy_access_rule.example <policy_access_rule_id>
```

## LHS and RHS Values

LHS and RHS values differ based on object types. Refer to the following table:

| Object Type | LHS| RHS
|----------|-----------|----------
| [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> |
| [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> |
| [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter |
| [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> |
| [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> |
| [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> |
| [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid>  | "true" / "false" |
| [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id>  | "true" |
| [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id>  | <Attribute_value_to_match> |
| [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id>  | <Attribute_value_to_match>  |
| [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id>  | <Attribute_value_to_match>  |
