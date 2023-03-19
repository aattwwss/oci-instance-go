package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.validate()
	if err != nil {
		log.Fatal(err)
	}

	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal(err)
	}

	coreClient, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal(err)
	}

	if len(cfg.AvailabilityDomains) == 0 {
		cfg.AvailabilityDomains, err = ListAvailabilityDomains(identityClient, cfg.TenancyID)
		if err != nil {
			log.Fatal(err)
		}
	}

	instances := ListInstances(coreClient, cfg.TenancyID)
	existingInstances := checkExistingInstances(cfg, instances)
	if existingInstances != "" {
		log.Println(existingInstances)
		return
	}

	for _, domain := range cfg.AvailabilityDomains {
		resp, err := createInstance(coreClient, cfg, domain)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	}
}
func ListAvailabilityDomains(identityClient identity.IdentityClient, compartmentId string) ([]string, error) {

	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	req := identity.ListAvailabilityDomainsRequest{CompartmentId: common.String(compartmentId)}

	resp, err := identityClient.ListAvailabilityDomains(context.Background(), req)
	helpers.FatalIfError(err)

	var domainNames []string
	for _, item := range resp.Items {
		domainNames = append(domainNames, *item.Name)
	}
	return domainNames, nil
}

func ListInstances(client core.ComputeClient, compartmentId string) []core.Instance {
	req := core.ListInstancesRequest{Page: common.String(""),
		Limit:         common.Int(78),
		SortBy:        core.ListInstancesSortByTimecreated,
		SortOrder:     core.ListInstancesSortOrderAsc,
		CompartmentId: common.String(compartmentId)}

	// Send the request using the service client
	resp, err := client.ListInstances(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	return resp.Items
}

func checkExistingInstances(cfg config, instances []core.Instance) string {
	shape := cfg.Shape
	maxInstances := cfg.MaxInstances
	var displayNames []string
	var states []core.InstanceLifecycleStateEnum
	for _, instance := range instances {
		if *instance.Shape == shape && instance.LifecycleState != core.InstanceLifecycleStateTerminated {
			displayNames = append(displayNames, *instance.DisplayName)
			states = append(states, instance.LifecycleState)
		}
	}

	if len(displayNames) < maxInstances {
		return ""
	}

	msg := fmt.Sprintf("Already have an instance(s) %v in state(s) (respectively) %v. User: %v\n", displayNames, states, cfg.UserID)
	return msg
}

func createInstance(client core.ComputeClient, cfg config, domain string) (core.LaunchInstanceResponse, error) {
	req := core.LaunchInstanceRequest{
		LaunchInstanceDetails: core.LaunchInstanceDetails{
			Metadata:           map[string]string{"ssh_authorized_keys": cfg.SSHPublicKey},
			Shape:              &cfg.Shape,
			CompartmentId:      &cfg.TenancyID,
			DisplayName:        common.String("instance-" + time.Now().Format("20060102-1504")),
			AvailabilityDomain: &domain,
			SourceDetails:      buildSourceDetails(cfg),
			CreateVnicDetails: &core.CreateVnicDetails{
				AssignPublicIp:         common.Bool(false),
				SubnetId:               &cfg.SubnetID,
				AssignPrivateDnsRecord: common.Bool(true),
			},
			AgentConfig: &core.LaunchInstanceAgentConfigDetails{
				PluginsConfig: []core.InstanceAgentPluginConfigDetails{
					{
						Name:         common.String("Compute Instance Monitoring"),
						DesiredState: "ENABLED",
					},
				},
				IsMonitoringDisabled: common.Bool(false),
				IsManagementDisabled: common.Bool(false),
			},
			DefinedTags:  make(map[string]map[string]interface{}),
			FreeformTags: make(map[string]string),
			InstanceOptions: &core.InstanceOptions{
				AreLegacyImdsEndpointsDisabled: common.Bool(false),
			},
			AvailabilityConfig: &core.LaunchInstanceAvailabilityConfigDetails{
				RecoveryAction: core.LaunchInstanceAvailabilityConfigDetailsRecoveryActionRestoreInstance,
			},
			ShapeConfig: &core.LaunchInstanceShapeConfigDetails{
				Ocpus:       &cfg.OCPUS,
				MemoryInGBs: &cfg.MemoryInGbs,
			},
		},
	}
	return client.LaunchInstance(context.Background(), req)
}
