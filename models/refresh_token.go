package model

import "time"

type RefreshToken struct {
	Id        string `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Token     string `gorm:"unique"`
	UserId    string
	ExpiresAt time.Time
	CreatedAt time.Time
}
