package webhooks

import "github.com/samuelbernardolip/gophercloud"

func triggerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("v1", "webhooks", id, "trigger")
}
