package main

import (
	"errors"
	"os"
)

func LoadConfiguration() (Region *string, RegistryId *string, Error error) {
	var region = os.Getenv("AWS_REGION")
	var registryId = os.Getenv("REGISTRY_ID")

	if len(region) == 0 {
		return nil, nil, errors.New("region is empty")
	}

	if len(registryId) == 0 {
		return nil, nil, errors.New("registry is is empty")
	}

	return &region, &registryId, nil
}
