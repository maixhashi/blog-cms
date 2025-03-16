package article_test

import (
	"go-react-app/model"
	"net/http"
	"testing"
)

func TestArticleController_GetAllArticles(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーの記事を全て取得する", func(t *testing.T) {
			// テスト用記事の作成
			articles := []model.ArticleRequest{
				{Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
				{Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
				{Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID}, // 別ユーザーの記事
			}
			
			for _, articleReq := range articles {
				article := articleReq.ToModel()
				articleDB.Create(&article)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWT(articleTestUser.ID)
			err := articleController.GetAllArticles(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllArticles() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			response := parseArticlesResponse(t, rec.Body.Bytes())
			
			if len(response) != 2 {
				t.Errorf("GetAllArticles() returned %d articles, want 2", len(response))
			}
			
			// 記事タイトルの確認
			titles := make(map[string]bool)
			for _, article := range response {
				titles[article.Title] = true
			}
			
			if !titles["Article 1"] || !titles["Article 2"] {
				t.Errorf("期待した記事が結果に含まれていません: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// データベース接続エラーなどのケースをモックして追加可能
		// 現在の実装では直接テストできないため省略
	})
}
