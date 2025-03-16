package article_test

import (
	"encoding/json"
	"go-react-app/model"
	"testing"
)

// テスト用記事を作成するヘルパー関数
func createTestArticle(title string, content string, userId uint) model.Article {
	request := model.ArticleRequest{
		Title:   title,
		Content: content,
		UserId:  userId,
	}
	
	article := request.ToModel()
	articleDB.Create(&article)
	return article
}

// レスポンスボディをパースするヘルパー関数
func parseArticleResponse(t *testing.T, responseBody []byte) model.ArticleResponse {
	var response model.ArticleResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 複数記事のレスポンスボディをパースするヘルパー関数
func parseArticlesResponse(t *testing.T, responseBody []byte) []model.ArticleResponse {
	var response []model.ArticleResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 記事が存在するか確認するヘルパー関数
func articleExists(articleId uint) bool {
	var count int64
	articleDB.Model(&model.Article{}).Where("id = ?", articleId).Count(&count)
	return count > 0
}
