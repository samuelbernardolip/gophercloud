package apiversions

import "github.com/samuelbernardolip/gophercloud"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
