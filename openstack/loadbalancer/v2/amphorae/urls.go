package amphorae

import "github.com/samuelbernardolip/gophercloud"

const (
	rootPath     = "octavia"
	resourcePath = "amphorae"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
