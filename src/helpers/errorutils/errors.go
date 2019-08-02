package errorutils

import "errors"

const UniqueConstrainViolationString = "duplicate key value violates unique constraint"

var InvalidRequest error
var UniqueConstrainViolation error

func NewInvalidRequest(message string) error {
	InvalidRequest = errors.New(message)
	return InvalidRequest
}

func NewUniqueConstrainViolation(message string) error {
	UniqueConstrainViolation = errors.New(message)
	return UniqueConstrainViolation
}
