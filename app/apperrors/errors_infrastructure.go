package apperrors

import (
	"fmt"
)

type ErrInfrastructure struct {
	Issue error

	Caller  string
	Calling string
}

const areaErrInfrastructure = "Infrastructure"

func (e *ErrInfrastructure) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrInfrastructure)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Calling: %s", e.Calling)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
