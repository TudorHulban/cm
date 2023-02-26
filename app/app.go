package app

import (
	"errors"
	"test/app/apperrors"
	"test/app/configuration"
	"test/app/services"

	"test/infra"
	"test/infra/web"

	"github.com/asaskevich/govalidator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type PiersApp struct {
	ServiceMain *services.ServiceMain
}

type AppCo struct {
	Configuration *configuration.Configuration
	ServiceMain   *services.ServiceMain

	webServer *web.WebServer

	LogLevel zerolog.Level
	Mode     int // production, testing, development
}

func NewApp(piers *PiersApp, config *configuration.Configuration) (*AppCo, error) {
	if _, errVa := govalidator.ValidateStruct(config); errVa != nil {
		return nil, apperrors.ErrValidation{
			Caller: "NewApp",
			Issue:  errVa,
		}
	}

	if piers.ServiceMain == nil {
		return nil, errors.New("passed service main is nil")
	}

	webServer, errIni := infra.InitializeWeb(piers.ServiceMain, config)
	if errIni != nil {
		return nil, &apperrors.ErrInfrastructure{
			Caller:  "NewApp",
			Calling: "infra.InitializeWeb",
			Issue:   errIni,
		}
	}

	return &AppCo{
		Configuration: config,
		webServer:     webServer,
	}, nil
}

func (a *AppCo) initLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(a.LogLevel)

	log.Info().Msgf("app log level set to %s", a.LogLevel.String())
}

func (a *AppCo) Start() error {
	a.initLogging()

	return a.webServer.Start()
}

func (a *AppCo) Close() error {
	return a.webServer.Stop()
}
