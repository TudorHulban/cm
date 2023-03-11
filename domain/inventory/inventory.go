package inventory

import (
	"sync"
)

type Inventory struct {
	Targets       map[TargetID]*Target
	CurrentTarget TargetID

	sync.Mutex
}

func NewInventory() *Inventory {
	return &Inventory{
		Targets: make(map[TargetID]*Target),
	}
}

// func (inv *Inventory) AddEntry(serviceName, version string) error {
// 	inv.Lock()
// 	defer inv.Unlock()

// 	_, exists := inv.Items[ServiceName(serviceName)]
// 	if !exists {
// 		inv.Items[ServiceName(serviceName)] = []*Entry{
// 			{
// 				Timestamp: time.Now(),
// 				Version:   version,
// 			},
// 		}

// 		return nil
// 	}

// 	inv.Items[ServiceName(serviceName)] = append(inv.Items[ServiceName(serviceName)], &Entry{
// 		Timestamp: time.Now(),
// 		Version:   version,
// 	})

// 	return nil
// }

// func (inv *Inventory) FindInventoryForService(name string) ([]*Entry, error) {
// 	if entries, exists := inv.Items[ServiceName(name)]; exists {
// 		return entries, nil
// 	}

// 	return nil, apperrors.ErrRecordNotFound{}
// }
