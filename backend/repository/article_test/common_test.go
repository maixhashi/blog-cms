package article_test

import (
    "go-react-app/model"
    "testing"
)

// ArticleRequestからテスト用記事を作成するヘルパー関数
func createTestArticle(title string, content string, userId uint) *model.Article {
    request := model.ArticleRequest{
        Title:   title,
        Content: content,
        UserId:  userId,
    }
    
    article := request.ToModel()
    articleDB.Create(&article)
    return &article
}

// 記事の検証ヘルパー関数
func validateArticle(t *testing.T, article *model.Article) {
    if article.ID == 0 {
        t.Error("Article ID should not be zero")
    }
    if article.CreatedAt.IsZero() {
        t.Error("CreatedAt should not be zero")
    }
    if article.UpdatedAt.IsZero() {
        t.Error("UpdatedAt should not be zero")
    }
}

// レスポンスの検証ヘルパー関数
func validateArticleResponse(t *testing.T, response model.ArticleResponse, expectedTitle string, expectedContent string) {
    if response.ID == 0 {
        t.Error("Response ID should not be zero")
    }
    if response.Title != expectedTitle {
        t.Errorf("Expected title %s, got %s", expectedTitle, response.Title)
    }
    if response.Content != expectedContent {
        t.Errorf("Expected content %s, got %s", expectedContent, response.Content)
    }
    if response.CreatedAt.IsZero() {
        t.Error("CreatedAt should not be zero")
    }
    if response.UpdatedAt.IsZero() {
        t.Error("UpdatedAt should not be zero")
    }
}