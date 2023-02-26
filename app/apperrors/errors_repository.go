package apperrors

import "fmt"

const (
	ErrorMsgRepositoryCreation      = "repository creation: %s\n"
	ErrorMsgRepositoryMigrationsRun = "repository migrations: %s\n"
)

type ErrRepository struct {
	Issue           error
	Caller          string
	NameConstructor string
	NameMethod      string
}

const areaErrRepository = "Repository"

func (e ErrRepository) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrRepository)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	if len(e.NameMethod) == 0 {
		res[2] = fmt.Sprintf("NameConstructor: %s", e.NameConstructor)
	} else {
		res[2] = fmt.Sprintf("NameMethod: %s", e.NameMethod)
	}

	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
