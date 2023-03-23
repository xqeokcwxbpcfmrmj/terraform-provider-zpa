---
subcategory: "SAML Attributes"
layout: "zscaler"
page_title: "ZPA: saml_attributes"
description: |-
  Get information about SAML attributes from an Identity Provider (IdP) in the Zscaler Private Access cloud.
---

# Data Source: zpa_saml_attribute

Use the **zpa_saml_attribute** data source to get information about a SAML Attributes from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

## Example Usage

```hcl
# ZPA SAML Attribute Data Source
data "zpa_saml_attribute" "email_user_sso" {
    name = "Email_User SSO"
    idp_name = "idp_name"
}
```

```hcl
# ZPA SAML Attribute Data Source
data "zpa_saml_attribute" "department" {
    name = "DepartmentName_IdP_Name_Users"
    idp_name = "idp_name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the saml attribute to be exported.
* `id` - (Optional) The ID of the machine group to be exported.
* `idp_name` - (Optional) The name of the IdP corresponding to the SAML attribute.

-> **NOTE** When multiple Identity Providers (IdP) are onboarded in ZPA, the parameter ``idp_name`` is required in order to reture the attribute from the correct IdP.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `creation_time` - (string)
* `idp_id` - (string) The ID of the IdP corresponding to the SAML attribute.
* `name` - (string)
* `modified_by` (string)
* `modified_time` (string)
* `saml_name` - (string)
* `user_attribute` - (string)
