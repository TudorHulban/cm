package inventory

import (
	"sync"
	"test/app/apperrors"
	"test/domain"
	"time"

	"github.com/asaskevich/govalidator"
)

type Service struct {
	ID          string //should be unique among total inventory
	Name        string
	Version     string
	LastCheckin time.Time

	Entries []*domain.Entry

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

	return &Service{
		ID:      params.ID,
		Name:    params.Name,
		Version: params.Version,
		Entries: entries,
	}, nil
}

func (s *Service) AddEntries(entries ...*domain.Entry) {
	s.Entries = append(s.Entries, entries...)
}

func (s *Service) CheckIn() {
	s.Lock()
	defer s.Unlock()

	s.LastCheckin = time.Now()
}
