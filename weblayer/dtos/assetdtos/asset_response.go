package assetdtos

import (
	"platform2.0-go-challenge/models/assets"
)

type AssetReponse struct {
	UserID    int               `json:"user_id"`
	Audiences []assets.Audience `json:"audiences"`
	Charts    []assets.Chart    `json:"charts"`
	Insights  []assets.Insight  `json:"insights"`
}
