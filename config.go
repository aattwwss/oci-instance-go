package main

import (
	"log"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type config struct {
	Region             string `env:"OCI_REGION"`
	UserID             string `env:"OCI_USER_ID"`
	TenancyID          string `env:"OCI_TENANCY_ID"`
	KeyFingerprint     string `env:"OCI_KEY_FINGERPRINT"`
	PrivateKeyFilename string `env:"OCI_PRIVATE_KEY_FILENAME"`
	SubnetID           string `env:"OCI_SUBNET_ID"`
	ImageID            string `env:"OCI_IMAGE_ID"`
	OCPUS              int    `env:"OCI_OCPUS"`
	MemoryInGbs        int    `env:"OCI_MEMORY_IN_GBS"`
	Shape              string `env:"OCI_SHAPE"`
	MaxInstances       int    `env:"OCI_MAX_INSTANCES"`
	SSHPublicKey       string `env:"OCI_SSH_PUBLIC_KEY"`

	// Optional
	AvailabilityDomain  string `env:"OCI_AVAILABILITY_DOMAIN"`
	BootVolumeSizeInGbs int    `env:"OCI_BOOT_VOLUME_SIZE_IN_GBS"`
	BootVolumeId        string `env:"OCI_BOOT_VOLUME_ID"`
}

func loadConfig() config {
	err := godotenv.Load() // load .env file
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
