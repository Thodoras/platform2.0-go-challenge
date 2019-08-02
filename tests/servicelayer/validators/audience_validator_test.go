package validators

import (
	"testing"

	"platform2.0-go-challenge/src/helpers/testutils"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/validators"
)

func mockAudienceRequestHappyPath() models.Audience {
	return models.Audience{
		ID:                     1,
		UserID:                 1,
		Gender:                 "f",
		BirthCountry:           "brz",
		AgeGroups:              "20-30",
		HoursSpent:             3,
		NumOfPurchasesPerMonth: 4,
	}
}

func TestAudienceValidateorShouldReturnNoErrorWhenHappyPath(t *testing.T) {

	mockRequest := mockAudienceRequestHappyPath()

	err := validators.ValidateAudience(mockRequest)

	testutils.AssertNoError(err, t)
}

func TestAudienceValidateorShouldReturnErrorWhenInvalidGender(t *testing.T) {
	mockRequest1 := mockAudienceRequestHappyPath()
	mockRequest1.Gender = "t"
	mockRequest2 := mockAudienceRequestHappyPath()
	mockRequest2.Gender = "male"

	err1 := validators.ValidateAudience(mockRequest1)
	err2 := validators.ValidateAudience(mockRequest2)

	testutils.AssertError(err1, t)
	testutils.AssertError(err2, t)
}

func TestAudienceValidateorShouldReturnErrorWhenInvalidBirthCountry(t *testing.T) {
	mockRequest1 := mockAudienceRequestHappyPath()
	mockRequest1.BirthCountry = "GRECE"
	mockRequest2 := mockAudienceRequestHappyPath()
	mockRequest2.BirthCountry = "u.s.a."

	err1 := validators.ValidateAudience(mockRequest1)
	err2 := validators.ValidateAudience(mockRequest2)

	testutils.AssertError(err1, t)
	testutils.AssertError(err2, t)
}

func TestAudienceValidateorShouldReturnErrorWhenInvalidAgeGroups(t *testing.T) {
	mockRequest1 := mockAudienceRequestHappyPath()
	mockRequest1.AgeGroups = "20/30"
	mockRequest2 := mockAudienceRequestHappyPath()
	mockRequest2.AgeGroups = "25"

	err1 := validators.ValidateAudience(mockRequest1)
	err2 := validators.ValidateAudience(mockRequest2)

	testutils.AssertError(err1, t)
	testutils.AssertError(err2, t)
}

func TestAudienceValidateorShouldReturnErrorWhenInvalidHoursSpent(t *testing.T) {
	mockRequest := mockAudienceRequestHappyPath()
	mockRequest.HoursSpent = -5

	err := validators.ValidateAudience(mockRequest)

	testutils.AssertError(err, t)
}

func TestAudienceValidateorShouldReturnErrorWhenInvalidNumOfPurchasesPerMonth(t *testing.T) {
	mockRequest := mockAudienceRequestHappyPath()
	mockRequest.NumOfPurchasesPerMonth = -5

	err := validators.ValidateAudience(mockRequest)

	testutils.AssertError(err, t)
}
