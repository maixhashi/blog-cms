package article_test

import (
	"net/http"
	"testing"
)

func TestArticleController_CreateArticle(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しい記事を作成できる", func(t *testing.T) {
			// テストリクエストの準備
			reqBody := `{"title":"New Test Article","content":"This is a test article content"}`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(
				articleTestUser.ID, 
				http.MethodPost, 
				"/articles", 
				reqBody,
			)
			err := articleController.CreateArticle(c)
			
			// 検証
			if err != nil {
				t.Errorf("CreateArticle() error = %v", err)
			}
			
			if rec.Code != http.StatusCreated {
				t.Errorf("CreateArticle() status code = %d, want %d", rec.Code, http.StatusCreated)
			}
			
			// レスポンスボディをパース
			response := parseArticleResponse(t, rec.Body.Bytes())
			
			if response.Title != "New Test Article" || response.Content != "This is a test article content" {
				t.Errorf("CreateArticle() = %v, want title=%s, content=%s", response, "New Test Article", "This is a test article content")
			}
			
			// データベースから直接確認
			if !articleExists(response.ID) {
				t.Error("CreateArticle() did not save article to database")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する記事は作成できない", func(t *testing.T) {
			// タイトルが空のリクエスト
			reqBody := `{"title":"","content":"Invalid article with empty title"}`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(
				articleTestUser.ID, 
				http.MethodPost, 
				"/articles", 
				reqBody,
			)
			err := articleController.CreateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("CreateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("CreateArticle() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(
				articleTestUser.ID, 
				http.MethodPost, 
				"/articles", 
				invalidJSON,
			)
			err := articleController.CreateArticle(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("CreateArticle() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("CreateArticle() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}