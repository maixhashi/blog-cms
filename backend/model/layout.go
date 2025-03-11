package model

import "time"

type Layout struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User      User      `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type LayoutResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
