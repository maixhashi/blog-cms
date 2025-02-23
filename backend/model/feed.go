package model

import "time"

type Feed struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Title         string     `json:"title" gorm:"not null"`           // フィードの名前 (例: "TechCrunch")
	URL           string     `json:"url" gorm:"not null;unique"`      // RSS/AtomのURL
	SiteURL       string     `json:"site_url"`                        // フィード元のサイトURL
	Description   string     `json:"description"`                     // フィードの説明・概要
	LastFetchedAt *time.Time `json:"last_fetched_at"`                 // 最後にフィードを取得した日時（NULL 許容）
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"` // 作成日時
	UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"` // 更新日時
	User          User       `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	UserId        uint       `json:"user_id" gorm:"not null"`
}

type FeedResponse struct {
	ID            uint       `json:"id"`
	Title         string     `json:"title"`
	URL           string     `json:"url"`
	SiteURL       string     `json:"site_url"`
	Description   string     `json:"description"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
