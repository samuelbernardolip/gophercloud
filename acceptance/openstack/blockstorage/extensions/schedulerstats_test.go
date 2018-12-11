// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/samuelbernardolip/gophercloud/acceptance/clients"
	"github.com/samuelbernardolip/gophercloud/acceptance/tools"
	"github.com/samuelbernardolip/gophercloud/openstack/blockstorage/extensions/schedulerstats"
	th "github.com/samuelbernardolip/gophercloud/testhelper"
)

func TestSchedulerStatsList(t *testing.T) {
	clients.RequireAdmin(t)

	blockClient, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	listOpts := schedulerstats.ListOpts{
		Detail: true,
	}

	allPages, err := schedulerstats.List(blockClient, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allStats, err := schedulerstats.ExtractStoragePools(allPages)
	th.AssertNoErr(t, err)

	for _, stat := range allStats {
		tools.PrintResource(t, stat)
	}
}
