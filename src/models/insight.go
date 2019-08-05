package models

// Insight models a insight in database
type Insight struct {
	UserAsset
	Text string `json:"text"`
}
