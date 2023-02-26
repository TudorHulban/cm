package services

import (
	"test/app/apperrors"

	"github.com/asaskevich/govalidator"
)

type ParamsDeleteTargetConfiguration struct {
	Target string `valid:"required"`
}

func (s *ServiceMain) DeleteTargetConfiguration(params *ParamsDeleteTargetConfiguration) error {
	if _, errVa := govalidator.ValidateStruct(params); errVa != nil {
		return apperrors.ErrValidation{
			Caller: "DeleteTargetConfiguration",
			Issue:  errVa,
		}
	}

	return s.configuration.DeleteTargetConfiguration(params.Target)
}
