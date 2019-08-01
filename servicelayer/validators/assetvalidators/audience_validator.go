package assetvalidators

import (
	"errors"
	"regexp"

	"platform2.0-go-challenge/helpers/errorutils"

	"platform2.0-go-challenge/models/assets"
)

const prefix string = "Validation error: "

var genderRegex = regexp.MustCompile(`^(m|f)$`)
var birthCountryRegex = regexp.MustCompile(`^[A-Za-z]{1,3}$`)
var ageGroupsRegex = regexp.MustCompile(`^\d{1,3}-\d{1,3}$`)

func ValidateAudience(audience assets.Audience) error {
	var err error
	err = validateGender(audience.Gender)
	if err != nil {
		return err
	}

	err = validateBirthCountry(audience.BirthCountry)
	if err != nil {
		return err
	}

	err = validateAgeGroups(audience.AgeGroups)
	if err != nil {
		return err
	}

	err = validateHoursSpent(audience.HoursSpent)
	if err != nil {
		return err
	}

	err = validateNumOfPurchasesPerMonth(audience.NumOfPurchasesPerMonth)
	if err != nil {
		return err
	}

	return nil
}

func validateGender(gender string) error {
	if genderRegex.MatchString(gender) {
		return nil
	}
	errorutils.InvalidRequest = errors.New(prefix + "Invalid gender")
	return errorutils.InvalidRequest
}

func validateBirthCountry(birthCounty string) error {
	if birthCountryRegex.MatchString(birthCounty) {
		return nil
	}
	errorutils.InvalidRequest = errors.New(prefix + "Invalid birth country")
	return errorutils.InvalidRequest
}

func validateAgeGroups(ageGroup string) error {
	if ageGroupsRegex.MatchString(ageGroup) {
		return nil
	}
	errorutils.InvalidRequest = errors.New(prefix + "Invalid age group")
	return errorutils.InvalidRequest
}

func validateHoursSpent(hoursSpent int) error {
	if hoursSpent >= 0 {
		return nil
	}
	errorutils.InvalidRequest = errors.New(prefix + "Invalid hours spent")
	return errorutils.InvalidRequest
}

func validateNumOfPurchasesPerMonth(numOfPurchasesPerMonth int) error {
	if numOfPurchasesPerMonth >= 0 {
		return nil
	}
	errorutils.InvalidRequest = errors.New(prefix + "Invalid number of purchases per month")
	return errorutils.InvalidRequest
}
