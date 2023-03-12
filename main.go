package main

import (
	"embed"
	"os"

	"test/app"
	"test/app/apperrors"
	"test/app/configuration"
	"test/app/services"
	"test/domain/inventory"

	"test/domain"
	"test/helpers"

	"github.com/rs/zerolog/log"
)

//go:embed configs
var _fs embed.FS

func main() {
	config := configuration.Configuration{
		ConfigurationWebServer: configuration.ConfigurationWebServer{
			APIsPathCommon: "",
			Port:           "9998",
		},
	}

	fsEntries, errRead := loadFSEntries()
	if errRead != nil {
		log.Error().Msg(helpers.ReplEOL(errRead.Error()))

		os.Exit(apperrors.OSExitForFileOperationsIssues)
	}

	configFiles, errLoad := loadVarsFiles(fsEntries)
	if errLoad != nil {
		log.Error().Msg(helpers.ReplEOL(errLoad.Error()))

		os.Exit(apperrors.OSExitForFileOperationsIssues)
	}

	configuration, errCr := domain.NewConfigurationFrom(configFiles)
	if errCr != nil {
		log.Error().Msg(helpers.ReplEOL(errCr.Error()))

		os.Exit(apperrors.OSExitForConfigurationIssues)
	}

	inventory := inventory.NewInventory()

	service, errServ := services.NewServiceMain(&services.PiersServiceMain{
		Configuration: configuration,
		Inventory:     inventory,
	})
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
