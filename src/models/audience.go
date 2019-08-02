package models

// Audience models audience in database
type Audience struct {
	ID                     int    `json:"id"`
	UserID                 int    `json:"user_id"`
	Gender                 string `json:"gender"`
	BirthCountry           string `json:"birth_country"`
	AgeGroups              string `json:"age_groups"`
	HoursSpent             int    `json:"hours_spent"`
	NumOfPurchasesPerMonth int    `json:"num_of_purchases_per_month"`
}
