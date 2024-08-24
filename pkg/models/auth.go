package models

import "time"

type AuthToken struct {
	ID int `json:"id"`
	UserId int `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
