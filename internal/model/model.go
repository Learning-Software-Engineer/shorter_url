package model

import "time"

type User struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type URLMapping struct {
	ShortURL       string    `json:"short_url"`
	LongURL        string    `json:"long_url"`
	CreationDate   time.Time `json:"creation_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	UserID         int       `json:"user_id"`
	ClickCount     int       `json:"click_count"`
}
