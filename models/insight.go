package models

// Insight models a insight in database
type Insight struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Text   string `json:"text"`
}
