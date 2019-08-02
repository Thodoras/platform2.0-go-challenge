package validators

import (
	"testing"

	"platform2.0-go-challenge/src/helpers/testutils"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/validators"
)

func mockInsightRequestHappyPath() models.Insight {
	return models.Insight{
		ID:     1,
		UserID: 1,
		Text:   "foo",
	}
}

func TestInsightValidateorShouldReturnNoErrorWhenHappyPath(t *testing.T) {

	mockRequest := mockInsightRequestHappyPath()

	err := validators.ValidateInsight(mockRequest)

	testutils.AssertNoError(err, t)
}

func TestInsightValidateorShouldReturnErrorWhenInvalidText(t *testing.T) {

	mockRequest := mockInsightRequestHappyPath()
	mockRequest.Text = ""

	err := validators.ValidateInsight(mockRequest)

	testutils.AssertError(err, t)
}
