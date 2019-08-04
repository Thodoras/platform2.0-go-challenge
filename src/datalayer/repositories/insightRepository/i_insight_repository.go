package insightRepository

import "platform2.0-go-challenge/src/models"

type IInsightRepository interface {
	GetInsights(id int) ([]models.Insight, error)
	GetInsightsPaginated(userID, limit, offset int) ([]models.Insight, error)
	AddInsight(insight models.Insight) (int, error)
	EditInsight(insight models.Insight) (int64, error)
	DeleteInsight(id int) (int64, error)
}
