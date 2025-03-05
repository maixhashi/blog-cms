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
		article  model.Article
		hasError bool
	}{
		{
			name: "Valid article with valid title",
			article: model.Article{
				Title:   "Valid Title",
				Content: "Some content",
				UserId:  user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty title",
			article: model.Article{
				Title:   "",
				Content: "Some content",
				UserId:  user.ID,
			},
			hasError: true,
		},
		{
			name: "Valid article with content",
			article: model.Article{
				Title:   "Valid Title",
				Content: "This is a valid content",
				UserId:  user.ID,
			},
			hasError: false,
		},
		{
			name: "Valid article with no content",
			article: model.Article{
				Title:   "Valid Title",
				Content: "",
				UserId:  user.ID,
			},
			hasError: false, // コンテンツは必須ではない
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ArticleValidate(tc.article)
			if (err != nil) != tc.hasError {
				t.Errorf("ArticleValidate() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// テスト用の長いタイトルを生成
func generateLongTitle(length int) string {
	return strings.Repeat("a", length)
}
