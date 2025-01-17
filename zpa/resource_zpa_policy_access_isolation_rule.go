package zpa

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	client "github.com/zscaler/zscaler-sdk-go/zpa"
	"github.com/zscaler/zscaler-sdk-go/zpa/services/policysetcontroller"
)

func resourcePolicyIsolationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyIsolationRuleCreate,
		Read:   resourcePolicyIsolationRuleRead,
		Update: resourcePolicyIsolationRuleUpdate,
		Delete: resourcePolicyIsolationRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importPolicyStateContextFunc([]string{"ISOLATE", "BYPASS_ISOLATE"}),
		},

		Schema: MergeSchema(
			CommonPolicySchema(),
			map[string]*schema.Schema{
				"action": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "  This is for providing the rule action.",
					ValidateFunc: validation.StringInSlice([]string{
						"ISOLATE",
						"BYPASS_ISOLATE",
					}, false),
				},
				"conditions": GetPolicyConditionsSchema([]string{
					"APP",
					"CLIENT_TYPE",
					"EDGE_CONNECTOR_GROUP",
					"POSTURE",
					"MACHINE_GRP",
					"TRUSTED_NETWORK",
					"PLATFORM",
					"IDP",
					"SAML",
					"SCIM",
					"SCIM_GROUP",
				}),
			},
		),
	}
}

func resourcePolicyIsolationRuleCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req, err := expandCreatePolicyIsolationRule(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating zpa policy isolation rule with request\n%+v\n", req)
	if ValidateConditions(req.Conditions, zClient) {
		policysetcontroller, _, err := zClient.policysetcontroller.Create(req)
		if err != nil {
			return err
		}
		d.SetId(policysetcontroller.ID)
		order, ok := d.GetOk("rule_order")
		if ok {
			reorder(order, policysetcontroller.PolicySetID, "ISOLATION_POLICY", policysetcontroller.ID, zClient)
		}
		return resourcePolicyIsolationRuleRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the zpa policy isolation (%s) operands, please make sure you are using valid inputs for APP type, LHS & RHS", req.Name)
	}

}

func resourcePolicyIsolationRuleRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	globalPolicySet, _, err := zClient.policysetcontroller.GetByPolicyType("ISOLATION_POLICY")
	if err != nil {
		return err
	}
	log.Printf("[INFO] Getting Policy Set Rule: globalPolicySet:%s id: %s\n", globalPolicySet.ID, d.Id())
	resp, _, err := zClient.policysetcontroller.GetPolicyRule(globalPolicySet.ID, d.Id())
	if err != nil {
		if obj, ok := err.(*client.ErrorResponse); ok && obj.IsObjectNotFound() {
			log.Printf("[WARN] Removing policy rule %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Got Policy Set Isolation Rule:\n%+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("action", resp.Action)
	_ = d.Set("operator", resp.Operator)
	_ = d.Set("policy_set_id", resp.PolicySetID)
	_ = d.Set("policy_type", resp.PolicyType)
	_ = d.Set("zpn_cbi_profile_id", resp.ZpnCbiProfileID)
	_ = d.Set("zpn_isolation_profile_id", resp.ZpnIsolationProfileID)
	_ = d.Set("rule_order", resp.RuleOrder)
	_ = d.Set("conditions", flattenPolicyConditions(resp.Conditions))

	return nil
}

func resourcePolicyIsolationRuleUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicySet, _, err := zClient.policysetcontroller.GetByPolicyType("ISOLATION_POLICY")
	if err != nil {
		return err
	}
	ruleID := d.Id()
	log.Printf("[INFO] Updating policy isolation rule ID: %v\n", ruleID)
	req, err := expandCreatePolicyRule(d)
	if err != nil {
		return err
	}
	if ValidateConditions(req.Conditions, zClient) {
		if _, _, err := zClient.policysetcontroller.GetPolicyRule(globalPolicySet.ID, ruleID); err != nil {
			if respErr, ok := err.(*client.ErrorResponse); ok && respErr.IsObjectNotFound() {
				d.SetId("")
				return nil
			}
		}

		if _, err := zClient.policysetcontroller.Update(globalPolicySet.ID, ruleID, req); err != nil {
			return err
		}
		if d.HasChange("rule_order") {
			order, ok := d.GetOk("rule_order")
			if ok {
				reorder(order, globalPolicySet.ID, "ISOLATION_POLICY", ruleID, zClient)
			}
		}
		return resourcePolicyIsolationRuleRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the zpa policy isolation (%s) operands, please make sure you are using valid inputs for APP type, LHS & RHS", req.Name)
	}

}

func resourcePolicyIsolationRuleDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicySet, _, err := zClient.policysetcontroller.GetByPolicyType("ISOLATION_POLICY")
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting policy isolation rule with id %v\n", d.Id())

	if _, err := zClient.policysetcontroller.Delete(globalPolicySet.ID, d.Id()); err != nil {
		return err
	}

	return nil

}

func expandCreatePolicyIsolationRule(d *schema.ResourceData) (*policysetcontroller.PolicyRule, error) {
	policySetID, ok := d.Get("policy_set_id").(string)
	if !ok {
		return nil, fmt.Errorf("policy_set_id is not set")
	}
	log.Printf("[INFO] action_id:%v\n", d.Get("action_id"))
	conditions, err := ExpandPolicyConditions(d)
	if err != nil {
		return nil, err
	}
	return &policysetcontroller.PolicyRule{
		ID:                    d.Get("id").(string),
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		Action:                d.Get("action").(string),
		ActionID:              d.Get("action_id").(string),
		CustomMsg:             d.Get("custom_msg").(string),
		BypassDefaultRule:     d.Get("bypass_default_rule").(bool),
		DefaultRule:           d.Get("default_rule").(bool),
		Operator:              d.Get("operator").(string),
		PolicySetID:           policySetID,
		PolicyType:            d.Get("policy_type").(string),
		ZpnCbiProfileID:       d.Get("zpn_cbi_profile_id").(string),
		ZpnIsolationProfileID: d.Get("zpn_isolation_profile_id").(string),
		Priority:              d.Get("priority").(string),
		RuleOrder:             d.Get("rule_order").(string),
		Conditions:            conditions,
	}, nil
}
