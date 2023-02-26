package apperrors

import (
	"errors"
	"fmt"
)

const (
	ErrorMsgConnectionCreation = "DB connection creation: %s"
)

// ErrRecordNotFound is used to indicate that selecting an individual object
// yielded no result. Declared as type, not value, for consistency reasons.
type ErrRecordNotFound struct{} //TODO: add optional message

const ErrRecordNotFoundMessage = "record not found"

func (ErrRecordNotFound) Error() string {
	return ErrRecordNotFoundMessage
}

func (ErrRecordNotFound) Unwrap() error {
	return errors.New(ErrRecordNotFoundMessage)
}

// ErrMultipleRowsFound is used to indicate that selecting an individual object
// yielded multiple results.
type ErrMultipleRowsFound struct {
	HowMany int64
}

const ErrMultipleRowsFoundMessage = "multiple rows found: %d"

func (e ErrMultipleRowsFound) Error() string {
	return fmt.Sprintf(ErrMultipleRowsFoundMessage, e.HowMany)
}

func (ErrMultipleRowsFound) Unwrap() error {
	return errors.New(ErrRecordNotFoundMessage)
}
