package services

import (
	"test/app/apperrors"

	"github.com/asaskevich/govalidator"
)

type ParamsAddEntry struct {
	ServiceName string `valid:"required"`
	Version     string `valid:"required"`
}

func (s *ServiceMain) AddInventoryEntry(params *ParamsAddEntry) error {
	if _, errVa := govalidator.ValidateStruct(params); errVa != nil {
		return apperrors.ErrValidation{
			Caller: "AddEntry",
			Issue:  errVa,
		}
	}

	return s.inventory.AddEntry(params.ServiceName, params.Version)
}
