package availabilityzones

import "github.com/samuelbernardolip/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
