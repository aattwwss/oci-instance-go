package main

//
// import "time"
//
// type CreateInstanceReq struct {
// 	Metadata           Metadata           `json:"metadata"`
// 	Shape              string             `json:"shape"`
// 	CompartmentID      string             `json:"compartmentId"`
// 	DisplayName        string             `json:"displayName"`
// 	AvailabilityDomain string             `json:"availabilityDomain"`
// 	SourceDetails      []SourceDetail     `json:"sourceDetails"`
// 	CreateVnicDetails  CreateVnicDetails  `json:"createVnicDetails"`
// 	AgentConfig        AgentConfig        `json:"agentConfig"`
// 	DefinedTags        DefinedTags        `json:"definedTags"`
// 	FreeformTags       FreeformTags       `json:"freeformTags"`
// 	InstanceOptions    InstanceOptions    `json:"instanceOptions"`
// 	AvailabilityConfig AvailabilityConfig `json:"availabilityConfig"`
// 	ShapeConfig        ShapeConfig        `json:"shapeConfig"`
// }
//
// type Metadata struct {
// 	SSHAuthorizedKeys string `json:"ssh_authorized_keys"`
// }
//
// type CreateVnicDetails struct {
// 	AssignPublicIP         bool   `json:"assignPublicIp"`
// 	SubnetID               string `json:"subnetId"`
// 	AssignPrivateDNSRecord bool   `json:"assignPrivateDnsRecord"`
// }
//
// type PluginsConfig struct {
// 	Name         string `json:"name"`
// 	DesiredState string `json:"desiredState"`
// }
//
// type AgentConfig struct {
// 	PluginsConfig        []PluginsConfig `json:"pluginsConfig"`
// 	IsMonitoringDisabled bool            `json:"isMonitoringDisabled"`
// 	IsManagementDisabled bool            `json:"isManagementDisabled"`
// }
//
// type DefinedTags struct {
// }
//
// type FreeformTags struct {
// }
//
// type InstanceOptions struct {
// 	AreLegacyImdsEndpointsDisabled bool `json:"areLegacyImdsEndpointsDisabled"`
// }
//
// type AvailabilityConfig struct {
// 	RecoveryAction string `json:"recoveryAction"`
// }
//
// type ShapeConfig struct {
// 	Ocpus       int `json:"ocpus"`
// 	MemoryInGBs int `json:"memoryInGBs"`
// }
//
// type SourceDetail struct {
// 	SourceType          string `json:"sourceType"`
// 	ImageId             string `json:"imageId,omitempty"`
// 	BootVolumeId        string `json:"bootVolumeId,omitempty"`
// 	BootVolumeSizeInGBs int    `json:"bootVolumeSizeInGBs,omitempty"`
// }
//
// func NewSourceDetail(cfg config) SourceDetail {
// 	if cfg.BootVolumeId != "" {
// 		return SourceDetail{
// 			SourceType:   "bootVolume",
// 			BootVolumeId: cfg.BootVolumeId,
// 		}
// 	}
// 	return SourceDetail{
// 		SourceType: "image",
// 		ImageId:    cfg.ImageID,
// 	}
// }
//
// func NewCreateInstanceReq(cfg config, domain string) CreateInstanceReq {
// 	return CreateInstanceReq{
// 		Metadata: Metadata{
// 			SSHAuthorizedKeys: cfg.SSHPublicKey},
// 		Shape:              cfg.Shape,
// 		CompartmentID:      cfg.TenancyID,
// 		DisplayName:        "instance-" + time.Now().Format("20060102-1504"),
// 		AvailabilityDomain: domain,
// 		SourceDetails:      []SourceDetail{NewSourceDetail(cfg)},
// 		CreateVnicDetails: CreateVnicDetails{
// 			AssignPublicIP:         false,
// 			SubnetID:               cfg.SubnetID,
// 			AssignPrivateDNSRecord: true,
// 		},
// 		AgentConfig: AgentConfig{
// 			PluginsConfig: []PluginsConfig{
// 				{
// 					Name:         "Compute Instance Monitoring",
// 					DesiredState: "ENABLED",
// 				},
// 			},
// 			IsMonitoringDisabled: false,
// 			IsManagementDisabled: false,
// 		},
// 		DefinedTags:  DefinedTags{},
// 		FreeformTags: FreeformTags{},
// 		InstanceOptions: InstanceOptions{
// 			AreLegacyImdsEndpointsDisabled: false,
// 		},
// 		AvailabilityConfig: AvailabilityConfig{
// 			RecoveryAction: "RESTORE_INSTANCE",
// 		},
// 		ShapeConfig: ShapeConfig{
// 			Ocpus:       cfg.OCPUS,
// 			MemoryInGBs: cfg.MemoryInGbs,
// 		},
// 	}
// }
