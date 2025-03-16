package feed_test

import (
	"go-react-app/model"
	"testing"
)

func TestFeedUsecase_GetAllFeeds(t *testing.T) {
	setupFeedTest()
	
	// テストデータの作成
	feeds := []model.Feed{
		{Title: "Feed 1", URL: "https://example.com/feed1", UserId: feedTestUser.ID},
		{Title: "Feed 2", URL: "https://example.com/feed2", UserId: feedTestUser.ID},
		{Title: "Feed 3", URL: "https://example.com/feed3", UserId: feedOtherUser.ID}, // 別ユーザーのフィード
	}
	
	for _, feed := range feeds {
		feedDB.Create(&feed)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのフィードのみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d のフィードを取得します", feedTestUser.ID)
			
			feedResponses, err := feedUsecase.GetAllFeeds(feedTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllFeeds() error = %v", err)
			}
			
			if len(feedResponses) != 2 {
				t.Errorf("GetAllFeeds() got %d feeds, want 2", len(feedResponses))
			}
			
			// フィードタイトルの確認
			titles := make(map[string]bool)
			for _, feed := range feedResponses {
				titles[feed.Title] = true
				t.Logf("取得したフィード: ID=%d, Title=%s", feed.ID, feed.Title)
			}
			
			if !titles["Feed 1"] || !titles["Feed 2"] {
				t.Errorf("期待したフィードが結果に含まれていません: %v", feedResponses)
			}
			
			// レスポンス形式の検証
			for _, feed := range feedResponses {
				if feed.ID == 0 || feed.Title == "" || feed.CreatedAt.IsZero() || feed.UpdatedAt.IsZero() {
					t.Errorf("GetAllFeeds() returned invalid feed: %+v", feed)
				}
			}
		})
	})
}

func TestFeedUsecase_GetFeedById(t *testing.T) {
	setupFeedTest()
	
	// テストデータの作成
	feed := createTestFeed("Test Feed", "https://example.com/test", feedTestUser.ID)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するフィードを正しく取得する", func(t *testing.T) {
			t.Logf("フィードID %d を取得します", feed.ID)
			
			response, err := feedUsecase.GetFeedById(feedTestUser.ID, feed.ID)
			
			if err != nil {
				t.Errorf("GetFeedById() error = %v", err)
			}
			
			if validateFeedResponse(t, response, feed.Title, feed.URL) {
				t.Logf("正常に取得: ID=%d, Title=%s", response.ID, response.Title)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してフィードを取得しようとします", nonExistentFeedID)
			
			_, err := feedUsecase.GetFeedById(feedTestUser.ID, nonExistentFeedID)
			
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーのフィードは取得できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := createTestFeed("Other User's Feed", "https://example.com/other", feedOtherUser.ID)
			t.Logf("他ユーザーのフィード(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserFeed.ID, feedTestUser.ID)
			
			_, err := feedUsecase.GetFeedById(feedTestUser.ID, otherUserFeed.ID)
			
			if err == nil {
				t.Error("他のユーザーのフィードを取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
