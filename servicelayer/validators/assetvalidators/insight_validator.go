package assetvalidators

import (
	"platform2.0-go-challenge/helpers/errorutils"
	"platform2.0-go-challenge/models/assets"
)

func ValidateInsight(insight assets.Insight) error {
	var err error
	err = validateText(insight.Text)
	if err != nil {
		return err
	}

	return nil
}

func validateText(text string) error {
	if text != "" {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid text")
}
