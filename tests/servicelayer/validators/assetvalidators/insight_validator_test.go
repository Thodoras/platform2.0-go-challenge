package assetvalidators

import (
	"testing"

	"platform2.0-go-challenge/helpers/testutils"
	"platform2.0-go-challenge/models/assets"
	"platform2.0-go-challenge/servicelayer/validators/assetvalidators"
)

func mockInsightRequestHappyPath() assets.Insight {
	return assets.Insight{
		ID:     1,
		UserID: 1,
		Text:   "foo",
	}
}

func TestInsightValidateorShouldReturnNoErrorWhenHappyPath(t *testing.T) {

	mockRequest := mockInsightRequestHappyPath()

	err := assetvalidators.ValidateInsight(mockRequest)

	testutils.AssertNoError(err, t)
}

func TestInsightValidateorShouldReturnErrorWhenInvalidText(t *testing.T) {

	mockRequest := mockInsightRequestHappyPath()
	mockRequest.Text = ""

	err := assetvalidators.ValidateInsight(mockRequest)

	testutils.AssertError(err, t)
}
