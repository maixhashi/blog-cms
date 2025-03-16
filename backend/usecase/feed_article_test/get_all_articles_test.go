package feed_article_test

import (
	"testing"
)

func TestFeedArticleUsecase_GetAllArticles(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("すべての記事を正しく取得できる", func(t *testing.T) {
			// テスト実行
			responses, err := feedArticleUc.GetAllArticles(testUserId)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if len(responses) != 3 {
				t.Errorf("GetAllArticles() got %d articles, want 3", len(responses))
			}
			
			// すべての記事のIDを確認
			articleIDs := make(map[string]bool)
			for _, article := range responses {
				articleIDs[article.ID] = true
			}
			
			expectedIDs := []string{"article1", "article2", "article3"}
			for _, id := range expectedIDs {
				if !articleIDs[id] {
					t.Errorf("GetAllArticles() missing article with ID: %s", id)
				}
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("記事の取得に失敗した場合はエラーを返す", func(t *testing.T) {
			// モックリポジトリにエラーを返すよう設定
			mockRepo.getUserArticleErr = true
			
			// テスト実行
			_, err := feedArticleUc.GetAllArticles(testUserId)
			
			// 検証
			if err == nil {
				t.Error("GetAllArticles() should return error")
			}
			
			// モックをリセット
			mockRepo.getUserArticleErr = false
		})
	})
}
