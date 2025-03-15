package feed_test

import (
    "go-react-app/model"
    "testing"
)

func TestFeedRepository_GetAllFeeds(t *testing.T) {
    setupFeedTest()
    
    feeds := []model.Feed{
        {Title: "Feed 1", URL: "https://example.com/feed1", UserId: feedTestUser.ID},
        {Title: "Feed 2", URL: "https://example.com/feed2", UserId: feedTestUser.ID},
        {Title: "Feed 3", URL: "https://example.com/feed3", UserId: feedOtherUser.ID},
    }
    
    for _, feed := range feeds {
        feedDB.Create(&feed)
    }
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("正しいユーザーIDのフィードのみを取得する", func(t *testing.T) {
            var result []model.Feed
            err := feedRepo.GetAllFeeds(&result, feedTestUser.ID)
            
            if err != nil {
                t.Errorf("GetAllFeeds() error = %v", err)
            }
            
            if len(result) != 2 {
                t.Errorf("GetAllFeeds() got %d feeds, want 2", len(result))
            }
            
            titles := make(map[string]bool)
            for _, feed := range result {
                titles[feed.Title] = true
            }
            
            if !titles["Feed 1"] || !titles["Feed 2"] {
                t.Errorf("期待したフィードが結果に含まれていません: %v", result)
            }
        })
    })
}
