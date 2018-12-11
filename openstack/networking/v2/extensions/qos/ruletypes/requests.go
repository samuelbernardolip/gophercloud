package ruletypes

import (
	"github.com/samuelbernardolip/gophercloud"
	"github.com/samuelbernardolip/gophercloud/pagination"
)

// ListRuleTypes returns the list of rule types from the server
func ListRuleTypes(c *gophercloud.ServiceClient) (result pagination.Pager) {
	return pagination.NewPager(c, listRuleTypesURL(c), func(r pagination.PageResult) pagination.Page {
		return ListRuleTypesPage{pagination.SinglePageBase(r)}
	})
}
