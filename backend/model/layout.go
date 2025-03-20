package model

import "time"

// データベースモデル
type Layout struct {
	ID         uint              `json:"id" gorm:"primaryKey" example:"1"`
	Title      string            `json:"title" gorm:"not null" example:"ブログのメインレイアウト"`
	UserId     uint              `json:"user_id" gorm:"not null" example:"1"`
	User       User              `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Components []LayoutComponent `json:"-" gorm:"foreignKey:LayoutId"`
	CreatedAt  time.Time         `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt  time.Time         `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// リクエスト用の構造体
type LayoutRequest struct {
	Title  string `json:"title" validate:"required" example:"ブログのメインレイアウト"`
	UserId uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// レスポンス用の構造体
type LayoutResponse struct {
	ID         uint                     `json:"id" example:"1"`
	Title      string                   `json:"title" example:"ブログのメインレイアウト"`
	Components []LayoutComponentResponse `json:"components,omitempty"`
	CreatedAt  time.Time                `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt  time.Time                `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}
// LayoutからLayoutResponseへの変換メソッド
func (l *Layout) ToResponse() LayoutResponse {
	response := LayoutResponse{
		ID:        l.ID,
		Title:     l.Title,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
	
	// Componentsがロードされている場合のみ変換
	if len(l.Components) > 0 {
		components := make([]LayoutComponentResponse, len(l.Components))
		for i, component := range l.Components {
			components[i] = component.ToResponse()
		}
		response.Components = components
	}
	
	return response
}

// LayoutRequestからLayoutへの変換メソッド
func (lr *LayoutRequest) ToModel() Layout {
	return Layout{
		Title:  lr.Title,
		UserId: lr.UserId,
	}
}