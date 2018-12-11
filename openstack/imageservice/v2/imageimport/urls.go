package imageimport

import "github.com/samuelbernardolip/gophercloud"

const (
	infoPath     = "info"
	resourcePath = "import"
)

func infoURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(infoPath, resourcePath)
}
