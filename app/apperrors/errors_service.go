package apperrors

import "fmt"

const (
	ErrorMsgServiceCreation = "service creation: %s"
)

type ErrService struct {
	Issue error

	Caller          string
	NameConstructor string
	NameMethod      string
}

const areaErrService = "Service"

func (e ErrService) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrService)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	if len(e.NameMethod) == 0 {
		res[2] = fmt.Sprintf("NameConstructor: %s", e.NameConstructor)
	} else {
		res[2] = fmt.Sprintf("NameMethod: %s", e.NameMethod)
	}

	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
