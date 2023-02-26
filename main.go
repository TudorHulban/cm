package main

import (
	"embed"
	"os"

	"test/app"
	"test/app/apperrors"
	"test/app/configuration"
	"test/app/services"

	"test/domain"
	"test/helpers"

	"github.com/rs/zerolog/log"
)

//go:embed configs/*.json
var _fs embed.FS

func main() {
	config := configuration.Configuration{
		ConfigurationWebServer: configuration.ConfigurationWebServer{
			APIsPathCommon: "",
			Port:           "9998",
		},
	}

	configFiles, errLoad := loadConfigFiles()
	if errLoad != nil {
		log.Error().Msg(helpers.ReplEOL(errLoad.Error()))

		os.Exit(apperrors.OSExitForFileOperationsIssues)
	}

	configuration, errCr := domain.NewConfigurationFrom(configFiles)
	if errCr != nil {
		log.Error().Msg(helpers.ReplEOL(errCr.Error()))

		os.Exit(apperrors.OSExitForConfigurationIssues)
	}

	service, errServ := services.NewServiceMain(configuration)
	if errServ != nil {
		log.Error().Msg(helpers.ReplEOL(errServ.Error()))

		os.Exit(apperrors.OSExitForServiceIssues)
	}

	a, errApp := app.NewApp(&app.PiersApp{
		ServiceMain: service,
	}, &config)
	if errApp != nil {
		log.Error().Msg(helpers.ReplEOL(errApp.Error()))

		os.Exit(apperrors.OSExitForInitializationIssues)
	}

	a.Start()
}
