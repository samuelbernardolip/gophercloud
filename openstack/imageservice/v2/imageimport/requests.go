package imageimport

import "github.com/samuelbernardolip/gophercloud"

// ImportMethod represents valid Import API method.
type ImportMethod string

const (
	// GlanceDirectMethod represents glance-direct Import API method.
	GlanceDirectMethod ImportMethod = "glance-direct"

	// WebDownloadMethod represents web-download Import API method.
	WebDownloadMethod ImportMethod = "web-download"
)

// Get retrieves Import API information data.
func Get(c *gophercloud.ServiceClient) (r GetResult) {
	_, r.Err = c.Get(infoURL(c), &r.Body, nil)
	return
}
