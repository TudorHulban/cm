package inventory

import (
	"sync"
	"test/app/apperrors"
	"time"
)

type ServiceName string

type Entry struct {
	Timestamp time.Time
	Version   string
}

type Inventory struct {
	Items map[ServiceName][]*Entry
	sync.Mutex
}

func NewInventory() *Inventory {
	return &Inventory{
		Items: make(map[ServiceName][]*Entry),
	}
}

func (inv *Inventory) AddEntry(serviceName, version string) error {
	inv.Lock()
	defer inv.Unlock()

	_, exists := inv.Items[ServiceName(serviceName)]
	if !exists {
		inv.Items[ServiceName(serviceName)] = []*Entry{
			{
				Timestamp: time.Now(),
				Version:   version,
			},
		}

		return nil
	}

	inv.Items[ServiceName(serviceName)] = append(inv.Items[ServiceName(serviceName)], &Entry{
		Timestamp: time.Now(),
		Version:   version,
	})

	return nil
}

func (inv *Inventory) FindInventoryForService(name string) ([]*Entry, error) {
	if entries, exists := inv.Items[ServiceName(name)]; exists {
		return entries, nil
	}

	return nil, apperrors.ErrRecordNotFound{}
}
