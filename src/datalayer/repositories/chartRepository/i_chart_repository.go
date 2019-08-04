package chartRepository

import "platform2.0-go-challenge/src/models"

type IChartRepository interface {
	GetCharts(userID int) ([]models.Chart, error)
	GetChartsPaginated(userID, limit, offset int) ([]models.Chart, error)
	AddChart(chart models.Chart) (int, error)
	EditChart(chart models.Chart) (int64, error)
	DeleteChart(id int) (int64, error)
}
