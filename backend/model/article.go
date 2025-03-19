package model

import "time"

// データベースモデル
type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:"type:text"`
	Published bool      `json:"published" gorm:"default:false"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User      User      `json:"-" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

// ArticleRequest 記事作成・更新リクエスト
type ArticleRequest struct {
	Title     string `json:"title" validate:"required" example:"Goプログラミングの基礎"`
	Content   string `json:"content" example:"Goは静的型付け言語です..."`
	Published bool   `json:"published" example:"true"`
	Tags      string `json:"tags" example:"Go,プログラミング,チュートリアル"`
	UserId    uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// ArticleResponse 記事のレスポンス
type ArticleResponse struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"Goプログラミングの基礎"`
	Content   string    `json:"content" example:"Goは静的型付け言語です..."`
	Published bool      `json:"published" example:"true"`
	Tags      string    `json:"tags" example:"Go,プログラミング,チュートリアル"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}
// ArticleからArticleResponseへの変換メソッド
func (a *Article) ToResponse() ArticleResponse {
	return ArticleResponse{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		Published: a.Published,
		Tags:      a.Tags,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

// ArticleRequestからArticleへの変換メソッド
func (ar *ArticleRequest) ToModel() Article {
	return Article{
		Title:     ar.Title,
		Content:   ar.Content,
		Published: ar.Published,
		Tags:      ar.Tags,
		UserId:    ar.UserId,
	}
}
