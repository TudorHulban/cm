package domain

import "errors"

type TargetID string // can be in the form of environment/domain/service

type Target struct {
	ID   TargetID
	Name string

	Services []Service
}

func NewTarget(id TargetID, name string, services ...Service) (*Target, error) {
	if len(name) == 0 {
		return nil, errors.New("passed target name is nil")
	}

	return &Target{
		ID:       id,
		Name:     name,
		Services: services,
	}, nil
}
