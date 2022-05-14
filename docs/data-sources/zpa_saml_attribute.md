---
subcategory: "SAML Attributes"
layout: "zpa"
page_title: "ZPA: saml_attributes"
description: |-
  Gets a ZPA SAML Attribute details.

---

# zpa_saml_attribute

The **zpa_saml_attribute** data source provides details about a specific SAML Attributes imported into Zscaler Private Access cloud by an Identity Provider (IdP).
This data source is required when creating:

1. Access policy Rule
2. Access policy timeout rule
3. Access policy forwarding rule

## Example Usage

```hcl
# ZPA SAML Attribute Data Source
data "zpa_saml_attribute" "email_user_sso" {
    name = "Email_User SSO"
}
```

The following arguments are supported:

* `name` - (Required) The name of the saml attribute to be exported.

## Attribute Reference

* `idp_name` - (Optional)
* `idp_id` - (Optional)
* `saml_name` - (Optional)
* `user_attribute` - (Optional)