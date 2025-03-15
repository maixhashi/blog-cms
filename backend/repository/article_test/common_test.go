package article_test

import (
    "go-react-app/model"
    "testing"
)

func createTestArticle(title string, content string, userId uint) *model.Article {
    article := &model.Article{
        Title:   title,
        Content: content,
        UserId:  userId,
    }
    articleDB.Create(article)
    return article
}

func validateArticle(t testing.T, article *model.Article) {
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