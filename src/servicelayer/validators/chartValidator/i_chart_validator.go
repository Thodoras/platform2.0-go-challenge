package chartValidator

import "platform2.0-go-challenge/src/models"

type IChartValidator interface {
	ValidateChart(chart models.Chart) error
}
