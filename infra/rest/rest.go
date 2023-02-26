package rest

type Rest struct{}

func NewRest() (*Rest, error) {
	return &Rest{}, nil
}
