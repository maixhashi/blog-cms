package feed_article_test

import (
	"go-react-app/model"
	"net/http"
	"net/http/httptest"
	
	"github.com/stretchr/testify/mock"
)

// テスト用のHTTPサーバーを作成するヘルパー関数
func createMockRSSServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockRSSXML))
	}))
	return server
}

// エラーを返すHTTPサーバーを作成するヘルパー関数
func createErrorRSSServer(statusCode int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte("Error response"))
	}))
	return server
}

// 無効なXMLを返すHTTPサーバーを作成するヘルパー関数
func createInvalidXMLServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is not valid XML"))
	}))
	return server
}

// モックフィードを作成するヘルパー関数
func createMockFeed(id uint, url string) model.Feed {
	return model.Feed{
		ID:     id,
		UserId: feedArticleTestUser.ID,
		URL:    url,
	}
}

// GetFeedByIdのモックを設定するヘルパー関数
func setupGetFeedByIdMock(feed model.Feed) {
	mockFeedRepo.On("GetFeedById", mock.AnythingOfType("*model.Feed"), feedArticleTestUser.ID, feed.ID).
		Run(func(args mock.Arguments) {
			feedArg := args.Get(0).(*model.Feed)
			*feedArg = feed
		}).
		Return(feed, nil)
}
