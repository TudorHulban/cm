package services

import (
	"test/app/apperrors"
	"test/domain"

	"github.com/asaskevich/govalidator"
)

type ParamsFindTargetConfiguration struct {
	Target string
}

func (s *ServiceMain) GetTargetConfiguration(params *ParamsFindTargetConfiguration) (*domain.TargetConfiguration, error) {
	return s.configuration.FindTargetConfiguration(params.Target)
}

type ParamsGetVariableValues struct {
	Name    string `valid:"required"`
	Targets []string
}

func (s *ServiceMain) GetVariableValues(params *ParamsGetVariableValues) ([]domain.TargetValue, error) {
	if _, errVa := govalidator.ValidateStruct(params); errVa != nil {
		return nil, apperrors.ErrValidation{
			Caller: "GetVariableValues",
			Issue:  errVa,
		}
	}

	return s.configuration.FindVariableValues(params.Name, params.Targets...)
}

func (s *ServiceMain) GetTargets() ([]domain.TargetID, domain.TargetID) {
	return s.configuration.Targets, s.configuration.CurrentTarget
}

func (s *ServiceMain) GetCurrentTarget() domain.TargetID {
	return s.configuration.CurrentTarget
}
