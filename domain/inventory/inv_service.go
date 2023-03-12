package inventory

import (
	"encoding/json"
	"errors"
	"io"
	"sync"
	"test/app/apperrors"
	"test/domain"
	"time"

	"github.com/asaskevich/govalidator"
)

type ServiceID string

type Service struct {
	Name        string    `json:"ServiceName"`
	Version     string    `json:"ServiceVersion"`
	LastCheckin time.Time `json:"-"`

	Entries []*domain.Entry `json:"Entries"`

	sync.Mutex
}

type ParamsNewService struct {
	ID      string `valid:"required"`
	Name    string `valid:"required"`
	Version string `valid:"required"`
}

func (inv *Inventory) NewService(params *ParamsNewService, entries ...*domain.Entry) (*Service, error) {
	if _, errVa := govalidator.ValidateStruct(params); errVa != nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewTarget",
			Issue:  errVa,
		}
	}

	res := Service{
		Name:    params.Name,
		Version: params.Version,
		Entries: entries,
	}

	inv.Lock()
	defer inv.Unlock()

	// TODO: check if already existing.

	inv.Services[ServiceID(params.ID)] = &res

	return &res, nil
}

// NewServicesFrom ads new services from content.
func (inv *Inventory) NewServicesFrom(serviceFiles map[string][]byte) error {
	if len(serviceFiles) == 0 {
		return errors.New("no service files contents passed")
	}

	for serviceID, serviceRaw := range serviceFiles {
		if _, exists := inv.Services[ServiceID(serviceID)]; exists {
			return apperrors.ErrValidation{
				Caller: "NewServicesFrom",
				Issue:  nil, //TODO: add
			}
		}

		newService, errCr := NewServiceFrom(serviceRaw)
		if errCr != nil {
			return apperrors.ErrValidation{
				Caller: "NewServicesFrom",
				Issue:  errCr,
			}
		}

		inv.AddService(ServiceID(serviceID), newService)
	}

	return nil
}

func NewServiceFrom(raw []byte) (*Service, error) {
	var res Service

	if errUn := json.Unmarshal(raw, &res); errUn != nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewServiceFrom",
			Issue:  errUn,
		}
	}

	return &res, nil
}

func (s *Service) AddEntries(entries ...*domain.Entry) {
	s.Lock()
	defer s.Unlock()

	s.Entries = append(s.Entries, entries...)
}

func (s *Service) GetEntry(name domain.OSVariableName) (*domain.Entry, error) {
	s.Lock()
	defer s.Unlock()

	for _, entry := range s.Entries {
		if entry.Name == name {
			return entry, nil
		}
	}

	return nil, apperrors.ErrRecordNotFound{}
}

// UpdateEntries updates entry if it exists.
func (s *Service) UpdateEntries(entries ...*domain.Entry) {
	s.Lock()
	defer s.Unlock()

	for _, entry := range entries {
		if reconstructedEntry, exists := s.GetEntry(entry.Name); exists == nil {
			reconstructedEntry.Value = entry.Value

			continue
		}
	}
}

func (s *Service) To(w io.Writer) (int, error) {
	raw, errMa := json.MarshalIndent(s, "", "")
	if errMa != nil {
		return 0, apperrors.ErrValidation{
			Caller: "To",
			Issue:  errMa,
		}
	}

	return w.Write(raw)
}

func (s *Service) CheckIn() {
	s.Lock()
	defer s.Unlock()

	s.LastCheckin = time.Now()
}
