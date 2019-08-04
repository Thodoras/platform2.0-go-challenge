package userValidator

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

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (u *UserValidator) ValidateUser(user models.User) error {
	var err error
	err = u.validateName(user.Name)
	if err != nil {
		return err
	}

	err = u.validatePassword(user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserValidator) validateName(name string) error {
	if nameRegex.MatchString(name) {
		return nil
	}

	return errorutils.NewInvalidRequest(invalidNameDescription)
}

func (u *UserValidator) validatePassword(password string) error {
	if passwordFormat.MatchString(password) &&
		oneUpperCaseLetter.MatchString(password) &&
		oneLowerCaseLetter.MatchString(password) &&
		oneDigit.MatchString(password) &&
		oneSpecialCharacter.MatchString(password) {
		return nil
	}

	return errorutils.NewInvalidRequest(invalidPasswordDescription)
}

func (u *UserValidator) ValidateLoginCredentials(user models.User) error {
	var err error
	err = u.validateNameExists(user.Name)
	if err != nil {
		return err
	}

	err = u.validatePasswordExists(user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserValidator) validateNameExists(name string) error {
	if name != "" {
		return nil
	}

	return errorutils.NewInvalidRequest(nameMissingDescription)
}

func (u *UserValidator) validatePasswordExists(password string) error {
	if password != "" {
		return nil
	}

	return errorutils.NewInvalidRequest(passwordMissingDescription)
}
