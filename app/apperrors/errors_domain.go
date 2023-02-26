package apperrors

import "fmt"

type ErrDomain struct {
	Issue error

	Caller          string
	NameConstructor string
	NameMethod      string
}

const areaErrDomain = "Domain"

func (e ErrDomain) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrDomain)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)

	if len(e.NameMethod) == 0 {
		res[2] = fmt.Sprintf("NameConstructor: %s", e.NameConstructor)
	} else {
		res[2] = fmt.Sprintf("NameMethod: %s", e.NameMethod)
	}

	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
