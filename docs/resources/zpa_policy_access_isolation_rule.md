---
subcategory: "Policy Set Controller"
layout: "zscaler"
page_title: "ZPA: policy_isolation_rule"
description: |-
  Creates and manages ZPA Policy Access Isolation Rule.
---

# Resource: zpa_policy_isolation_rule

The **zpa_policy_isolation_rule** resource creates a policy isolation access rule in the Zscaler Private Access cloud.

## Example Usage

```hcl
# Get Isolation Policy ID
data "zpa_policy_type" "isolation_policy" {
    policy_type = "ISOLATION_POLICY"
}

# Get Isolation Profile ID
data "zpa_isolation_profile" "isolation_profile" {
    name = "zpa_isolation_profile"
}

#Create Client Isolation Access Rule
resource "zpa_policy_isolation_rule" "this" {
  name                          = "Example_Isolation_Policy"
  description                   = "Example_Isolation_Policy"
  action                        = "ISOLATE"
  rule_order                     = 1
  operator = "AND"
  policy_set_id = data.zpa_policy_type.isolation_policy.id
  zpn_isolation_profile_id = data.zpa_isolation_profile.isolation_profile.id

  conditions {
    negated = false
    operator = "OR"
    operands {
      object_type = "CLIENT_TYPE"
      lhs = "id"
      rhs = "zpn_client_type_exporter"
    }
  }
}
```

### Required

* `name` - (Required) This is the name of the forwarding policy rule.
* `policy_set_id` - (Required) Use [zpa_policy_type](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_policy_type) data source to retrieve the necessary policy Set ID ``policy_set_id``
* `zpn_isolation_profile_id` - (Required) Use [zpa_isolation_profile](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_isolation_profile) data source to retrieve the necessary Isolation profile ID ``zpn_isolation_profile_id``

## Attributes Reference

* `action` - (Optional) This is for providing the rule action.
  * The supported actions for a policy isolation rule are: ``BYPASS_ISOLATE``, or ``ISOLATE``
* `description` - (Optional) This is the description of the access policy rule.
* `operator` (Optional) Supported values: ``AND``, ``OR``
* `policy_type` (Optional) Supported values: ``ISOLATION_POLICY``
* `rule_order` - (Optional)

* `conditions` - (Optional)
  * `negated` - (Optional) Supported values: ``true`` or ``false``
  * `operator` (Optional) Supported values: ``AND``, and ``OR``
  * `operands` (Optional) - Operands block must be repeated if multiple per `object_type` conditions are to be added to the rule.
    * `name` (Optional)
    * `lhs` (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
    * `rhs` (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
    * `idp_id` (Optional)
    * `object_type` (Optional) This is for specifying the policy critiera. Supported values: `APP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
    * `CLIENT_TYPE` (Optional) - The below options are the only ones supported in a timeout policy rule.
      * ``zpn_client_type_exporter`` "Web Browser"

## Import

Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language.
[Visit](https://github.com/zscaler/zscaler-terraformer)

Policy Access Isolation Rule can be imported by using `<POLICY ISOLATION RULE ID>` as the import ID.

For example:

```shell
terraform import zpa_policy_isolation_rule.example <policy_isolation_rule_id>
```

## LHS and RHS Values

LHS and RHS values differ based on object types. Refer to the following table:

| Object Type | LHS| RHS
|----------|-----------|----------
| [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> |
| [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | ``zpn_client_type_exporter`` |
| [PLATFORM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | ``linux``, ``android``, ``windows``, ``ios``, ``mac`` |
| [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> |
| [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> |
| [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> |
| [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid>  | "true" / "false" |
| [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id>  | "true" |
| [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id>  | <Attribute_value_to_match> |
| [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id>  | <Attribute_value_to_match>  |
| [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id>  | <Attribute_value_to_match>  |
