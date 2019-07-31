package assets

import (
	"platform2.0-go-challenge/models/assets"
)

// AssetReponse models response of get assets
type AssetReponse struct {
	ID        int               `json:"id"`
	Audiences []assets.Audience `json:"audiences"`
	Charts    []assets.Chart    `json:"charts"`
	Insights  []assets.Insight  `json:"insights"`
}
