// +build acceptance networking

package v2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samuelbernardolip/gophercloud/acceptance/clients"
	subnetpools "github.com/samuelbernardolip/gophercloud/acceptance/openstack/networking/v2/extensions/subnetpools"
	"github.com/samuelbernardolip/gophercloud/acceptance/tools"
	"github.com/samuelbernardolip/gophercloud/openstack/networking/v2/subnets"
	th "github.com/samuelbernardolip/gophercloud/testhelper"
)

func TestSubnetCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create Network
	network, err := CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer DeleteNetwork(t, client, network.ID)

	// Create Subnet
	subnet, err := CreateSubnet(t, client, network.ID)
	th.AssertNoErr(t, err)
	defer DeleteSubnet(t, client, subnet.ID)

	tools.PrintResource(t, subnet)

	// Update Subnet
	newSubnetName := tools.RandomString("TESTACC-", 8)
	newSubnetDescription := ""
	updateOpts := subnets.UpdateOpts{
		Name:        newSubnetName,
		Description: &newSubnetDescription,
	}
	_, err = subnets.Update(client, subnet.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	// Get subnet
	newSubnet, err := subnets.Get(client, subnet.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newSubnet)
	th.AssertEquals(t, newSubnet.Name, newSubnetName)
	th.AssertEquals(t, newSubnet.Description, newSubnetDescription)

	allPages, err := subnets.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allSubnets, err := subnets.ExtractSubnets(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, subnet := range allSubnets {
		if subnet.ID == newSubnet.ID {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestSubnetsDefaultGateway(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create Network
	network, err := CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer DeleteNetwork(t, client, network.ID)

	// Create Subnet
	subnet, err := CreateSubnetWithDefaultGateway(t, client, network.ID)
	th.AssertNoErr(t, err)
	defer DeleteSubnet(t, client, subnet.ID)

	tools.PrintResource(t, subnet)

	if subnet.GatewayIP == "" {
		t.Fatalf("A default gateway was not created.")
	}

	var noGateway = ""
	updateOpts := subnets.UpdateOpts{
		GatewayIP: &noGateway,
	}

	newSubnet, err := subnets.Update(client, subnet.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	if newSubnet.GatewayIP != "" {
		t.Fatalf("Gateway was not updated correctly")
	}
}

func TestSubnetsNoGateway(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create Network
	network, err := CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer DeleteNetwork(t, client, network.ID)

	// Create Subnet
	subnet, err := CreateSubnetWithNoGateway(t, client, network.ID)
	th.AssertNoErr(t, err)
	defer DeleteSubnet(t, client, subnet.ID)

	tools.PrintResource(t, subnet)

	if subnet.GatewayIP != "" {
		t.Fatalf("A gateway exists when it shouldn't.")
	}

	subnetParts := strings.Split(subnet.CIDR, ".")
	newGateway := fmt.Sprintf("%s.%s.%s.1", subnetParts[0], subnetParts[1], subnetParts[2])
	updateOpts := subnets.UpdateOpts{
		GatewayIP: &newGateway,
	}

	newSubnet, err := subnets.Update(client, subnet.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	if newSubnet.GatewayIP == "" {
		t.Fatalf("Gateway was not updated correctly")
	}
}

func TestSubnetsWithSubnetPool(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create Network
	network, err := CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer DeleteNetwork(t, client, network.ID)

	// Create SubnetPool
	subnetPool, err := subnetpools.CreateSubnetPool(t, client)
	th.AssertNoErr(t, err)
	defer subnetpools.DeleteSubnetPool(t, client, subnetPool.ID)

	// Create Subnet
	subnet, err := CreateSubnetWithSubnetPool(t, client, network.ID, subnetPool.ID)
	th.AssertNoErr(t, err)
	defer DeleteSubnet(t, client, subnet.ID)

	tools.PrintResource(t, subnet)

	if subnet.GatewayIP == "" {
		t.Fatalf("A subnet pool was not associated.")
	}
}

func TestSubnetsWithSubnetPoolNoCIDR(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create Network
	network, err := CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer DeleteNetwork(t, client, network.ID)

	// Create SubnetPool
	subnetPool, err := subnetpools.CreateSubnetPool(t, client)
	th.AssertNoErr(t, err)
	defer subnetpools.DeleteSubnetPool(t, client, subnetPool.ID)

	// Create Subnet
	subnet, err := CreateSubnetWithSubnetPoolNoCIDR(t, client, network.ID, subnetPool.ID)
	th.AssertNoErr(t, err)
	defer DeleteSubnet(t, client, subnet.ID)

	tools.PrintResource(t, subnet)

	if subnet.GatewayIP == "" {
		t.Fatalf("A subnet pool was not associated.")
	}
}
