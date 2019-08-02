package validators

import (
	"testing"

	"platform2.0-go-challenge/helpers/testutils"
	"platform2.0-go-challenge/models"
	"platform2.0-go-challenge/servicelayer/validators"
)

func mockUserRequestHappyPath() models.User {
	return models.User{
		ID:       1,
		Name:     "foo",
		Password: "Fo0b4r1!",
	}
}

func TestUserValidateorShouldReturnNoErrorWhenHappyPath(t *testing.T) {

	mockRequest := mockUserRequestHappyPath()

	err := validators.ValidateUser(mockRequest)

	testutils.AssertNoError(err, t)
}

func TestUserValidateorShouldReturnErrorWhenInvalidName(t *testing.T) {
	mockRequest1 := mockUserRequestHappyPath()
	mockRequest1.Name = "" // No name
	mockRequest2 := mockUserRequestHappyPath()
	mockRequest2.Name = "fo" // Too short name
	mockRequest3 := mockUserRequestHappyPath()
	mockRequest3.Name = "abcd56789012345678901" // Too long name

	err1 := validators.ValidateUser(mockRequest1)
	err2 := validators.ValidateUser(mockRequest2)
	err3 := validators.ValidateUser(mockRequest3)

	testutils.AssertError(err1, t)
	testutils.AssertError(err2, t)
	testutils.AssertError(err3, t)
}

func TestUserValidateorShouldReturnErrorWhenInvalidPassword(t *testing.T) {
	mockRequest1 := mockUserRequestHappyPath()
	mockRequest1.Password = "" // No password
	mockRequest2 := mockUserRequestHappyPath()
	mockRequest2.Password = "Aa1!a" // Too short password
	mockRequest3 := mockUserRequestHappyPath()
	mockRequest3.Password = "Aa1!1234567890123" // Too long password
	mockRequest4 := mockUserRequestHappyPath()
	mockRequest4.Password = "aa1!1234" // No upper case letter
	mockRequest5 := mockUserRequestHappyPath()
	mockRequest5.Password = "AA1!1234" // No lower case letter
	mockRequest6 := mockUserRequestHappyPath()
	mockRequest6.Password = "Aa!!!@#$" // No digits
	mockRequest7 := mockUserRequestHappyPath()
	mockRequest7.Password = "Aa111234" // No special characters

	err1 := validators.ValidateUser(mockRequest1)
	err2 := validators.ValidateUser(mockRequest2)
	err3 := validators.ValidateUser(mockRequest3)
	err4 := validators.ValidateUser(mockRequest4)
	err5 := validators.ValidateUser(mockRequest5)
	err6 := validators.ValidateUser(mockRequest6)
	err7 := validators.ValidateUser(mockRequest7)

	testutils.AssertError(err1, t)
	testutils.AssertError(err2, t)
	testutils.AssertError(err3, t)
	testutils.AssertError(err4, t)
	testutils.AssertError(err5, t)
	testutils.AssertError(err6, t)
	testutils.AssertError(err7, t)
}

func TestLoginValidateorShouldReturnNoErrorWhenHappyPath(t *testing.T) {

	mockRequest := mockUserRequestHappyPath()

	err := validators.ValidateLoginCredentials(mockRequest)

	testutils.AssertNoError(err, t)
}

func TestLoginValidateorShouldReturnNoErrorWhenNameMissing(t *testing.T) {

	mockRequest := mockUserRequestHappyPath()
	mockRequest.Name = ""

	err := validators.ValidateLoginCredentials(mockRequest)

	testutils.AssertError(err, t)
}

func TestLoginValidateorShouldReturnNoErrorWhenNamePassword(t *testing.T) {

	mockRequest := mockUserRequestHappyPath()
	mockRequest.Password = ""

	err := validators.ValidateLoginCredentials(mockRequest)

	testutils.AssertError(err, t)
}
