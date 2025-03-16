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

// リクエスト用の構造体
type ArticleRequest struct {
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	Tags      string `json:"tags"`
	UserId    uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// レスポンス用の構造体
type ArticleResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
