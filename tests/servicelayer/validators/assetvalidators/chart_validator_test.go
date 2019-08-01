package assetvalidators

import (
	"testing"

	"platform2.0-go-challenge/helpers/testutils"
	"platform2.0-go-challenge/models/assets"
	"platform2.0-go-challenge/servicelayer/validators/assetvalidators"
)

func mockChartRequestHappyPath() assets.Chart {
	return assets.Chart{
		ID:         1,
		UserID:     1,
		Title:      "foo",
		AxisXTitle: "xfoo",
		AxisYTitle: "yfoo",
		Data:       "data foo",
	}
}

func TestChartValidateorShouldReturnNoErrorWhenHappyPath(t *testing.T) {

	mockRequest1 := mockChartRequestHappyPath()
	mockRequest2 := mockChartRequestHappyPath()
	mockRequest2.AxisXTitle = ""
	mockRequest2.AxisYTitle = ""

	err1 := assetvalidators.ValidateChart(mockRequest1)
	err2 := assetvalidators.ValidateChart(mockRequest2)

	testutils.AssertNoError(err1, t)
	testutils.AssertNoError(err2, t)
}

func TestChartValidateorShouldReturnErrorWhenInvalidTitle(t *testing.T) {
	mockRequest1 := mockChartRequestHappyPath()
	mockRequest1.Title = ""
	mockRequest2 := mockChartRequestHappyPath()
	mockRequest2.Title = "fo"

	err1 := assetvalidators.ValidateChart(mockRequest1)
	err2 := assetvalidators.ValidateChart(mockRequest2)

	testutils.AssertError(err1, t)
	testutils.AssertError(err2, t)
}

func TestChartValidateorShouldReturnErrorWhenInvalidData(t *testing.T) {
	mockRequest := mockChartRequestHappyPath()
	mockRequest.Title = ""

	err := assetvalidators.ValidateChart(mockRequest)

	testutils.AssertError(err, t)
	testutils.AssertError(err, t)
}
