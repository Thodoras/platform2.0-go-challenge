package services

import (
	"platform2.0-go-challenge/src/datalayer/repositories"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/validators"
	"platform2.0-go-challenge/src/weblayer/dtos"
)

const numOfAssets = 3

func GetAllAssets(id int, onlyFavourites bool) (*dtos.AssetReponse, error) {
	var response dtos.AssetReponse
	errs := make(chan error)

	go getAudiencesAsync(id, onlyFavourites, &response, errs)
	go getChartsAsync(id, onlyFavourites, &response, errs)
	go getInsightsAsync(id, onlyFavourites, &response, errs)

	var err error
	for i := 0; i < numOfAssets; i++ {
		temp := <-errs
		if temp != nil {
			err = temp
		}
	}

	response.UserID = id
	return &response, err
}

func getAudiencesAsync(userID int, onlyFavourites bool, response *dtos.AssetReponse, errs chan error) {
	audiences, err := repositories.GetAudiences(userID, onlyFavourites)
	response.Audiences = audiences
	errs <- err
}

func getChartsAsync(userID int, onlyFavourites bool, response *dtos.AssetReponse, errs chan error) {
	charts, err := repositories.GetCharts(userID, onlyFavourites)
	response.Charts = charts
	errs <- err
}

func getInsightsAsync(userID int, onlyFavourites bool, response *dtos.AssetReponse, errs chan error) {
	insights, err := repositories.GetInsights(userID, onlyFavourites)
	response.Insights = insights
	errs <- err
}

func GetAllAssetsPaginated(userID, limit, offset int, onlyFavourites bool) (*dtos.AssetReponse, error) {
	var response dtos.AssetReponse
	errs := make(chan error)

	go getAudiencesPaginatedAsync(userID, limit, offset, onlyFavourites, &response, errs)
	go getChartsPaginatedAsync(userID, limit, offset, onlyFavourites, &response, errs)
	go getInsightsPaginatedAsync(userID, limit, offset, onlyFavourites, &response, errs)

	var err error
	for i := 0; i < numOfAssets; i++ {
		temp := <-errs
		if temp != nil {
			err = temp
		}
	}

	response.UserID = userID
	return &response, err
}

func getAudiencesPaginatedAsync(userID, limit, offset int, onlyFavourites bool, response *dtos.AssetReponse, errs chan error) {
	audiences, err := repositories.GetAudiencesPaginated(userID, limit, offset, onlyFavourites)
	response.Audiences = audiences
	errs <- err
}

func getChartsPaginatedAsync(userID, limit, offset int, onlyFavourites bool, response *dtos.AssetReponse, errs chan error) {
	charts, err := repositories.GetChartsPaginated(userID, limit, offset, onlyFavourites)
	response.Charts = charts
	errs <- err
}

func getInsightsPaginatedAsync(userID, limit, offset int, onlyFavourites bool, response *dtos.AssetReponse, errs chan error) {
	insights, err := repositories.GetInsightsPaginated(userID, limit, offset, onlyFavourites)
	response.Insights = insights
	errs <- err
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
