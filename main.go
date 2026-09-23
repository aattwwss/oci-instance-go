package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func main() {
	var t int
	flag.IntVar(&t, "t", 0, "Number of minutes between retries")
	flag.Parse()

	if t == 0 {
		run()
		return
	}

	log.Printf("Starting script with %v minutes delay.", t)
	for true {
		run()
		time.Sleep(time.Duration(t) * time.Minute)
	}
}

func run() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.validate()
	if err != nil {
		log.Fatal(err)
	}

	cp, err := cfg.buildConfigProvider()
	if err != nil {
		log.Fatal(err)
	}

	cfg.TenancyID, err = cp.TenancyOCID()
	if err != nil {
		log.Fatal(err)
	}

	coreClient, err := core.NewComputeClientWithConfigurationProvider(cp)
	if err != nil {
		log.Fatal(err)
	}

	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(cp)
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

	// Names of live instances, used to pick the next free display name.
	takenNames := make(map[string]bool)
	for _, i := range instances {
		if i.LifecycleState != core.InstanceLifecycleStateTerminated {
			takenNames[*i.DisplayName] = true
		}
	}

	for _, domain := range cfg.AvailabilityDomains {
		log.Println("Trying domain: ", domain)
		_, err := createInstance(coreClient, cfg, domain, takenNames)
		if err == nil {
			handleSuccess()
			return
		}
		if !strings.Contains(err.Error(), "Out of host capacity") {
			log.Println("Something went wrong: ", err)
			return
		}
		log.Println("Domain out of capacity: ", domain)
	}
}
func ListAvailabilityDomains(client identity.IdentityClient, compartmentId string) ([]string, error) {
	req := identity.ListAvailabilityDomainsRequest{CompartmentId: common.String(compartmentId)}

	resp, err := client.ListAvailabilityDomains(context.Background(), req)
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

	msg := fmt.Sprintf("Already have an instance(s) %v in state(s) (respectively) %v\n", displayNames, states)
	return msg
}

func createInstance(client core.ComputeClient, cfg config, domain string, takenNames map[string]bool) (core.LaunchInstanceResponse, error) {
	req := core.LaunchInstanceRequest{
		LaunchInstanceDetails: core.LaunchInstanceDetails{
			Metadata:           map[string]string{"ssh_authorized_keys": cfg.SSHPublicKey},
			Shape:              &cfg.Shape,
			CompartmentId:      &cfg.TenancyID,
			DisplayName:        common.String(displayName(cfg, takenNames)),
			AvailabilityDomain: &domain,
			SourceDetails:      buildSourceDetails(cfg),
			CreateVnicDetails: &core.CreateVnicDetails{
				AssignPublicIp:         common.Bool(cfg.AssignPublicIP),
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
		},
	}
	// In-transit encryption for paravirtualized boot/data volume attachments.
	req.LaunchInstanceDetails.IsPvEncryptionInTransitEnabled = &cfg.PvEncryptionInTransit
	// Fixed shapes like VM.Standard.E2.1.Micro have built-in OCPU/memory;
	// sending a zeroed ShapeConfig causes InternalError from the API.
	if cfg.OCPUS > 0 || cfg.MemoryInGbs > 0 {
		req.LaunchInstanceDetails.ShapeConfig = &core.LaunchInstanceShapeConfigDetails{
			Ocpus:       &cfg.OCPUS,
			MemoryInGBs: &cfg.MemoryInGbs,
		}
	}
	return client.LaunchInstance(context.Background(), req)
}

func handleSuccess() {
	log.Println("Instance created")
}

func displayName(cfg config, takenNames map[string]bool) string {
	if cfg.DisplayName == "" {
		return "instance-" + time.Now().Format("20060102-1504")
	}
	name := cfg.DisplayName
	for n := 2; takenNames[name]; n++ {
		name = fmt.Sprintf("%s-%d", cfg.DisplayName, n)
	}
	return name
}
