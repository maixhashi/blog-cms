package feed_test

import (
    "go-react-app/model"
    "testing"
)

func TestFeedRepository_DeleteFeed(t *testing.T) {
    setupFeedTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("自分のフィードを削除できる", func(t *testing.T) {
            feed := model.Feed{
                Title:  "Feed to Delete",
                URL:    "https://example.com/delete",
                UserId: feedTestUser.ID,
            }
            feedDB.Create(&feed)
            
            err := feedRepo.DeleteFeed(feedTestUser.ID, feed.ID)
            
            if err != nil {
                t.Errorf("DeleteFeed() error = %v", err)
            }
            
            var count int64
            feedDB.Model(&model.Feed{}).Where("id = ?", feed.ID).Count(&count)
            
            if count != 0 {
                t.Error("DeleteFeed() did not delete the feed from database")
            }
        })
    })
    
    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないフィードIDでの削除はエラーになる", func(t *testing.T) {
            err := feedRepo.DeleteFeed(feedTestUser.ID, nonExistentFeedID)
            
            if err == nil {
                t.Error("DeleteFeed() with non-existent ID should return error")
            }
        })
        
        t.Run("他のユーザーのフィードは削除できない", func(t *testing.T) {
            otherUserFeed := model.Feed{
                Title:  "Other User's Feed",
                URL:    "https://example.com/other",
                UserId: feedOtherUser.ID,
            }
            feedDB.Create(&otherUserFeed)
            
            err := feedRepo.DeleteFeed(feedTestUser.ID, otherUserFeed.ID)
            
            if err == nil {
                t.Error("DeleteFeed() should not allow deleting other user's feed")
            }
        })
    })
}
