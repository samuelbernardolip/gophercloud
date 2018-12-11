// +build acceptance imageservice imageimport

package v2

import (
	"testing"

	"github.com/samuelbernardolip/gophercloud/acceptance/clients"
	"github.com/samuelbernardolip/gophercloud/acceptance/tools"
	th "github.com/samuelbernardolip/gophercloud/testhelper"
)

func TestGetImportInfo(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	importInfo, err := GetImportInfo(t, client)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, importInfo)
}
