package services

import (
	"platform2.0-go-challenge/datalayer/repositories"
	"platform2.0-go-challenge/models/assets"
	"platform2.0-go-challenge/weblayer/dtos/assetdtos"
)

func GetAllAssets(id int) (*assetdtos.AssetReponse, error) {
	var response assetdtos.AssetReponse
	var err error
	response.Audiences, err = repositories.GetAudiences(id)
	if err != nil {
		return nil, err
	}
	response.Charts, err = repositories.GetCharts(id)
	if err != nil {
		return nil, err
	}
	response.Insights, err = repositories.GetInsights(id)
	if err != nil {
		return nil, err
	}
	response.UserID = id
	return &response, nil
}

func AddAudience(audience assets.Audience) (int, error) {
	return repositories.AddAudience(audience)
}

func AddChart(chart assets.Chart) (int, error) {
	return repositories.AddChart(chart)
}

func AddInsight(insight assets.Insight) (int, error) {
	return repositories.AddInsight(insight)
}

func EditAudience(audience assets.Audience) (int64, error) {
	return repositories.EditAudience(audience)
}

func EditChart(chart assets.Chart) (int64, error) {
	return repositories.EditChart(chart)
}

func EditInsight(insight assets.Insight) (int64, error) {
	return repositories.EditInsight(insight)
}
