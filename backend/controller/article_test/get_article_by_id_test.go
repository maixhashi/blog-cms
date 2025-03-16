package article_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestArticleController_GetArticleById(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在する記事を正しく取得する", func(t *testing.T) {
			// テスト用記事の作成
			article := createTestArticle("Test Article", "Test Content", articleTestUser.ID)
			
			// テスト実行
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				article.ID, 
				http.MethodGet, 
				fmt.Sprintf("/articles/%d", article.ID), 
				"",
			)
			err := articleController.GetArticleById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetArticleById() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetArticleById() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseArticleResponse(t, rec.Body.Bytes())
			
			if response.ID != article.ID || response.Title != article.Title || response.Content != article.Content {
				t.Errorf("GetArticleById() = %v, want id=%d, title=%s, content=%s", response, article.ID, article.Title, article.Content)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーの記事は取得できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := createTestArticle("Other User's Article", "Other User's Content", articleOtherUser.ID)
			
			// テスト実行 - testUserとして他のユーザーの記事にアクセス
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				otherUserArticle.ID, 
				http.MethodGet, 
				fmt.Sprintf("/articles/%d", otherUserArticle.ID), 
				"",
			)
			err := articleController.GetArticleById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetArticleById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetArticleById() with other user's article status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			// テスト実行 - 存在しないIDの記事にアクセス
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				nonExistentArticleID, 
				http.MethodGet, 
				fmt.Sprintf("/articles/%d", nonExistentArticleID), 
				"",
			)
			err := articleController.GetArticleById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetArticleById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetArticleById() with non-existent ID status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}
