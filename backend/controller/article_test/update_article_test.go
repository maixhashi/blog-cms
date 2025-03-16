package article_test

import (
	"fmt"
	"go-react-app/model"  // model パッケージをインポート
	"net/http"
	"testing"
)

func TestArticleController_UpdateArticle(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("既存の記事を更新できる", func(t *testing.T) {
			// テスト用記事の作成
			article := createTestArticle("Original Article", "Original Content", articleTestUser.ID)
			
			// 更新リクエストの準備
			updateReqBody := `{"title":"Updated Article","content":"Updated Content"}`
			
			// テスト実行
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				article.ID, 
				http.MethodPut, 
				fmt.Sprintf("/articles/%d", article.ID), 
				updateReqBody,
			)
			err := articleController.UpdateArticle(c)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateArticle() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("UpdateArticle() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseArticleResponse(t, rec.Body.Bytes())
			
			if response.ID != article.ID || response.Title != "Updated Article" || response.Content != "Updated Content" {
				t.Errorf("UpdateArticle() = %v, want id=%d, title=%s, content=%s", 
					response, article.ID, "Updated Article", "Updated Content")
			}
			
			// データベースから直接確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, article.ID)
			if dbArticle.Title != "Updated Article" || dbArticle.Content != "Updated Content" {
				t.Errorf("UpdateArticle() did not update article correctly, got title=%s, content=%s", 
					dbArticle.Title, dbArticle.Content)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := createTestArticle("Other User's Article", "Other User's Content", articleOtherUser.ID)
			
			// 更新リクエストの準備
			updateReqBody := `{"title":"Attempted Update","content":"Attempted Content Update"}`
			
			// テスト実行 - testUserとして他のユーザーの記事を更新しようとする
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				otherUserArticle.ID, 
				http.MethodPut, 
				fmt.Sprintf("/articles/%d", otherUserArticle.ID), 
				updateReqBody,
			)
			err := articleController.UpdateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateArticle() with other user's article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
			
			// データベースに変更が反映されていないことを確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, otherUserArticle.ID)
			if dbArticle.Title != "Other User's Article" {
				t.Errorf("UpdateArticle() should not update other user's article, but got title=%s", dbArticle.Title)
			}
		})
		
		t.Run("存在しないIDの記事は更新できない", func(t *testing.T) {
			// 更新リクエストの準備
			updateReqBody := `{"title":"Update Non-existent","content":"Update Content"}`
			
			// テスト実行 - 存在しないIDの記事を更新しようとする
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				nonExistentArticleID, 
				http.MethodPut, 
				fmt.Sprintf("/articles/%d", nonExistentArticleID), 
				updateReqBody,
			)
			err := articleController.UpdateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateArticle() with non-existent article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("バリデーションエラーが発生する記事は更新できない", func(t *testing.T) {
			// テスト用記事の作成
			article := createTestArticle("Article for Validation Test", "Original Content", articleTestUser.ID)
			
			// タイトルが空の無効なリクエスト
			invalidReqBody := `{"title":"","content":"Content with empty title"}`
			
			// テスト実行
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				article.ID, 
				http.MethodPut, 
				fmt.Sprintf("/articles/%d", article.ID), 
				invalidReqBody,
			)
			err := articleController.UpdateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateArticle() with invalid article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
			
			// データベースに変更が反映されていないことを確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, article.ID)
			if dbArticle.Title != "Article for Validation Test" {
				t.Errorf("UpdateArticle() should not update article with validation error, but got title=%s", dbArticle.Title)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// テスト用記事の作成
			article := createTestArticle("Article for JSON Test", "Original Content", articleTestUser.ID)
			
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				article.ID, 
				http.MethodPut, 
				fmt.Sprintf("/articles/%d", article.ID), 
				invalidJSON,
			)
			err := articleController.UpdateArticle(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("UpdateArticle() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("UpdateArticle() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}
