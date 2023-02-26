package infra

import (
	"test/app/apperrors"
	"test/app/configuration"
	"test/app/services"

	"test/infra/apis"
	"test/infra/controller"
	"test/infra/rest"
	"test/infra/web"
)

func InitializeWeb(service *services.ServiceMain, config *configuration.Configuration) (*web.WebServer, error) {
	control, errCo := controller.NewController(&controller.PiersControllerWeb{
		ServiceMain: service,
	})
	if errCo != nil {
		return nil, &apperrors.ErrInfrastructure{
			Caller:  "InitializeWeb",
			Calling: "controller.NewController",
			Issue:   errCo,
		}
	}

	apis, errApis := apis.NewAPI(&apis.PiersAPI{
		ServiceMain: service,
		Controller:  control,
	})
	if errApis != nil {
		return nil, &apperrors.ErrInfrastructure{
			Caller:  "InitializeWeb",
			Calling: "apis.NewAPIS",
			Issue:   errApis,
		}
	}

	rest, errRest := rest.NewRest()
	if errRest != nil {
		return nil, &apperrors.ErrInfrastructure{
			Caller:  "InitializeWeb",
			Calling: "rest.NewRest",
			Issue:   errRest,
		}
	}

	web, errCr := web.NewWebServer(&web.PiersWebServer{
		APIS: apis,
		Rest: rest,
	}, config.AsWebServerConfiguration())
	if errCr != nil {
		return nil, &apperrors.ErrInfrastructure{
			Caller:  "InitializeWeb",
			Calling: "web.NewWebServer",
			Issue:   errCr,
		}
	}

	return web, nil
}
