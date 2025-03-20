package model

import "time"

// データベースモデル
type LayoutComponent struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Name      string    `json:"name" gorm:"not null" example:"ヘッダーコンポーネント"`
	Type      string    `json:"type" gorm:"not null" example:"header"`
	Content   string    `json:"content" example:"<h1>ブログタイトル</h1>"`
	X         int       `json:"x" gorm:"default:0" example:"0"`
	Y         int       `json:"y" gorm:"default:0" example:"0"`
	Width     int       `json:"width" gorm:"default:100" example:"100"`
	Height    int       `json:"height" gorm:"default:100" example:"50"`
	UserId    uint      `json:"user_id" gorm:"not null" example:"1"`
	User      User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	LayoutId  *uint     `json:"layout_id" example:"1"`
	Layout    *Layout   `json:"-" gorm:"foreignKey:LayoutId"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// リクエスト用の構造体
type LayoutComponentRequest struct {
	Name    string `json:"name" validate:"required" example:"ヘッダーコンポーネント"`
	Type    string `json:"type" validate:"required" example:"header"`
	Content string `json:"content" example:"<h1>ブログタイトル</h1>"`
	X       int    `json:"x" example:"0"`
	Y       int    `json:"y" example:"0"`
	Width   int    `json:"width" example:"100"`
	Height  int    `json:"height" example:"50"`
	UserId  uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// 位置情報更新用のリクエスト構造体
type PositionRequest struct {
	X      int `json:"x" example:"10"`
	Y      int `json:"y" example:"20"`
	Width  int `json:"width,omitempty" example:"150"`
	Height int `json:"height,omitempty" example:"75"`
}

// レイアウト割り当て用のリクエスト構造体
type AssignLayoutRequest struct {
	LayoutId uint           `json:"layout_id" validate:"required" example:"1"`
	Position PositionRequest `json:"position"`
}

// レスポンス用の構造体
type LayoutComponentResponse struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"ヘッダーコンポーネント"`
	Type      string    `json:"type" example:"header"`
	Content   string    `json:"content" example:"<h1>ブログタイトル</h1>"`
	X         int       `json:"x" example:"0"`
	Y         int       `json:"y" example:"0"`
	Width     int       `json:"width" example:"100"`
	Height    int       `json:"height" example:"50"`
	LayoutId  *uint     `json:"layout_id,omitempty" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
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