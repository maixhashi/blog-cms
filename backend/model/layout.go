package model

import "time"

type Layout struct {
	ID         uint              `json:"id" gorm:"primaryKey"`
	Title      string            `json:"title" gorm:"not null"`
	UserId     uint              `json:"user_id" gorm:"not null"`
	User       User              `json:"user" gorm:"foreignKey:UserId"`
	Components []LayoutComponent `json:"components" gorm:"foreignKey:LayoutId"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type LayoutResponse struct {
	ID         uint                     `json:"id"`
	Title      string                   `json:"title"`
	Components []LayoutComponentResponse `json:"components,omitempty"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}