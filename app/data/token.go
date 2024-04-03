package data

import (
	"database/sql"
	"time"
)

type Token struct {
	Id        uint      `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	AppId     string    `json:"app_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type TokenProvider struct {
	*sql.DB
}

func NewTokenProvider(db *sql.DB) *TokenProvider {
	return &TokenProvider{db}
}
