package feed_test

import (
    "go-react-app/model"
    "testing"
    "time"
)

func TestFeedRepository_UpdateFeed(t *testing.T) {
    setupFeedTest()
    
    feed := model.Feed{
        Title:  "Original Feed",
        URL:    "https://example.com/original",
        UserId: feedTestUser.ID,
    }
    feedDB.Create(&feed)
    
    time.Sleep(10 * time.Millisecond)
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("フィードのタイトルとURLを更新できる", func(t *testing.T) {
            updatedFeed := model.Feed{
                Title: "Updated Feed",
                URL:   "https://example.com/updated",
            }
            
            err := feedRepo.UpdateFeed(&updatedFeed, feedTestUser.ID, feed.ID)
            
            if err != nil {
                t.Errorf("UpdateFeed() error = %v", err)
            }
            
            var dbFeed model.Feed
            feedDB.First(&dbFeed, feed.ID)
            
            if dbFeed.Title != "Updated Feed" || dbFeed.URL != "https://example.com/updated" {
                t.Errorf("UpdateFeed() database feed not updated correctly")
            }
        })
    })

    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないフィードIDでの更新はエラーになる", func(t *testing.T) {
            invalidFeed := model.Feed{Title: "Invalid Update"}
            err := feedRepo.UpdateFeed(&invalidFeed, feedTestUser.ID, nonExistentFeedID)
            
            if err == nil {
                t.Error("UpdateFeed() should return error for non-existent ID")
            }
        })

        t.Run("他のユーザーのフィードは更新できない", func(t *testing.T) {
            otherUserFeed := model.Feed{
                Title:  "Other User's Feed",
                URL:    "https://example.com/other",
                UserId: feedOtherUser.ID,
            }
            feedDB.Create(&otherUserFeed)
            
            updateAttempt := model.Feed{Title: "Attempted Update"}
            err := feedRepo.UpdateFeed(&updateAttempt, feedTestUser.ID, otherUserFeed.ID)
            
            if err == nil {
                t.Error("UpdateFeed() should not allow updating other user's feed")
            }
        })
    })
}
