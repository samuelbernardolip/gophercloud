package buildinfo

import "github.com/samuelbernardolip/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
