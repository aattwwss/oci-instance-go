package main

import (
	"errors"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/oracle/oci-go-sdk/v65/core"
)

type config struct {
	Region             string  `env:"OCI_REGION"`
	UserID             string  `env:"OCI_USER_ID"`
	TenancyID          string  `env:"OCI_TENANCY_ID"`
	KeyFingerprint     string  `env:"OCI_KEY_FINGERPRINT"`
	PrivateKeyFilename string  `env:"OCI_PRIVATE_KEY_FILENAME"`
	SubnetID           string  `env:"OCI_SUBNET_ID"`
	ImageID            string  `env:"OCI_IMAGE_ID"`
	OCPUS              float32 `env:"OCI_OCPUS"`
	MemoryInGbs        float32 `env:"OCI_MEMORY_IN_GBS"`
	Shape              string  `env:"OCI_SHAPE"`
	MaxInstances       int     `env:"OCI_MAX_INSTANCES" envDefault:"1"`
	SSHPublicKey       string  `env:"OCI_SSH_PUBLIC_KEY"`

	// Optional
	AvailabilityDomains []string `env:"OCI_AVAILABILITY_DOMAIN" envSeparator:","`
	BootVolumeSizeInGbs int64    `env:"OCI_BOOT_VOLUME_SIZE_IN_GBS"`
	BootVolumeId        string   `env:"OCI_BOOT_VOLUME_ID"`
}

func (cfg config) validate() error {
	if cfg.BootVolumeId != "" && cfg.BootVolumeSizeInGbs != 0 {
		return errors.New("OCI_BOOT_VOLUME_ID and OCI_BOOT_VOLUME_SIZE_IN_GBS cannot be used together")
	}
	if cfg.BootVolumeId != "" && len(cfg.AvailabilityDomains) == 0 {
		return errors.New("OCI_AVAILABILITY_DOMAIN must be specified as string if using OCI_BOOT_VOLUME_ID")
	}
	return nil
}
func loadConfig() (config, error) {
	err := godotenv.Load() // load .env file
	if err != nil {
		return config{}, err
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	err = cfg.validate()
	if err != nil {
		return config{}, err
	}

	return cfg, nil
}

func buildSourceDetails(cfg config) core.InstanceSourceDetails {
	if cfg.BootVolumeId != "" {
		return core.InstanceSourceViaBootVolumeDetails{
			BootVolumeId: &cfg.BootVolumeId,
		}
	}

	var bootVolume *int64
	if cfg.BootVolumeSizeInGbs > 0 {
		bootVolume = &cfg.BootVolumeSizeInGbs
	}
	return core.InstanceSourceViaImageDetails{
		ImageId:             &cfg.ImageID,
		BootVolumeSizeInGBs: bootVolume,
	}
}
