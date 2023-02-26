package controller

import (
	"errors"
	"test/app/services"
)

type PiersControllerWeb struct {
	ServiceMain *services.ServiceMain
}

type ControllerWeb struct {
	ServiceMain *services.ServiceMain
}

func NewController(piers *PiersControllerWeb) (*ControllerWeb, error) {
	if piers.ServiceMain == nil {
		return nil, errors.New("passed service main is nil")
	}

	return &ControllerWeb{
		ServiceMain: piers.ServiceMain,
	}, nil
}
