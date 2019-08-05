package models

type UserAsset struct {
	ID        int  `json:"id"`
	UserID    int  `json:"user_id"`
	Favourite bool `json:"favourite"`
}
