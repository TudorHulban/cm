package services

import (
	"test/domain/inventory"
)

type ParamsGetInventoryForService struct {
	ServiceName string `valid:"required"`
}

func (s *ServiceMain) GetInventoryForService(params *ParamsGetInventoryForService) ([]*inventory.Entry, error) {
	return s.inventory.FindInventoryForService(params.ServiceName)
}
