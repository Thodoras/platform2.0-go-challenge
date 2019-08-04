package insightValidator

import "platform2.0-go-challenge/src/models"

type IInsightValidator interface {
	ValidateInsight(insight models.Insight) error
}
