package dtos

import "platform2.0-go-challenge/src/models"

type AssetReponse struct {
	UserID    int               `json:"user_id"`
	Audiences []models.Audience `json:"audiences"`
	Charts    []models.Chart    `json:"charts"`
	Insights  []models.Insight  `json:"insights"`
}
