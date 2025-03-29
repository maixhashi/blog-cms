package model

import "time"

const (
	BookTitleMaxLength = 200
	BookAuthorMaxLength = 100
)

// Book データベースの書籍モデル
type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`
	Title       string    `json:"title" gorm:"not null" example:"Go言語による並行処理"`
	Author      string    `json:"author" gorm:"not null" example:"Katherine Cox-Buday"`
	Description string    `json:"description" example:"Go言語の並行処理について解説した書籍"`
	ISBN        string    `json:"isbn" gorm:"index" example:"9784873118468"`
	ImageURL    string    `json:"image_url" example:"http://books.google.com/books/content?id=..."`
	PublishedDate string  `json:"published_date" example:"2018-06-15"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	User        User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId      uint      `json:"user_id" gorm:"not null" example:"1"`
}

// BookRequest 書籍作成・更新リクエスト
type BookRequest struct {
	Title        string `json:"title" validate:"required,max=200" example:"Go言語による並行処理"`
	Author       string `json:"author" validate:"required,max=100" example:"Katherine Cox-Buday"`
	Description  string `json:"description" example:"Go言語の並行処理について解説した書籍"`
	ISBN         string `json:"isbn" example:"9784873118468"`
	ImageURL     string `json:"image_url" example:"http://books.google.com/books/content?id=..."`
	PublishedDate string `json:"published_date" example:"2018-06-15"`
	UserId       uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// BookResponse 書籍のレスポンス
type BookResponse struct {
	ID           uint      `json:"id" example:"1"`
	Title        string    `json:"title" example:"Go言語による並行処理"`
	Author       string    `json:"author" example:"Katherine Cox-Buday"`
	Description  string    `json:"description" example:"Go言語の並行処理について解説した書籍"`
	ISBN         string    `json:"isbn" example:"9784873118468"`
	ImageURL     string    `json:"image_url" example:"http://books.google.com/books/content?id=..."`
	PublishedDate string    `json:"published_date" example:"2018-06-15"`
	CreatedAt    time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// ToResponse BookからBookResponseへの変換メソッド
func (b *Book) ToResponse() BookResponse {
	return BookResponse{
		ID:           b.ID,
		Title:        b.Title,
		Author:       b.Author,
		Description:  b.Description,
		ISBN:         b.ISBN,
		ImageURL:     b.ImageURL,
		PublishedDate: b.PublishedDate,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}

// ToModel BookRequestからBookへの変換メソッド
func (br *BookRequest) ToModel() Book {
	return Book{
		Title:        br.Title,
		Author:       br.Author,
		Description:  br.Description,
		ISBN:         br.ISBN,
		ImageURL:     br.ImageURL,
		PublishedDate: br.PublishedDate,
		UserId:       br.UserId,
	}
}
