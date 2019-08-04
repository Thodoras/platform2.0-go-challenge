package chartValidator

import (
	"platform2.0-go-challenge/src/helpers/errorutils"
	"platform2.0-go-challenge/src/models"
)

type ChartValidator struct{}

func NewChartValidator() *ChartValidator {
	return &ChartValidator{}
}

func (c *ChartValidator) ValidateChart(chart models.Chart) error {
	var err error
	err = c.validateTitle(chart.Title)
	if err != nil {
		return err
	}

	err = c.validateData(chart.Data)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChartValidator) validateTitle(title string) error {
	if len(title) >= 3 {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid title")
}

func (c *ChartValidator) validateData(data string) error {
	if data != "" {
		return nil
	}
	return errorutils.NewInvalidRequest("Invalid data")
}
