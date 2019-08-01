package services

import (
	"platform2.0-go-challenge/datalayer/repositories"
	"platform2.0-go-challenge/models/assets"
	"platform2.0-go-challenge/servicelayer/validators/assetvalidators"
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
	err := assetvalidators.ValidateAudience(audience)
	if err != nil {
		return 0, err
	}

	return repositories.AddAudience(audience)
}

func AddChart(chart assets.Chart) (int, error) {
	err := assetvalidators.ValidateChart(chart)
	if err != nil {
		return 0, err
	}

	return repositories.AddChart(chart)
}

func AddInsight(insight assets.Insight) (int, error) {
	err := assetvalidators.ValidateInsight(insight)
	if err != nil {
		return 0, err
	}

	return repositories.AddInsight(insight)
}

func EditAudience(audience assets.Audience) (int64, error) {
	err := assetvalidators.ValidateAudience(audience)
	if err != nil {
		return 0, err
	}

	return repositories.EditAudience(audience)
}

func EditChart(chart assets.Chart) (int64, error) {
	err := assetvalidators.ValidateChart(chart)
	if err != nil {
		return 0, err
	}

	return repositories.EditChart(chart)
}

func EditInsight(insight assets.Insight) (int64, error) {
	err := assetvalidators.ValidateInsight(insight)
	if err != nil {
		return 0, err
	}

	return repositories.EditInsight(insight)
}

func DeleteAudience(id int) (int64, error) {
	return repositories.DeleteAudience(id)
}

func DeleteChart(id int) (int64, error) {
	return repositories.DeleteChart(id)
}

func DeleteInsight(id int) (int64, error) {
	return repositories.DeleteInsight(id)
}
