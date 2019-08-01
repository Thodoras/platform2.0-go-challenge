package errorutils

import "errors"

const prefix string = "Validation error: "

var InvalidRequest error

func NewInvalidRequest(message string) error {
	InvalidRequest = errors.New(prefix + message)
	return InvalidRequest
}
