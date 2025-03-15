package feed_test

import (
    "go-react-app/model"
    "testing"
)

func createTestFeed(title string, url string, userId uint) *model.Feed {
    feed := &model.Feed{
        Title:  title,
        URL:    url,
        UserId: userId,
    }
    feedDB.Create(feed)
    return feed
}

func validateFeed(t *testing.T, feed *model.Feed) {
    if feed.ID == 0 {
        t.Error("Feed ID should not be zero")
    }
    if feed.CreatedAt.IsZero() {
        t.Error("CreatedAt should not be zero")
    }
    if feed.UpdatedAt.IsZero() {
        t.Error("UpdatedAt should not be zero")
    }
}
