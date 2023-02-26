package apperrors

import "errors"

const ErrorMsgForContextExpiration = "context expired"

var (
	OSExitForConfigurationIssues        = 1
	OSExitForDatabaseIssues             = 2
	OSExitForRepositoryIssues           = 3
	OSExitForRepositoryMigrationsIssues = 4
	OSExitForServiceIssues              = 5
	OSExitForFileOperationsIssues       = 6
	OSExitForGRPCIssues                 = 7
	OSExitForSeederIssues               = 8
	OSExitForWebServerIssues            = 9
	OSExitForInitializationIssues       = 10
)

var errItemNotFound = errors.New("item not found")
