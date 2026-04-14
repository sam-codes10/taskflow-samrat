package models

type UserAuthRes struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

