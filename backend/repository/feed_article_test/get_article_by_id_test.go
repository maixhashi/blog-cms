package feed_article_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFeedArticleRepository_GetArticleByID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		// HTTPサーバーのモック作成
		server := createMockRSSServer()
		defer server.Close()
		
		// GetFeedByIdのモック設定
		mockFeed := createMockFeed(1, server.URL)
		mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, uint(1)).
			Run(func(args mock.Arguments) {
				feed := args.Get(0).(*model.Feed)
				*feed = mockFeed
			}).
			Return(mockFeed, nil)
		
		// テスト対象の関数を実行
		article, err := feedArticleRepo.GetArticleByID(feedArticleTestUser.ID, 1, "article1")
		
		// 検証
		assert.NoError(t, err)
		assert.Equal(t, "article1", article.ID)
		assert.Equal(t, uint(1), article.FeedID)
		assert.Equal(t, "Test Article 1", article.Title)
		
		// モックが期待通り呼ばれたことを確認
		mockFeedRepo.AssertExpectations(t)
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しない記事ID", func(t *testing.T) {
			// 新しいテストサーバーを作成
			server := createMockRSSServer()
			defer server.Close()
			
			// モックをリセットしてから、新しいモック設定を行う
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = repository.NewFeedArticleRepository(mockFeedRepo)
			
			// GetFeedByIdのモック設定
			mockFeed := createMockFeed(1, server.URL)
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, uint(1)).
				Run(func(args mock.Arguments) {
					feed := args.Get(0).(*model.Feed)
					*feed = mockFeed
				}).
				Return(mockFeed, nil)
			
			// 存在しない記事IDでテスト
			_, err := feedArticleRepo.GetArticleByID(feedArticleTestUser.ID, 1, "nonexistent")
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "記事が見つかりません")
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
	})
}
