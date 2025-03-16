package article_test

import (
	"go-react-app/model"
	"testing"
)

func TestArticleUsecase_CreateArticle(t *testing.T) {
	setupArticleUsecaseTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しい記事を作成できる", func(t *testing.T) {
			// テスト用記事リクエスト
			validRequest := model.ArticleRequest{
				Title:   "New Test Article",
				Content: "This is a test article content",
				UserId:  articleTestUser.ID,
			}
			
			t.Logf("記事作成: Title=%s, UserId=%d", validRequest.Title, validRequest.UserId)
			
			// テスト実行
			response, err := articleUsecase.CreateArticle(validRequest)
			
			// 検証
			if err != nil {
				t.Errorf("CreateArticle() error = %v", err)
			}
			
			verifyArticleResponse(t, response, validRequest.Title, validRequest.Content, validRequest.UserId)
			t.Logf("生成された記事ID: %d", response.ID)
			
			// データベースから直接確認
			verifyDatabaseArticle(t, response.ID, validRequest.Title, validRequest.Content, articleTestUser.ID)
			t.Logf("データベース保存確認: Title=%s, Content=%s, UserId=%d", validRequest.Title, validRequest.Content, articleTestUser.ID)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する記事は作成できない", func(t *testing.T) {
			// 無効な記事（タイトルなし）
			invalidRequest := model.ArticleRequest{
				Title:   "", // 空のタイトル
				Content: "Invalid article content",
				UserId:  articleTestUser.ID,
			}
			
			t.Logf("無効な記事作成を試行: Title=%s (空)", invalidRequest.Title)
			
			_, err := articleUsecase.CreateArticle(invalidRequest)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効な記事でエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに保存されていないことを確認
			var count int64
			articleDb.Model(&model.Article{}).Where("content = ? AND title = ?", invalidRequest.Content, invalidRequest.Title).Count(&count)
			if count > 0 {
				t.Error("バリデーションエラーの記事がデータベースに保存されています")
			} else {
				t.Log("バリデーションエラーの記事は保存されていないことを確認")
			}
		})
	})
}