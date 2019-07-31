package services

import (
	"platform2.0-go-challenge/datalayer/repositories"
	"platform2.0-go-challenge/models/assets"
)

func GetAllAssets(id string) error {
	return nil
}

func AddAudience(audience assets.Audience) (int, error) {
	return repositories.AddAudience(audience)
}

func AddChart(chart assets.Chart) (int, error) {
	return repositories.AddChart(chart)
}
