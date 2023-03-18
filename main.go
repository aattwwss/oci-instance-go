package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func main() {
	domains := ListAvailabilityDomains()
	instances := ListInstances(domains[0].Name)
	b, _ := json.Marshal(instances)
	fmt.Println(string(b))
}
func ListAvailabilityDomains() []identity.AvailabilityDomain {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := identity.ListAvailabilityDomainsRequest{CompartmentId: common.String("ocid1.tenancy.oc1..aaaaaaaai7czu4a2llxchv7veudgj7cbg5fi3rerivgbmf2h7q4wrg54i37q")}

	// Send the request using the service client
	resp, err := client.ListAvailabilityDomains(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	return resp.Items
}

func ListInstances(availabilityDomain *string) []core.Instance {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := core.ListInstancesRequest{Page: common.String(""),
		AvailabilityDomain: availabilityDomain,
		Limit:              common.Int(78),
		SortBy:             core.ListInstancesSortByTimecreated,
		SortOrder:          core.ListInstancesSortOrderAsc,
		CompartmentId:      common.String("ocid1.tenancy.oc1..aaaaaaaai7czu4a2llxchv7veudgj7cbg5fi3rerivgbmf2h7q4wrg54i37q")}

	// Send the request using the service client
	resp, err := client.ListInstances(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	return resp.Items
}
