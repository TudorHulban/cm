package web

import (
	"errors"

	"test/app/apperrors"
	"test/infra/apis"
	"test/infra/rest"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type PiersWebServer struct {
	APIS *apis.API
	Rest *rest.Rest
}

type ConfigurationWebServer struct {
	APIsPathCommon string
	Port           string `valid:"required,port"`
}

type WebServer struct {
	App *fiber.App

	errShutdown error

	APIS *apis.API
	Rest *rest.Rest

	ConfigurationWebServer
}

func NewWebServer(piers *PiersWebServer, config *ConfigurationWebServer) (*WebServer, error) {
	if _, errVa := govalidator.ValidateStruct(config); errVa != nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewWebServer",
			Issue:  errVa,
		}
	}

	if piers.APIS == nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewWebServer",
			Issue:  errors.New("passed service apis is nil"),
		}
	}

	if piers.Rest == nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewWebServer",
			Issue:  errors.New("passed service rest is nil"),
		}
	}

	return &WebServer{
		ConfigurationWebServer: *config,
		App:                    fiber.New(),
		APIS:                   piers.APIS,
		Rest:                   piers.Rest,
	}, nil
}

func (w *WebServer) initLogging() error {
	return nil
}

func (w *WebServer) Start() error {
	w.initLogging()
	w.AddRoutes()

	log.Info().Msgf("server started on http://127.0.0.1:%s", w.Port)
	if w.errShutdown = w.App.Listen(":" + w.Port); w.errShutdown != nil {
		return w.errShutdown
	}

	return nil
}

func (w *WebServer) Stop() error {
	return nil
}
