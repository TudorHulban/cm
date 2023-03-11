package inventory

import (
	"fmt"
	"sync"
	"test/app/apperrors"
	"time"

	"github.com/asaskevich/govalidator"
)

type TargetID string // can be in the form of environment/domain/service

type Target struct {
	Name        string
	LastCheckin time.Time

	Services []*Service

	sync.Mutex
}

type ParamsNewTarget struct {
	ID   string `valid:"required"`
	Name string `valid:"required"`
}

func (inv *Inventory) NewTarget(params *ParamsNewTarget, services ...*Service) (*Target, error) {
	if _, errVa := govalidator.ValidateStruct(params); errVa != nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewTarget",
			Issue:  errVa,
		}
	}

	for id, target := range inv.Targets {
		if id == TargetID(params.ID) {
			return nil, apperrors.ErrValidation{
				Caller: "NewTarget",
				Issue:  fmt.Errorf("already existing target ID:%s", params.ID),
			}
		}

		if target.Name == params.Name {
			return nil, apperrors.ErrValidation{
				Caller: "NewTarget",
				Issue:  fmt.Errorf("already existing target Name:%s", params.Name),
			}
		}
	}

	res := Target{
		Name:     params.Name,
		Services: services,
	}

	inv.Targets[TargetID(params.ID)] = &res

	return &res, nil
}

func (t *Target) AddServices(services ...*Service) {
	t.Services = append(t.Services, services...)
}

func (inv *Inventory) CheckIn(target string) error {
	t, exists := inv.Targets[TargetID(target)]
	if !exists {
		return apperrors.ErrValidation{}
	}

	t.Lock()
	defer t.Unlock()

	t.LastCheckin = time.Now()

	return nil
}
