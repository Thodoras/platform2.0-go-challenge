package assetService

import (
	"platform2.0-go-challenge/src/datalayer/repositories/audienceRepository"
	"platform2.0-go-challenge/src/datalayer/repositories/chartRepository"
	"platform2.0-go-challenge/src/datalayer/repositories/insightRepository"
	"platform2.0-go-challenge/src/models"
	"platform2.0-go-challenge/src/servicelayer/validators/audienceValidator"
	"platform2.0-go-challenge/src/servicelayer/validators/chartValidator"
	"platform2.0-go-challenge/src/servicelayer/validators/insightValidator"
	"platform2.0-go-challenge/src/weblayer/dtos"
)

type AssetService struct {
	audienceRepository audienceRepository.IAudienceRepository
	chartRepository    chartRepository.IChartRepository
	insightRepository  insightRepository.IInsightRepository
	audienceValidator  audienceValidator.IAudienceValidator
	chartValidator     chartValidator.IChartValidator
	insightValidator   insightValidator.IInsightValidator
}

func NewAssetService(
	audienceRepository audienceRepository.IAudienceRepository,
	chartRepository chartRepository.IChartRepository,
	insightRepository insightRepository.IInsightRepository,
	audienceValidator audienceValidator.IAudienceValidator,
	chartValidator chartValidator.IChartValidator,
	insightValidator insightValidator.IInsightValidator,
) *AssetService {
	return &AssetService{
		audienceRepository: audienceRepository,
		chartRepository:    chartRepository,
		insightRepository:  insightRepository,
		audienceValidator:  audienceValidator,
		chartValidator:     chartValidator,
		insightValidator:   insightValidator,
	}

}

const numOfAssets = 3

func (a *AssetService) GetAllAssets(id int) (*dtos.AssetReponse, error) {
	var response dtos.AssetReponse
	errs := make(chan error)

	go a.getAudiencesAsync(id, &response, errs)
	go a.getChartsAsync(id, &response, errs)
	go a.getInsightsAsync(id, &response, errs)

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

func (a *AssetService) getAudiencesAsync(userID int, response *dtos.AssetReponse, errs chan error) {
	audiences, err := a.audienceRepository.GetAudiences(userID)
	response.Audiences = audiences
	errs <- err
}

func (a *AssetService) getChartsAsync(userID int, response *dtos.AssetReponse, errs chan error) {
	charts, err := a.chartRepository.GetCharts(userID)
	response.Charts = charts
	errs <- err
}

func (a *AssetService) getInsightsAsync(userID int, response *dtos.AssetReponse, errs chan error) {
	insights, err := a.insightRepository.GetInsights(userID)
	response.Insights = insights
	errs <- err
}

func (a *AssetService) GetAllAssetsPaginated(userID, limit, offset int) (*dtos.AssetReponse, error) {
	var response dtos.AssetReponse
	errs := make(chan error)

	go a.getAudiencesPaginatedAsync(userID, limit, offset, &response, errs)
	go a.getChartsPaginatedAsync(userID, limit, offset, &response, errs)
	go a.getInsightsPaginatedAsync(userID, limit, offset, &response, errs)

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

func (a *AssetService) getAudiencesPaginatedAsync(userID, limit, offset int, response *dtos.AssetReponse, errs chan error) {
	audiences, err := a.audienceRepository.GetAudiencesPaginated(userID, limit, offset)
	response.Audiences = audiences
	errs <- err
}

func (a *AssetService) getChartsPaginatedAsync(userID, limit, offset int, response *dtos.AssetReponse, errs chan error) {
	charts, err := a.chartRepository.GetChartsPaginated(userID, limit, offset)
	response.Charts = charts
	errs <- err
}

func (a *AssetService) getInsightsPaginatedAsync(userID, limit, offset int, response *dtos.AssetReponse, errs chan error) {
	insights, err := a.insightRepository.GetInsightsPaginated(userID, limit, offset)
	response.Insights = insights
	errs <- err
}

func (a *AssetService) AddAudience(audience models.Audience) (int, error) {
	err := a.audienceValidator.ValidateAudience(audience)
	if err != nil {
		return 0, err
	}

	return a.audienceRepository.AddAudience(audience)
}

func (a *AssetService) AddChart(chart models.Chart) (int, error) {
	err := a.chartValidator.ValidateChart(chart)
	if err != nil {
		return 0, err
	}

	return a.chartRepository.AddChart(chart)
}

func (a *AssetService) AddInsight(insight models.Insight) (int, error) {
	err := a.insightValidator.ValidateInsight(insight)
	if err != nil {
		return 0, err
	}

	return a.insightRepository.AddInsight(insight)
}

func (a *AssetService) EditAudience(audience models.Audience) (int64, error) {
	err := a.audienceValidator.ValidateAudience(audience)
	if err != nil {
		return 0, err
	}

	return a.audienceRepository.EditAudience(audience)
}

func (a *AssetService) EditChart(chart models.Chart) (int64, error) {
	err := a.chartValidator.ValidateChart(chart)
	if err != nil {
		return 0, err
	}

	return a.chartRepository.EditChart(chart)
}

func (a *AssetService) EditInsight(insight models.Insight) (int64, error) {
	err := a.insightValidator.ValidateInsight(insight)
	if err != nil {
		return 0, err
	}

	return a.insightRepository.EditInsight(insight)
}

func (a *AssetService) DeleteAudience(id int) (int64, error) {
	return a.audienceRepository.DeleteAudience(id)
}

func (a *AssetService) DeleteChart(id int) (int64, error) {
	return a.chartRepository.DeleteChart(id)
}

func (a *AssetService) DeleteInsight(id int) (int64, error) {
	return a.insightRepository.DeleteInsight(id)
}
