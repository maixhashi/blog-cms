package repository

import (
	"errors"
	"go-react-app/model"
	"go-react-app/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	feedArticleDB   *gorm.DB
	mockFeedRepo    *MockFeedRepository
	feedArticleRepo IFeedArticleRepository
	articleTestUser model.User
)

// フィードリポジトリのモック
type MockFeedRepository struct {
	mock.Mock
}

func (m *MockFeedRepository) GetAllFeeds(feeds *[]model.Feed, userId uint) error {
	args := m.Called(feeds, userId)
	// モックが呼ばれたときに引数として渡されたfeedsスライスにデータを設定
	if feeds != nil && args.Get(0) != nil {
		mockFeeds := args.Get(0).([]model.Feed)
		*feeds = mockFeeds
	}
	return args.Error(1)
}

func (m *MockFeedRepository) GetFeedById(feed *model.Feed, userId uint, feedId uint) error {
	args := m.Called(feed, userId, feedId)
	// モックが呼ばれたときに引数として渡されたfeedにデータを設定
	if feed != nil && args.Get(0) != nil {
		mockFeed := args.Get(0).(model.Feed)
		*feed = mockFeed
	}
	return args.Error(1)
}

func (m *MockFeedRepository) CreateFeed(feed *model.Feed) error {
	args := m.Called(feed)
	return args.Error(0)
}

func (m *MockFeedRepository) UpdateFeed(feed *model.Feed, userId uint, feedId uint) error {
	args := m.Called(feed, userId, feedId)
	return args.Error(0)
}

func (m *MockFeedRepository) DeleteFeed(userId uint, feedId uint) error {
	args := m.Called(userId, feedId)
	return args.Error(0)
}

// RSSフィードのモックレスポンス
const mockRSSXML = `<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>Test Feed</title>
  <entry>
    <id>article1</id>
    <title>Test Article 1</title>
    <link href="https://example.com/article1" rel="alternate"/>
    <summary type="html"><![CDATA[Summary of article 1]]></summary>
    <category term="Tech"/>
    <published>2023-01-01T00:00:00Z</published>
    <updated>2023-01-02T00:00:00Z</updated>
    <author>
      <name>Test Author</name>
    </author>
    <content type="html"><![CDATA[Content of article 1]]></content>
  </entry>
  <entry>
    <id>article2</id>
    <title>Test Article 2</title>
    <link href="https://example.com/article2" rel="alternate"/>
    <summary type="html"><![CDATA[Summary of article 2]]></summary>
    <category term="News"/>
    <published>2023-01-03T00:00:00Z</published>
    <updated>2023-01-04T00:00:00Z</updated>
    <author>
      <name>Another Author</name>
    </author>
    <content type="html"><![CDATA[Content of article 2]]></content>
  </entry>
</feed>`

func setupFeedArticleTest() {
	// テストDBのセットアップ
	feedArticleDB = testutils.SetupTestDB()
	
	// テストユーザーの作成
	articleTestUser = testutils.CreateTestUser(feedArticleDB)
	
	// モックフィードリポジトリの作成
	mockFeedRepo = new(MockFeedRepository)
	
	// フィード記事リポジトリのインスタンス化
	feedArticleRepo = NewFeedArticleRepository(mockFeedRepo)
}
func TestFeedArticleRepository_GetArticlesByFeedID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		// HTTPサーバーのモック作成
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockRSSXML))
		}))
		defer server.Close()
		
		// GetFeedByIdのモック設定
		mockFeed := model.Feed{
			ID:     1,
			UserId: articleTestUser.ID,	
			URL:    server.URL, // モックサーバーのURL
		}
		mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, uint(1)).
			Run(func(args mock.Arguments) {
				feed := args.Get(0).(*model.Feed)
				*feed = mockFeed
			}).
			Return(mockFeed, nil)
		
		// テスト対象の関数を実行
		articles, err := feedArticleRepo.GetArticlesByFeedID(articleTestUser.ID, 1)
		
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
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, uint(999)).
				Return(model.Feed{}, assert.AnError)
			
			// テスト対象の関数を実行
			_, err := feedArticleRepo.GetArticlesByFeedID(articleTestUser.ID, 999)
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "フィードの取得に失敗")
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
	})
}

func TestFeedArticleRepository_GetArticleByID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		// HTTPサーバーのモック作成
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockRSSXML))
		}))
		defer server.Close()
		
		// GetFeedByIdのモック設定
		mockFeed := model.Feed{
			ID:     1,
			UserId: articleTestUser.ID,
			URL:    server.URL, // モックサーバーのURL
		}
		mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, uint(1)).
			Run(func(args mock.Arguments) {
				feed := args.Get(0).(*model.Feed)
				*feed = mockFeed
			}).
			Return(mockFeed, nil)
		
		// テスト対象の関数を実行
		article, err := feedArticleRepo.GetArticleByID(articleTestUser.ID, 1, "article1")
		
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
			// 新しいテストサーバーを作成し、そのスコープ内で保持する
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/xml")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockRSSXML))
			}))
			// テストケース終了時にサーバーを確実に閉じる
			defer server.Close()
			
			// モックをリセットしてから、新しいモック設定を行う
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = NewFeedArticleRepository(mockFeedRepo)
			
			// GetFeedByIdのモック設定
			mockFeed := model.Feed{
				ID:     1,
				UserId: articleTestUser.ID,
				URL:    server.URL, // 新しく作ったサーバーのURLを使用
			}
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, uint(1)).
				Run(func(args mock.Arguments) {
					feed := args.Get(0).(*model.Feed)
					*feed = mockFeed
				}).
				Return(mockFeed, nil)
			
			// 存在しない記事IDでテスト
			_, err := feedArticleRepo.GetArticleByID(articleTestUser.ID, 1, "nonexistent")
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "記事が見つかりません")
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
	})
}

func TestFeedArticleRepository_GetAllArticles(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		// HTTPサーバーのモック作成
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockRSSXML))
		}))
		defer server.Close()
		
		// GetAllFeedsのモック設定
		mockFeeds := []model.Feed{
			{ID: 1, UserId: articleTestUser.ID, URL: server.URL},
			{ID: 2, UserId: articleTestUser.ID, URL: server.URL},
		}
		mockFeedRepo.On("GetAllFeeds", mock.AnythingOfType("*[]model.Feed"), articleTestUser.ID).
			Run(func(args mock.Arguments) {
				feeds := args.Get(0).(*[]model.Feed)
				*feeds = mockFeeds
			}).
			Return(mockFeeds, nil)
		
		// 各フィードに対するGetFeedByIdのモック設定
		for _, feed := range mockFeeds {
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, feed.ID).
				Run(func(args mock.Arguments) {
					feedArg := args.Get(0).(*model.Feed)
					*feedArg = feed
				}).
				Return(feed, nil)
		}
		
		// テスト対象の関数を実行
		articles, err := feedArticleRepo.GetAllArticles(articleTestUser.ID)
		
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
			feedArticleRepo = NewFeedArticleRepository(mockFeedRepo)
			
			// GetAllFeedsが必ずエラーを返すように設定
			mockErr := errors.New("フィードの取得に失敗しました")
			mockFeedRepo.On("GetAllFeeds", mock.AnythingOfType("*[]model.Feed"), articleTestUser.ID).
				Run(func(args mock.Arguments) {
					// 空のスライスを設定（nilではなく）
					feeds := args.Get(0).(*[]model.Feed)
					*feeds = []model.Feed{}
				}).
				Return(nil, mockErr)
			
			// テスト対象の関数を実行
			articles, err := feedArticleRepo.GetAllArticles(articleTestUser.ID)
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "フィードの取得に失敗")
			assert.Nil(t, articles)
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
		
		t.Run("一部のフィード取得に失敗しても残りのフィードからは記事を取得する", func(t *testing.T) {
			// 正常なサーバーを作成
			goodServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/xml")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockRSSXML))
			}))
			defer goodServer.Close()
			
			// エラーを返すサーバーを作成
			errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Access Denied"))
			}))
			defer errorServer.Close()
			
			// 新しいモックを作成
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = NewFeedArticleRepository(mockFeedRepo)
			
			// GetAllFeedsのモック設定 - 正常なフィードと問題のあるフィードの混在
			mockFeeds := []model.Feed{
				{ID: 1, UserId: articleTestUser.ID, URL: goodServer.URL},    // 正常なフィード
				{ID: 2, UserId: articleTestUser.ID, URL: errorServer.URL},   // 問題のあるフィード
				{ID: 3, UserId: articleTestUser.ID, URL: goodServer.URL},    // 正常なフィード
			}
			
			mockFeedRepo.On("GetAllFeeds", mock.AnythingOfType("*[]model.Feed"), articleTestUser.ID).
				Run(func(args mock.Arguments) {
					feeds := args.Get(0).(*[]model.Feed)
					*feeds = mockFeeds
				}).
				Return(mockFeeds, nil)
			
			// 各フィードに対するGetFeedByIdのモック設定
			for _, feed := range mockFeeds {
				mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, feed.ID).
					Run(func(args mock.Arguments) {
						feedArg := args.Get(0).(*model.Feed)
						*feedArg = feed
					}).
					Return(feed, nil)
			}
			
			// テスト対象の関数を実行
			articles, err := feedArticleRepo.GetAllArticles(articleTestUser.ID)
			
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
		
		t.Run("無効なXMLフォーマットのフィード", func(t *testing.T) {
			// 無効なXMLを返すサーバー
			invalidXMLServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/xml")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("This is not valid XML"))
			}))
			defer invalidXMLServer.Close()
			
			// 新しいモックを作成
			mockFeedRepo = new(MockFeedRepository)
			feedArticleRepo = NewFeedArticleRepository(mockFeedRepo)
			
			// GetFeedByIdのモック設定
			mockFeed := model.Feed{
				ID:     1,
				UserId: articleTestUser.ID,
				URL:    invalidXMLServer.URL,
			}
			mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), articleTestUser.ID, uint(1)).
				Run(func(args mock.Arguments) {
					feed := args.Get(0).(*model.Feed)
					*feed = mockFeed
				}).
				Return(mockFeed, nil)
			
			// テスト対象の関数を実行
			_, err := feedArticleRepo.GetArticlesByFeedID(articleTestUser.ID, 1)
			
			// エラーが返されることを検証
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "XMLのパース")
			
			// モックが期待通り呼ばれたことを確認
			mockFeedRepo.AssertExpectations(t)
		})
	})
}
