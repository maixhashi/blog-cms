package feed_test

import (
    "go-react-app/model"
    "testing"
)

func TestFeedRepository_CreateFeed(t *testing.T) {
    setupFeedTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("新しいフィードを作成できる", func(t *testing.T) {
            feed := model.Feed{
                Title:  "New Feed",
                URL:    "https://example.com/new",
                UserId: feedTestUser.ID,
            }
            
            err := feedRepo.CreateFeed(&feed)
            
            if err != nil {
                t.Errorf("CreateFeed() error = %v", err)
            }
            
            validateFeed(t, &feed)
            
            var savedFeed model.Feed
            feedDB.First(&savedFeed, feed.ID)
            
            if savedFeed.Title != "New Feed" || savedFeed.URL != "https://example.com/new" {
                t.Errorf("CreateFeed() = %v, want title=%s, url=%s", savedFeed, "New Feed", "https://example.com/new")
            }
        })
    })
}
