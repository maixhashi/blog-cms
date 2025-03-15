package feed_article_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	feedArticleDB      *gorm.DB
	mockFeedRepo       *MockFeedRepository
	feedArticleRepo    repository.IFeedArticleRepository
	feedArticleTestUser model.User
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
	feedArticleTestUser = testutils.CreateTestUser(feedArticleDB)
	
	// モックフィードリポジトリの作成
	mockFeedRepo = new(MockFeedRepository)
	
	// フィード記事リポジトリのインスタンス化
	feedArticleRepo = repository.NewFeedArticleRepository(mockFeedRepo)
}
