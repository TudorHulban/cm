package domain

import "errors"

type Service struct {
	Name          string
	Configuration []OSVariableName
}

func NewService(name string, vars ...OSVariableName) (*Service, error) {
	if len(name) == 0 {
		return nil, errors.New("passed service name is nil")
	}

	return &Service{
		Name:          name,
		Configuration: vars,
	}, nil
}
