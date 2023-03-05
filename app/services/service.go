package services

import (
	"errors"
	"test/domain"
	"test/domain/inventory"
)

type PiersServiceMain struct {
	Configuration *domain.Configuration
	Inventory     *inventory.Inventory
}

type ServiceMain struct {
	configuration *domain.Configuration
	inventory     *inventory.Inventory
}

func NewServiceMain(piers *PiersServiceMain) (*ServiceMain, error) {
	if piers.Configuration == nil {
		return nil, errors.New("passed configuration is nil")
	}

	if piers.Inventory == nil {
		return nil, errors.New("passed inventory is nil")
	}

	return &ServiceMain{
		configuration: piers.Configuration,
		inventory:     piers.Inventory,
	}, nil
}
