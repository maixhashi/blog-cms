package feed_article_test

import (
	"errors"
	"go-react-app/model"
	"go-react-app/repository"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFeedArticleRepository_GetAllArticles(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		// HTTPサーバーのモック作成
		server := createMockRSSServer()
		defer server.Close()
		
		// GetAllFeedsのモック設定
		mockFeeds := []model.Feed{
			{ID: 1, UserId: feedArticleTestUser.ID, URL: server.URL},
			{ID: 2, UserId: feedArticleTestUser.ID, URL: server.URL},
		}
		mockFeedRepo.On("GetAllFeeds", mock.AnythingOfType("*[]model.Feed"), feedArticleTestUser.ID).
			Run(func(args mock.Arguments) {
				feeds := args.Get(0).(*[]model.Feed)
				*feeds = mockFeeds
			}).
			Return(mockFeeds, nil)
		
		// 各フィードに対するGetFeedByIdのモック設定
		for _, feed := range mockFeeds {
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, feed.ID).
				Run(func(args mock.Arguments) {
					feedArg := args.Get(0).(*model.Feed)
					*feedArg = feed
				}).
				Return(feed, nil)
		}
		
		// テスト対象の関数を実行
		articles, err := feedArticleRepo.GetAllArticles(feedArticleTestUser.ID)
		
		// 検証
		assert.NoError(t, err)
		assert.Len(t, articles, 4) // 2つのフィードそれぞれに2つの記事 = 合計4つ
		
		// フィードIDが正しく設定されているか確認
		feedIDCount := make(map[uint]int)
		for _, article := range articles {
			feedIDCount[article.FeedID]++
		}
		
		// 各フィードから2つずつ記事が取得されているか確認
		assert.Equal(t, 2, feedIDCount[1])
		assert.Equal(t, 2, feedIDCount[2])
		
		// モックが期待通り呼ばれたことを確認
		mockFeedRepo.AssertExpectations(t)
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("フィード一覧取得エラー", func(t *testing.T) {
			// モックをリセットして新しいモック設定
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = repository.NewFeedArticleRepository(mockFeedRepo)
			
			// GetAllFeedsが必ずエラーを返すように設定
			mockErr := errors.New("フィードの取得に失敗しました")
			mockFeedRepo.On("GetAllFeeds", mock.AnythingOfType("*[]model.Feed"), feedArticleTestUser.ID).
				Run(func(args mock.Arguments) {
					// 空のスライスを設定（nilではなく）
					feeds := args.Get(0).(*[]model.Feed)
					*feeds = []model.Feed{}
				}).
				Return(nil, mockErr)
			
			// テスト対象の関数を実行
			articles, err := feedArticleRepo.GetAllArticles(feedArticleTestUser.ID)
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "フィードの取得に失敗")
			assert.Nil(t, articles)
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
		
		t.Run("一部のフィード取得に失敗しても残りのフィードからは記事を取得する", func(t *testing.T) {
			// 正常なサーバーを作成
			goodServer := createMockRSSServer()
			defer goodServer.Close()
			
			// エラーを返すサーバーを作成
			errorServer := createErrorRSSServer(http.StatusForbidden)
			defer errorServer.Close()
			
			// 新しいモックを作成
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = repository.NewFeedArticleRepository(mockFeedRepo)
			
			// GetAllFeedsのモック設定 - 正常なフィードと問題のあるフィードの混在
			mockFeeds := []model.Feed{
				{ID: 1, UserId: feedArticleTestUser.ID, URL: goodServer.URL},    // 正常なフィード
				{ID: 2, UserId: feedArticleTestUser.ID, URL: errorServer.URL},   // 問題のあるフィード
				{ID: 3, UserId: feedArticleTestUser.ID, URL: goodServer.URL},    // 正常なフィード
			}
			
			mockFeedRepo.On("GetAllFeeds", mock.AnythingOfType("*[]model.Feed"), feedArticleTestUser.ID).
				Run(func(args mock.Arguments) {
					feeds := args.Get(0).(*[]model.Feed)
					*feeds = mockFeeds
				}).
				Return(mockFeeds, nil)
			
			// 各フィードに対するGetFeedByIdのモック設定
			for _, feed := range mockFeeds {
				mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, feed.ID).
					Run(func(args mock.Arguments) {
						feedArg := args.Get(0).(*model.Feed)
						*feedArg = feed
					}).
					Return(feed, nil)
			}
			
			// テスト対象の関数を実行
			articles, err := feedArticleRepo.GetAllArticles(feedArticleTestUser.ID)
			
			// 正常に終了し、一部の記事が取得できることを検証
			assert.NoError(t, err)
			assert.NotEmpty(t, articles)
			
			// フィードIDが正しく設定されているか確認
			feedIDCount := make(map[uint]int)
			for _, article := range articles {
				feedIDCount[article.FeedID]++
			}
			
			// ID 1と3のフィードからは記事が取得できるが、ID 2のフィードからは取得できない
			assert.Equal(t, 2, feedIDCount[1])
			assert.Equal(t, 0, feedIDCount[2]) // エラーのあるフィードからは記事が取得できない
			assert.Equal(t, 2, feedIDCount[3])
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
	})
}
