package apperrors

import "fmt"

type ErrValidation struct {
	Issue  error
	Caller string
}

const areaErrValidation = "VALIDATION"

func (e ErrValidation) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrValidation)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n"
}
