package services

import (
	"platform2.0-go-challenge/datalayer/repositories"
	"platform2.0-go-challenge/models"
	"platform2.0-go-challenge/servicelayer/validators"
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

func AddAudience(audience models.Audience) (int, error) {
	err := validators.ValidateAudience(audience)
	if err != nil {
		return 0, err
	}

	return repositories.AddAudience(audience)
}

func AddChart(chart models.Chart) (int, error) {
	err := validators.ValidateChart(chart)
	if err != nil {
		return 0, err
	}

	return repositories.AddChart(chart)
}

func AddInsight(insight models.Insight) (int, error) {
	err := validators.ValidateInsight(insight)
	if err != nil {
		return 0, err
	}

	return repositories.AddInsight(insight)
}

func EditAudience(audience models.Audience) (int64, error) {
	err := validators.ValidateAudience(audience)
	if err != nil {
		return 0, err
	}

	return repositories.EditAudience(audience)
}

func EditChart(chart models.Chart) (int64, error) {
	err := validators.ValidateChart(chart)
	if err != nil {
		return 0, err
	}

	return repositories.EditChart(chart)
}

func EditInsight(insight models.Insight) (int64, error) {
	err := validators.ValidateInsight(insight)
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
