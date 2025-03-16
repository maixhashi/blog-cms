package article_test

import (
	"go-react-app/model"
	"testing"
)

func TestArticleUsecase_GetAllArticles(t *testing.T) {
	setupArticleUsecaseTest()
	
	// テストデータの作成
	articles := []model.Article{
		{Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
		{Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
		{Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID}, // 別ユーザーの記事
	}
	
	for _, article := range articles {
		articleDb.Create(&article)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDの記事のみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d の記事を取得します", articleTestUser.ID)
			
			articleResponses, err := articleUsecase.GetAllArticles(articleTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if len(articleResponses) != 2 {
				t.Errorf("GetAllArticles() got %d articles, want 2", len(articleResponses))
			}
			
			// 記事タイトルの確認
			titles := make(map[string]bool)
			for _, article := range articleResponses {
				titles[article.Title] = true
				t.Logf("取得した記事: ID=%d, Title=%s", article.ID, article.Title)
			}
			
			if !titles["Article 1"] || !titles["Article 2"] {
				t.Errorf("期待した記事が結果に含まれていません: %v", articleResponses)
			}
			
			// レスポンス形式の検証
			for _, article := range articleResponses {
				if article.ID == 0 || article.Title == "" || article.CreatedAt.IsZero() || article.UpdatedAt.IsZero() {
					t.Errorf("GetAllArticles() returned invalid article: %+v", article)
				}
			}
		})
	})
}

func TestArticleUsecase_GetArticleById(t *testing.T) {
	setupArticleUsecaseTest()
	
	// テストデータの作成
	article := createTestArticle(t, "Test Article", "Test Content", articleTestUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在する記事を正しく取得する", func(t *testing.T) {
			t.Logf("記事ID %d を取得します", article.ID)
			
			response, err := articleUsecase.GetArticleById(articleTestUser.ID, article.ID)
			
			if err != nil {
				t.Errorf("GetArticleById() error = %v", err)
			}
			
			verifyArticleResponse(t, response, article.Title, article.Content, article.UserId)
			t.Logf("正常に取得: ID=%d, Title=%s", response.ID, response.Title)
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定して記事を取得しようとします", nonExistentArticleID)
			
			_, err := articleUsecase.GetArticleById(articleTestUser.ID, nonExistentArticleID)
			
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーの記事は取得できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := createTestArticle(t, "Other User's Article", "Other User's Content", articleOtherUser.ID)
			t.Logf("他ユーザーの記事(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserArticle.ID, articleTestUser.ID)
			
			_, err := articleUsecase.GetArticleById(articleTestUser.ID, otherUserArticle.ID)
			
			if err == nil {
				t.Error("他のユーザーの記事を取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
