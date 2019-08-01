package testutils

import "testing"

func AssertNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}

func AssertError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
