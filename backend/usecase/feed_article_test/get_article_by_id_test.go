package feed_article_test

import (
	"go-react-app/model"
	"reflect"
	"testing"
)

func TestFeedArticleUsecase_GetArticleByID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した記事IDの記事を正しく取得できる", func(t *testing.T) {
			// テスト実行
			response, err := feedArticleUc.GetArticleByID(testUserId, 1, "article1")
			
			// 検証
			if err != nil {
				t.Errorf("GetArticleByID() error = %v", err)
			}
			
			expectedResponse := model.FeedArticleResponse{
				ID:          testArticles[0].ID,
				FeedID:      testArticles[0].FeedID,
				Title:       testArticles[0].Title,
				URL:         testArticles[0].URL,
				Summary:     testArticles[0].Summary,
				Categories:  testArticles[0].Categories,
				PublishedAt: testArticles[0].PublishedAt,
				Author:      testArticles[0].Author,
			}
			
			if !reflect.DeepEqual(response, expectedResponse) {
				t.Errorf("GetArticleByID() response mismatch: got %+v, want %+v", response, expectedResponse)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しない記事IDを指定した場合はエラーを返す", func(t *testing.T) {
			// テスト実行
			_, err := feedArticleUc.GetArticleByID(testUserId, 1, "non-existent-id")
			
			// 検証
			if err == nil {
				t.Error("GetArticleByID() should return error for non-existent article ID")
			}
		})
		
		t.Run("フィードの取得に失敗した場合はエラーを返す", func(t *testing.T) {
			// モックリポジトリにエラーを返すよう設定
			mockRepo.shouldReturnErr = true
			
			// テスト実行
			_, err := feedArticleUc.GetArticleByID(testUserId, 1, "article1")
			
			// 検証
			if err == nil {
				t.Error("GetArticleByID() should return error when feed fetch fails")
			}
			
			// モックをリセット
			mockRepo.shouldReturnErr = false
		})
	})
}
