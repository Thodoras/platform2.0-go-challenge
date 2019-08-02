package validators

import (
	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/models"
)

func ValidateInsight(insight models.Insight) error {
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
