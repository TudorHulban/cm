package inventory

import (
	"os"
	"test/domain"
	"testing"
)

func TestNewServiceFrom(t *testing.T) {
	serv := Service{
		Name:    "abcd",
		Version: "123",
		Entries: []*domain.Entry{
			{
				Name:  "one",
				Value: "value one",
			},
		},
	}

	serv.To(os.Stdout)
}
