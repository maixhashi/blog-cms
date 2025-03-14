package model

import "time"

type LayoutComponent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	Content   string    `json:"content"`
	X         int       `json:"x" gorm:"default:0"` // X座標位置
	Y         int       `json:"y" gorm:"default:0"` // Y座標位置
	Width     int       `json:"width" gorm:"default:100"` // 幅
	Height    int       `json:"height" gorm:"default:100"` // 高さ
	UserId    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId"`
	LayoutId  *uint     `json:"layout_id"` // nullableにする
	Layout    *Layout   `json:"layout" gorm:"foreignKey:LayoutId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LayoutComponentResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	LayoutId  *uint     `json:"layout_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
