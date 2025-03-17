package model

import "time"

// データベースモデル
type LayoutComponent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	Content   string    `json:"content"`
	X         int       `json:"x" gorm:"default:0"`
	Y         int       `json:"y" gorm:"default:0"`
	Width     int       `json:"width" gorm:"default:100"`
	Height    int       `json:"height" gorm:"default:100"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	LayoutId  *uint     `json:"layout_id"`
	Layout    *Layout   `json:"-" gorm:"foreignKey:LayoutId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// リクエスト用の構造体
type LayoutComponentRequest struct {
	Name    string `json:"name" validate:"required"`
	Type    string `json:"type" validate:"required"`
	Content string `json:"content"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	UserId  uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// 位置情報更新用のリクエスト構造体
type PositionRequest struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

// レイアウト割り当て用のリクエスト構造体
type AssignLayoutRequest struct {
	LayoutId uint           `json:"layout_id" validate:"required"`
	Position PositionRequest `json:"position"`
}

// レスポンス用の構造体
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

// LayoutComponentからLayoutComponentResponseへの変換メソッド
func (lc *LayoutComponent) ToResponse() LayoutComponentResponse {
	return LayoutComponentResponse{
		ID:        lc.ID,
		Name:      lc.Name,
		Type:      lc.Type,
		Content:   lc.Content,
		X:         lc.X,
		Y:         lc.Y,
		Width:     lc.Width,
		Height:    lc.Height,
		LayoutId:  lc.LayoutId,
		CreatedAt: lc.CreatedAt,
		UpdatedAt: lc.UpdatedAt,
	}
}

// LayoutComponentRequestからLayoutComponentへの変換メソッド
func (lcr *LayoutComponentRequest) ToModel() LayoutComponent {
	return LayoutComponent{
		Name:    lcr.Name,
		Type:    lcr.Type,
		Content: lcr.Content,
		X:       lcr.X,
		Y:       lcr.Y,
		Width:   lcr.Width,
		Height:  lcr.Height,
		UserId:  lcr.UserId,
	}
}