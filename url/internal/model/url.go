package model

import "time"

type Url struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"not null;unique;index"`
	ShortUrl  string    `json:"short_url" gorm:"not null;unique;index"`
	LongUrl   string    `json:"long_url" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserId    uint      `json:"user_id" gorm:"not null;index"`
}
