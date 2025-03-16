package article_test

import (
	"go-react-app/model"
	"testing"
)

// 記事の検証ヘルパー関数
func verifyArticleResponse(t *testing.T, response model.ArticleResponse, expectedTitle string, expectedContent string, expectedUserId uint) {
	if response.ID == 0 || response.Title != expectedTitle || response.Content != expectedContent {
		t.Errorf("記事の内容が期待値と一致しません: got=%+v, want title=%s, content=%s", 
			response, expectedTitle, expectedContent)
	}
	
	// タイムスタンプの検証
	if response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
		t.Error("記事のタイムスタンプが正しく設定されていません")
	}
}

// データベースの記事を検証するヘルパー関数
func verifyDatabaseArticle(t *testing.T, articleId uint, expectedTitle string, expectedContent string, expectedUserId uint) {
	var dbArticle model.Article
	result := articleDb.First(&dbArticle, articleId)
	
	if result.Error != nil {
		t.Errorf("データベースから記事を取得できませんでした: %v", result.Error)
		return
	}
	
	if dbArticle.Title != expectedTitle || dbArticle.Content != expectedContent || dbArticle.UserId != expectedUserId {
		t.Errorf("データベースの記事が期待値と一致しません: got=%+v, want title=%s, content=%s, userId=%d", 
			dbArticle, expectedTitle, expectedContent, expectedUserId)
	}
}

// 記事が存在するか確認するヘルパー関数
func articleExists(articleId uint) bool {
	var count int64
	articleDb.Model(&model.Article{}).Where("id = ?", articleId).Count(&count)
	return count > 0
}
