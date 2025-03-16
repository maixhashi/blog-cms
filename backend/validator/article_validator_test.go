package validator

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"strings"
	"testing"
)

func TestArticleValidate(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewArticleValidator()

	testCases := []struct {
		name     string
		request  model.ArticleRequest
		hasError bool
	}{
		{
			name: "Valid article with valid title",
			request: model.ArticleRequest{
				Title:   "Valid Title",
				Content: "Some content",
				UserId:  user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty title",
			request: model.ArticleRequest{
				Title:   "",
				Content: "Some content",
				UserId:  user.ID,
			},
			hasError: true,
		},
		{
			name: "Valid article with content",
			request: model.ArticleRequest{
				Title:   "Valid Title",
				Content: "This is a valid content",
				UserId:  user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid article with no content",
			request: model.ArticleRequest{
				Title:   "Valid Title",
				Content: "",
				UserId:  user.ID,
			},
			hasError: false, // コンテンツは必須ではない
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateArticleRequest(tc.request)
			if (err != nil) != tc.hasError {
				t.Errorf("ValidateArticleRequest() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// テスト用の長いタイトルを生成
func generateLongTitle(length int) string {
	return strings.Repeat("a", length)
}
