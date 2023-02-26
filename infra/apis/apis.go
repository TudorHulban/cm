package apis

import (
	"errors"
	"test/app/services"
	"test/infra/controller"
)

type PiersAPI struct {
	ServiceMain *services.ServiceMain
	Controller  *controller.ControllerWeb
}

type API struct {
	serviceMain *services.ServiceMain
	controller  *controller.ControllerWeb
}

func NewAPI(piers *PiersAPI) (*API, error) {
	if piers.ServiceMain == nil {
		return nil, errors.New("passed service main is nil")
	}

	if piers.Controller == nil {
		return nil, errors.New("passed controller is nil")
	}

	return &API{
		serviceMain: piers.ServiceMain,
		controller:  piers.Controller,
	}, nil
}
