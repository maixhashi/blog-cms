package model

// GoogleBook Google Books APIからのレスポンスモデル
type GoogleBook struct {
	ID           string   `json:"id" example:"zyTCAlFPjgYC"`
	Title        string   `json:"title" example:"Go言語による並行処理"`
	Authors      []string `json:"authors" example:"Katherine Cox-Buday"`
	Description  string   `json:"description" example:"Go言語の並行処理について解説した書籍"`
	ISBN         string   `json:"isbn" example:"9784873118468"`
	ImageURL     string   `json:"image_url" example:"http://books.google.com/books/content?id=..."`
	PublishedDate string   `json:"published_date" example:"2018-06-15"`
}

// GoogleBookSearchRequest 検索リクエスト
type GoogleBookSearchRequest struct {
	Query string `json:"query" validate:"required" example:"Go言語"`
	MaxResults int `json:"max_results" example:"10"`
}

// GoogleBookSearchResponse 検索結果レスポンス
type GoogleBookSearchResponse struct {
	TotalItems int          `json:"total_items" example:"42"`
	Items      []GoogleBook `json:"items"`
}

// ToBookRequest GoogleBookからBookRequestへの変換メソッド
func (gb *GoogleBook) ToBookRequest() BookRequest {
	author := ""
	if len(gb.Authors) > 0 {
		author = gb.Authors[0]
	}
	
	return BookRequest{
		Title:        gb.Title,
		Author:       author,
		Description:  gb.Description,
		ISBN:         gb.ISBN,
		ImageURL:     gb.ImageURL,
		PublishedDate: gb.PublishedDate,
	}
}
