package models

// Audience models audience in database
type Audience struct {
	UserAsset
	Gender                 string `json:"gender"`
	BirthCountry           string `json:"birth_country"`
	AgeGroups              string `json:"age_groups"`
	HoursSpent             int    `json:"hours_spent"`
	NumOfPurchasesPerMonth int    `json:"num_of_purchases_per_month"`
}
