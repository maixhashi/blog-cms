package feed_article_test

import (
	"go-react-app/model"
	"reflect"
	"testing"
)

func TestFeedArticleUsecase_GetArticlesByFeedID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("フィードの記事を正しく取得できる", func(t *testing.T) {
			// テスト実行
			responses, err := feedArticleUc.GetArticlesByFeedID(testUserId, 1)
			
			// 検証
			if err != nil {
				t.Errorf("GetArticlesByFeedID() error = %v", err)
			}
			
			if len(responses) != 2 {
				t.Errorf("GetArticlesByFeedID() got %d articles, want 2", len(responses))
			}
			
			// レスポンスの確認
			expectedFirst := model.FeedArticleResponse{
				ID:          testArticles[0].ID,
				FeedID:      testArticles[0].FeedID,
				Title:       testArticles[0].Title,
				URL:         testArticles[0].URL,
				Summary:     testArticles[0].Summary,
				Categories:  testArticles[0].Categories,
				PublishedAt: testArticles[0].PublishedAt,
				Author:      testArticles[0].Author,
			}
			
			if !reflect.DeepEqual(responses[0], expectedFirst) {
				t.Errorf("GetArticlesByFeedID() response mismatch: got %+v, want %+v", responses[0], expectedFirst)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("フィードの取得に失敗した場合はエラーを返す", func(t *testing.T) {
			// モックリポジトリにエラーを返すよう設定
			mockRepo.shouldReturnErr = true
			
			// テスト実行
			_, err := feedArticleUc.GetArticlesByFeedID(testUserId, 1)
			
			// 検証
			if err == nil {
				t.Error("GetArticlesByFeedID() should return error")
			}
			
			// モックをリセット
			mockRepo.shouldReturnErr = false
		})
	})
}
