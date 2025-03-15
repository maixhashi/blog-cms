package feed_article_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFeedArticleRepository_GetArticlesByFeedID(t *testing.T) {
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
		articles, err := feedArticleRepo.GetArticlesByFeedID(feedArticleTestUser.ID, 1)
		
		// 検証
		assert.NoError(t, err)
		assert.Len(t, articles, 2) // XMLに2つの記事が含まれている
		
		// 最初の記事の内容を検証
		assert.Equal(t, "article1", articles[0].ID)
		assert.Equal(t, uint(1), articles[0].FeedID)
		assert.Equal(t, "Test Article 1", articles[0].Title)
		assert.Equal(t, "https://example.com/article1", articles[0].URL)
		assert.Equal(t, "Summary of article 1", articles[0].Summary)
		assert.Contains(t, articles[0].Categories, "Tech")
		assert.Equal(t, "Test Author", articles[0].Author)
		
		// 2番目の記事の内容を検証
		assert.Equal(t, "article2", articles[1].ID)
		assert.Equal(t, "Test Article 2", articles[1].Title)
		
		// モックが期待通り呼ばれたことを確認
		mockFeedRepo.AssertExpectations(t)
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("フィード取得エラー", func(t *testing.T) {
			// GetFeedByIdでエラーを返すようにモック設定
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, uint(999)).
				Return(createMockFeed(999, ""), assert.AnError)
			
			// テスト対象の関数を実行
			_, err := feedArticleRepo.GetArticlesByFeedID(feedArticleTestUser.ID, 999)
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "フィードの取得に失敗")
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
		
		t.Run("無効なXMLフォーマットのフィード", func(t *testing.T) {
			// 無効なXMLを返すサーバー
			invalidXMLServer := createInvalidXMLServer()
			defer invalidXMLServer.Close()
			
			// 新しいモックを作成
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = repository.NewFeedArticleRepository(mockFeedRepo)
			
			// GetFeedByIdのモック設定
			mockFeed := createMockFeed(1, invalidXMLServer.URL)
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, uint(1)).
				Run(func(args mock.Arguments) {
					feed := args.Get(0).(*model.Feed)
					*feed = mockFeed
				}).
				Return(mockFeed, nil)
			
			// テスト対象の関数を実行
			_, err := feedArticleRepo.GetArticlesByFeedID(feedArticleTestUser.ID, 1)
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "XMLのパース")
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
	})
}
