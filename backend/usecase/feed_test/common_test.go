package feed_test

import (
	"go-react-app/model"
	"testing"
)

// フィードレスポンスの基本的な検証を行うヘルパー関数
func validateFeedResponse(t *testing.T, feed model.FeedResponse, expectedTitle string, expectedURL string) bool {
	if feed.ID == 0 || feed.Title != expectedTitle || feed.URL != expectedURL {
		t.Errorf("Feed response validation failed: got %+v, want title=%s, url=%s", feed, expectedTitle, expectedURL)
		return false
	}
	
	if feed.CreatedAt.IsZero() || feed.UpdatedAt.IsZero() {
		t.Errorf("Feed response has invalid timestamps: %+v", feed)
		return false
	}
	
	return true
}

// データベースのフィードを検証するヘルパー関数
func validateDatabaseFeed(t *testing.T, feedID uint, expectedTitle string, expectedURL string, expectedUserID uint) bool {
	var dbFeed model.Feed
	result := feedDB.First(&dbFeed, feedID)
	
	if result.Error != nil {
		t.Errorf("Failed to retrieve feed from database: %v", result.Error)
		return false
	}
	
	if dbFeed.Title != expectedTitle || dbFeed.URL != expectedURL || dbFeed.UserId != expectedUserID {
		t.Errorf("Database feed validation failed: got %+v, want title=%s, url=%s, userId=%d", 
			dbFeed, expectedTitle, expectedURL, expectedUserID)
		return false
	}
	
	return true
}
