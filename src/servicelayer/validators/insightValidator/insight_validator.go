package insightValidator

import (
	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/models"
)

type InsightValidator struct{}

func NewInsightValidator() *InsightValidator {
	return &InsightValidator{}
}

func (i *InsightValidator) ValidateInsight(insight models.Insight) error {
	var err error
	err = i.validateText(insight.Text)
	if err != nil {
		return err
	}

	return nil
}

func (i *InsightValidator) validateText(text string) error {
	if text != "" {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid text")
}
