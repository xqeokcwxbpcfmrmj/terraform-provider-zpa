---
subcategory: "Policy Set Controller"
layout: "zscaler"
page_title: "ZPA: policy_inspection_rule"
description: |-
  Creates and manages ZPA Policy Access Inspection Rule.
---

# Resource: zpa_policy_inspection_rule

The **zpa_policy_inspection_rule** resource creates a policy inspection access rule in the Zscaler Private Access cloud.

## Example Usage 1

```hcl
# Retrieve Inspection policy type
data "zpa_policy_type" "inspection_policy" {
  policy_type = "INSPECTION_POLICY"
}

#Create Inspection Access Rule
resource "zpa_policy_inspection_rule" "this" {
  name                      = "Example"
  description               = "Example"
  action                    = "INSPECT"
  rule_order                = 1
  operator                  = "AND"
  policy_set_id             = data.zpa_policy_type.inspection_policy.id
  zpn_inspection_profile_id = zpa_inspection_profile.this.id
  conditions {
    operator = "OR"
    operands {
      object_type = "APP"
      lhs         = "id"
      rhs         = zpa_application_segment_inspection.this.id
    }
  }
}
```

## Example Usage 2

```hcl
# Retrieve Inspection policy type
data "zpa_policy_type" "inspection_policy" {
  policy_type = "INSPECTION_POLICY"
}

#Create Inspection Access Rule
resource "zpa_policy_inspection_rule" "this" {
  name                      = "Example"
  description               = "Example"
  action                    = "BYPASS_INSPECT"
  rule_order                = 1
  operator                  = "AND"
  policy_set_id             = data.zpa_policy_type.inspection_policy.id
  conditions {
    operator = "OR"
    operands {
      object_type = "APP"
      lhs         = "id"
      rhs         = zpa_application_segment_inspection.this.id
    }
  }
}
```

### Required

* `name` - (Required) This is the name of the policy inspection rule.
* `policy_set_id` - (Required) Use [zpa_policy_type](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_policy_type) data source to retrieve the necessary policy Set ID ``policy_set_id``

## Attributes Reference

* `action` - (Optional) This is for providing the rule action.
  * The supported actions for a policy inspection rule are: `BYPASS_INSPECT`, or `INSPECT`
* `zpn_inspection_profile_id` (Optional) An inspection profile is required if the `action` is set to `INSPECT`
* `action_id` - (Optional) This field defines the description of the server.
* `bypass_default_rule` - (Optional)
* `custom_msg` - (Optional) This is for providing a customer message for the user.
* `description` - (Optional) This is the description of the access policy rule.
* `operator` (Optional) Supported values: ``AND``, ``OR``
* `policy_type` (Optional) Supported values: ``INSPECTION_POLICY``
* `rule_order` - (Optional)

* `conditions` - (Optional)
  * `negated` - (Optional) Supported values: ``true`` or ``false``
  * `operator` (Optional) Supported values: ``AND``, and ``OR``
  * `operands` (Optional) - Operands block must be repeated if multiple per `object_type` conditions are to be added to the rule.
    * `name` (Optional)
    * `lhs` (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
    * `rhs` (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
    * `idp_id` (Optional)
    * `object_type` (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
    * `CLIENT_TYPE` (Optional) - The below options are the only ones supported in a timeout policy rule.
      * `zpn_client_type_exporter`
      * `zpn_client_type_browser_isolation`
      * `zpn_client_type_machine_tunnel`
      * `zpn_client_type_ip_anchoring`
      * `zpn_client_type_edge_connector`
      * `zpn_client_type_zapp`

## Import

Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
[Visit](https://github.com/zscaler/zscaler-terraformer)

Policy Access Inspection Rule can be imported by using `<POLICY INSPECTION RULE ID>` as the import ID.

For example:

```shell
terraform import zpa_policy_inspection_rule.example <policy_inspection_rule_id>
```

## LHS and RHS Values

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
