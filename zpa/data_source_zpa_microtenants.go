package zpa

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zscaler/zscaler-sdk-go/zpa/services/microtenants"
)

func dataSourceMicroTenants() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMicroTenantsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"criteria_attribute": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"criteria_attribute_values": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operator": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_role": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"user": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.",
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"customer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.",
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"eula": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"force_pwd_change": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_ids": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_locked": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"language_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_login_disabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pin_session": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"microtenant_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"microtenant_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timezone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tmp_password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"token_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This field is mandatory if twoFactorAuthEnabled is set.",
						},
						"two_factor_auth_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"two_factor_auth_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Mandatory only for POST. Not mandatory for PUT/DELETE requests.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.",
						},
						"modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.",
						},
						"modified_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.",
						},
					},
				},
			},
		},
	}
}

func dataSourceMicroTenantsRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var resp *microtenants.MicroTenant
	id, ok := d.Get("id").(string)
	if ok && id != "" {
		log.Printf("[INFO] Getting data for microtenant %s\n", id)
		res, _, err := zClient.microtenants.Get(id)
		if err != nil {
			return err
		}
		resp = res
	}
	name, ok := d.Get("name").(string)
	if ok && name != "" {
		log.Printf("[INFO] Getting data for microtenant name %s\n", name)
		res, _, err := zClient.microtenants.GetByName(name)
		if err != nil {
			return err
		}
		resp = res
	}
	if resp != nil {
		d.SetId(resp.ID)
		_ = d.Set("name", resp.Name)
		_ = d.Set("description", resp.Description)
		_ = d.Set("enabled", resp.Enabled)
		_ = d.Set("criteria_attribute", resp.CriteriaAttribute)
		_ = d.Set("criteria_attribute_values", resp.CriteriaAttributeValues)
		_ = d.Set("operator", resp.Operator)
		_ = d.Set("priority", resp.Priority)
		_ = d.Set("creation_time", resp.CreationTime)
		_ = d.Set("modified_by", resp.ModifiedBy)
		_ = d.Set("modified_time", resp.ModifiedTime)
		_ = d.Set("roles", flattenRoles(resp))

		if err := d.Set("user", flattenUserResource(resp.UserResource)); err != nil {
			return fmt.Errorf("failed to read user resource %s", err)
		}

	} else {
		return fmt.Errorf("couldn't find any microtenant with name '%s' or id '%s'", name, id)
	}

	return nil
}

func flattenRoles(role *microtenants.MicroTenant) []interface{} {
	roles := make([]interface{}, len(role.Roles))
	for i, roleItem := range role.Roles {
		roles[i] = map[string]interface{}{
			"id":          roleItem.ID,
			"name":        roleItem.Name,
			"custom_role": roleItem.CustomRole,
		}
	}

	return roles
}

func flattenUserResource(user *microtenants.UserResource) []interface{} {
	if user == nil {
		return nil
	}
	m := map[string]interface{}{
		"id":                      user.ID,
		"name":                    user.Name,
		"description":             user.Description,
		"comments":                user.Comments,
		"customer_id":             user.CustomerID,
		"display_name":            user.DisplayName,
		"email":                   user.Email,
		"eula":                    user.Eula,
		"force_pwd_change":        user.ForcePwdChange,
		"group_ids":               user.GroupIDs,
		"is_enabled":              user.IsEnabled,
		"is_locked":               user.IsLocked,
		"language_code":           user.LanguageCode,
		"local_login_disabled":    user.LocalLoginDisabled,
		"password":                user.Password,
		"phone_number":            user.PhoneNumber,
		"pin_session":             user.PinSession,
		"role_id":                 user.RoleID,
		"microtenant_id":          user.MicrotenantID,
		"microtenant_name":        user.MicrotenantName,
		"timezone":                user.Timezone,
		"tmp_password":            user.TmpPassword,
		"token_id":                user.TokenID,
		"two_factor_auth_enabled": user.TwoFactorAuthEnabled,
		"two_factor_auth_type":    user.TwoFactorAuthType,
		"username":                user.Username,
		"creation_time":           user.CreationTime,
		"modified_by":             user.ModifiedBy,
		"modified_time":           user.ModifiedTime,
	}

	return []interface{}{m}
}
