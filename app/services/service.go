package services

import (
	"errors"
	"test/domain"
)

type ServiceMain struct {
	configuration *domain.Configuration
}

func NewServiceMain(config *domain.Configuration) (*ServiceMain, error) {
	if config == nil {
		return nil, errors.New("passed configuration is nil")
	}

	return &ServiceMain{
		configuration: config,
	}, nil
}
