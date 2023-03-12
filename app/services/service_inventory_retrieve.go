package services

import (
	"test/app/apperrors"
	"test/domain/inventory"
)

type ParamsInventoryGetServices struct {
	TargetID       inventory.TargetID `valid:"required"`
	ServiceName    string             `valid:"required"`
	ServiceVersion string
}

func (s *ServiceMain) InventoryGetServices(params *ParamsInventoryGetServices) ([]*inventory.Service, error) {
	reconstructedServices := s.inventory.FindServices(params.ServiceName, params.ServiceVersion, params.TargetID)

	if len(reconstructedServices) == 0 {
		return nil, apperrors.ErrDomain{
			Caller:     "InventoryGetServices",
			NameMethod: "inventory.FindServices",
			Issue:      apperrors.ErrRecordNotFound{},
		}
	}

	return reconstructedServices, nil
}
