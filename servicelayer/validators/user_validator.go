package validators

import (
	"regexp"

	"platform2.0-go-challenge/helpers/errorutils"
	"platform2.0-go-challenge/models"
)

const nameDescription = "Name should have no empty spaces, and be of length from 3 to 20"
const passwordDescription = "Password should contain at least one lower case letter, one upper case and one digit and be of length from 6 to 16"

var nameRegex = regexp.MustCompile(`^[\w]{3,20}$`)
var oneUpperCaseLetter = regexp.MustCompile(`[A-Z]`)
var oneLowerCaseLetter = regexp.MustCompile(`[a-z]`)
var oneDigit = regexp.MustCompile(`\d`)
var oneSpecialCharacter = regexp.MustCompile(`[!@#$%^&*]`)
var passwordFormat = regexp.MustCompile(`^.{6,16}$`)

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

	return errorutils.NewInvalidRequest(nameDescription)
}

func validatePassword(password string) error {
	if passwordFormat.MatchString(password) &&
		oneUpperCaseLetter.MatchString(password) &&
		oneLowerCaseLetter.MatchString(password) &&
		oneDigit.MatchString(password) &&
		oneSpecialCharacter.MatchString(password) {
		return nil
	}

	return errorutils.NewInvalidRequest(passwordDescription)
}
