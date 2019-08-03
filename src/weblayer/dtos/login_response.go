package dtos

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}
