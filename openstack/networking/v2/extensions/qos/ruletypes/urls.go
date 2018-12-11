package ruletypes

import "github.com/samuelbernardolip/gophercloud"

func listRuleTypesURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("qos", "rule-types")
}
