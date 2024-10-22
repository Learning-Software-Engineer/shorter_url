package dto

import "time"

type ShortenRequest struct {
	LongURL     string    `json:"long_url"`
	CustomAlias string    `json:"custom_alias,omitempty"`
	ExpiryDate  time.Time `json:"expiry_date"`
	UserID      int       `json:"user_id"`
}

type ShortenResponse struct {
	ShortURL   string    `json:"short_url"`
	LongURL    string    `json:"long_url"`
	ExpiryDate time.Time `json:"expiry_date"`
	CreatedAt  time.Time `json:"created_at"`
}

type RedirectRequest struct {
	ShortURL string `json:"short_url"`
}

type RedirectResponse struct {
	RedirectURL string `json:"redirect_url"`
}
