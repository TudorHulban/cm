package sinventory

import (
	"errors"
	"test/domain/inventory"
)

type PiersServiceInventory struct {
	Inventory *inventory.Inventory
}

type ServiceInventory struct {
	inventory *inventory.Inventory
}

func NewServiceMain(piers *PiersServiceInventory) (*ServiceInventory, error) {
	if piers.Inventory == nil {
		return nil, errors.New("passed inventory is nil")
	}

	return &ServiceInventory{
		inventory: piers.Inventory,
	}, nil
}
