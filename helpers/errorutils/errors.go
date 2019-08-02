package errorutils

import "errors"

const prefix string = "Validation error: "
const UniqueConstrainViolationString = "duplicate key value violates unique constraint"

var InvalidRequest error
var UniqueConstrainViolation error

func NewInvalidRequest(message string) error {
	InvalidRequest = errors.New(prefix + message)
	return InvalidRequest
}

func NewUniqueConstrainViolation(message string) error {
	UniqueConstrainViolation = errors.New(prefix + message)
	return UniqueConstrainViolation
}
