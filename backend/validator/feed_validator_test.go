package validator

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"
)

func TestFeedValidate(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewFeedValidator()

	testCases := []struct {
		name     string
		feed     model.Feed
		hasError bool
	}{
		{
			name: "Valid feed with valid title",
			feed: model.Feed{
				Title:  "Valid Feed Title",
				URL:    "https://example.com/feed",
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty title",
			feed: model.Feed{
				Title:  "",
				URL:    "https://example.com/feed",
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Valid title with URL",
			feed: model.Feed{
				Title:  "Feed Title",
				URL:    "https://example.com/feed",
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Zero user ID",
			feed: model.Feed{
				Title:  "Valid Title",
				URL:    "https://example.com/feed",
				UserId: 0,
			},
			hasError: false, // UserIDはバリデーションしていないので、エラーにならないはず
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.FeedValidate(tc.feed)
			if (err != nil) != tc.hasError {
				t.Errorf("FeedValidate() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}
