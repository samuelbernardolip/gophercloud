package rescueunrescue

import "github.com/samuelbernardolip/gophercloud"

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
