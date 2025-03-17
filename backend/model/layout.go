package model

import "time"

// データベースモデル
type Layout struct {
	ID         uint              `json:"id" gorm:"primaryKey"`
	Title      string            `json:"title" gorm:"not null"`
	UserId     uint              `json:"user_id" gorm:"not null"`
	User       User              `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Components []LayoutComponent `json:"-" gorm:"foreignKey:LayoutId"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// リクエスト用の構造体
type LayoutRequest struct {
	Title  string `json:"title" validate:"required"`
	UserId uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// レスポンス用の構造体
type LayoutResponse struct {
	ID         uint                     `json:"id"`
	Title      string                   `json:"title"`
	Components []LayoutComponentResponse `json:"components,omitempty"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
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