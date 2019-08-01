package assetvalidators

import (
	"platform2.0-go-challenge/helpers/errorutils"
	"platform2.0-go-challenge/models/assets"
)

func ValidateChart(chart assets.Chart) error {
	var err error
	err = validateTitle(chart.Title)
	if err != nil {
		return err
	}

	err = validateData(chart.Data)
	if err != nil {
		return err
	}

	return nil
}

func validateTitle(title string) error {
	if len(title) >= 3 {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid title")
}

func validateData(data string) error {
	if data != "" {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid data")
}
