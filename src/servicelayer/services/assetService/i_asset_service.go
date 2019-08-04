package assetService

import (
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/weblayer/dtos"
)

type IAssetService interface {
	GetAllAssets(id int) (*dtos.AssetReponse, error)
	GetAllAssetsPaginated(userID, limit, offset int) (*dtos.AssetReponse, error)
	AddAudience(audience models.Audience) (int, error)
	AddChart(chart models.Chart) (int, error)
	AddInsight(insight models.Insight) (int, error)
	EditAudience(audience models.Audience) (int64, error)
	EditChart(chart models.Chart) (int64, error)
	EditInsight(insight models.Insight) (int64, error)
	DeleteAudience(id int) (int64, error)
	DeleteChart(id int) (int64, error)
	DeleteInsight(id int) (int64, error)
}
