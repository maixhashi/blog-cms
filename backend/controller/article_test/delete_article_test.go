package article_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestArticleController_DeleteArticle(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分の記事を削除できる", func(t *testing.T) {
			// テスト用記事の作成
			article := createTestArticle("Article to Delete", "Content to Delete", articleTestUser.ID)
			
			// テスト実行
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				article.ID, 
				http.MethodDelete, 
				fmt.Sprintf("/articles/%d", article.ID), 
				"",
			)
			err := articleController.DeleteArticle(c)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteArticle() error = %v", err)
			}
			
			if rec.Code != http.StatusNoContent {
				t.Errorf("DeleteArticle() status code = %d, want %d", rec.Code, http.StatusNoContent)
			}
			
			// データベースから削除されていることを確認
			if articleExists(article.ID) {
				t.Error("DeleteArticle() did not delete the article from database")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーの記事は削除できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := createTestArticle("Other User's Article to Not Delete", "Other User's Content", articleOtherUser.ID)
			
			// テスト実行 - testUserとして他のユーザーの記事を削除しようとする
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				otherUserArticle.ID, 
				http.MethodDelete, 
				fmt.Sprintf("/articles/%d", otherUserArticle.ID), 
				"",
			)
			err := articleController.DeleteArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteArticle() with other user's article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
			
			// データベースから削除されていないことを確認
			if !articleExists(otherUserArticle.ID) {
				t.Error("DeleteArticle() deleted other user's article from database")
			}
		})
		
		t.Run("存在しない記事の削除はエラーになる", func(t *testing.T) {
			// テスト実行 - 存在しないIDの記事を削除しようとする
			_, c, rec := setupEchoWithArticleId(
				articleTestUser.ID, 
				nonExistentArticleID, 
				http.MethodDelete, 
				fmt.Sprintf("/articles/%d", nonExistentArticleID), 
				"",
			)
			err := articleController.DeleteArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteArticle() with non-existent article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("無効な記事IDパラメータでエラーを返す", func(t *testing.T) {
			// 無効なIDパラメータ
			_, c, rec := setupEchoWithJWTAndBody(
				articleTestUser.ID, 
				http.MethodDelete, 
				"/articles/invalid", 
				"",
			)
			c.SetParamNames("articleId")
			c.SetParamValues("invalid")
			
			err := articleController.DeleteArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusBadRequest {
				t.Errorf("DeleteArticle() with invalid ID parameter status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}