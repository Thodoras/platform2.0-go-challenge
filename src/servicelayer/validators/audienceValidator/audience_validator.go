package audienceValidator

import (
	"regexp"

	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/models"
)

var genderRegex = regexp.MustCompile(`^(m|f)$`)
var birthCountryRegex = regexp.MustCompile(`^[A-Za-z]{1,3}$`)
var ageGroupsRegex = regexp.MustCompile(`^\d{1,3}-\d{1,3}$`)

type AudienceValidator struct{}

func NewAudienceValidator() *AudienceValidator {
	return &AudienceValidator{}
}

func (a *AudienceValidator) ValidateAudience(audience models.Audience) error {
	var err error
	err = a.validateGender(audience.Gender)
	if err != nil {
		return err
	}

	err = a.validateBirthCountry(audience.BirthCountry)
	if err != nil {
		return err
	}

	err = a.validateAgeGroups(audience.AgeGroups)
	if err != nil {
		return err
	}

	err = a.validateHoursSpent(audience.HoursSpent)
	if err != nil {
		return err
	}

	err = a.validateNumOfPurchasesPerMonth(audience.NumOfPurchasesPerMonth)
	if err != nil {
		return err
	}

	return nil
}

func (a *AudienceValidator) validateGender(gender string) error {
	if genderRegex.MatchString(gender) {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid gender")
}

func (a *AudienceValidator) validateBirthCountry(birthCounty string) error {
	if birthCountryRegex.MatchString(birthCounty) {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid birth country")
}

func (a *AudienceValidator) validateAgeGroups(ageGroup string) error {
	if ageGroupsRegex.MatchString(ageGroup) {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid age group")
}

func (a *AudienceValidator) validateHoursSpent(hoursSpent int) error {
	if hoursSpent >= 0 {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid hours spent")
}

func (a *AudienceValidator) validateNumOfPurchasesPerMonth(numOfPurchasesPerMonth int) error {
	if numOfPurchasesPerMonth >= 0 {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid number of purchases per month")
}
