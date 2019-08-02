package validators

import (
	"regexp"

	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/models"
)

const minNameLength = "3"
const maxNameLength = "20"
const minPasswordLength = "6"
const maxPasswordLength = "16"
const invalidNameDescription = "Name should have no empty spaces, and be of length from " + minNameLength + " to " + maxNameLength
const invalidPasswordDescription = "Password should contain at least one lower case letter, one upper case and one digit and be of length from " + minPasswordLength + " to " + maxPasswordLength
const nameMissingDescription = "Missing name"
const passwordMissingDescription = "Missing password"

var nameRegex = regexp.MustCompile(`^[\w]{` + minNameLength + `,` + maxNameLength + `}$`)
var oneUpperCaseLetter = regexp.MustCompile(`[A-Z]`)
var oneLowerCaseLetter = regexp.MustCompile(`[a-z]`)
var oneDigit = regexp.MustCompile(`\d`)
var oneSpecialCharacter = regexp.MustCompile(`[!@#$%^&*]`)
var passwordFormat = regexp.MustCompile(`^.{` + minPasswordLength + `,` + maxPasswordLength + `}$`)

func ValidateUser(user models.User) error {
	var err error
	err = validateName(user.Name)
	if err != nil {
		return err
	}

	err = validatePassword(user.Password)
	if err != nil {
		return err
	}

	return nil
}

func validateName(name string) error {
	if nameRegex.MatchString(name) {
		return nil
	}

	return errorutils.NewInvalidRequest(invalidNameDescription)
}

func validatePassword(password string) error {
	if passwordFormat.MatchString(password) &&
		oneUpperCaseLetter.MatchString(password) &&
		oneLowerCaseLetter.MatchString(password) &&
		oneDigit.MatchString(password) &&
		oneSpecialCharacter.MatchString(password) {
		return nil
	}

	return errorutils.NewInvalidRequest(invalidPasswordDescription)
}

func ValidateLoginCredentials(user models.User) error {
	var err error
	err = validateNameExists(user.Name)
	if err != nil {
		return err
	}

	err = validatePasswordExists(user.Password)
	if err != nil {
		return err
	}

	return nil
}

func validateNameExists(name string) error {
	if name != "" {
		return nil
	}

	return errorutils.NewInvalidRequest(nameMissingDescription)
}

func validatePasswordExists(password string) error {
	if password != "" {
		return nil
	}

	return errorutils.NewInvalidRequest(passwordMissingDescription)
}
