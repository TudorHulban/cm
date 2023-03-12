package inventory

import (
	"sync"
	"test/app/apperrors"
	"time"
)

type Inventory struct {
	Targets       map[TargetID]*Target
	Services      map[ServiceID]*Service
	CurrentTarget TargetID

	sync.Mutex
}

func NewInventory() *Inventory {
	return &Inventory{
		Targets: make(map[TargetID]*Target),
	}
}

func (inv *Inventory) AddService(id ServiceID, service *Service) {
	inv.Lock()
	defer inv.Unlock()

	// TODO: check if not already existing

	inv.Services[id] = service
}

func (inv *Inventory) CheckIn(target TargetID) error {
	t, exists := inv.Targets[target]
	if !exists {
		return apperrors.ErrValidation{}
	}

	t.Lock()
	defer t.Unlock()

	t.LastCheckin = time.Now()

	return nil
}

func (inv *Inventory) FindServices(name, version string, inTargets ...TargetID) []*Service {
	var res []*Service

	for _, targetID := range inTargets {
		target, exists := inv.Targets[TargetID(targetID)]
		if !exists {
			continue
		}

		go inv.CheckIn(targetID)

		res = append(res, target.FindServiceByName(name, version)...)
	}

	return res
}
