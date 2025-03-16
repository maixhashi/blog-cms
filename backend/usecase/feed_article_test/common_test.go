package feed_article_test

import (
	"go-react-app/model"
	"time"
)

// テスト用変数
var (
	testUserId uint = 1
)

// テスト用記事データ
var testArticles = []model.FeedArticle{
	{
		ID:          "article1",
		FeedID:      1,
		Title:       "Test Article 1",
		URL:         "https://example.com/article1",
		Summary:     "Summary of article 1",
		Categories:  []string{"tech", "news"},
		PublishedAt: time.Now(),
		Author:      "Test Author",
	},
	{
		ID:          "article2",
		FeedID:      1,
		Title:       "Test Article 2",
		URL:         "https://example.com/article2",
		Summary:     "Summary of article 2",
		Categories:  []string{"science"},
		PublishedAt: time.Now(),
		Author:      "Another Author",
	},
	{
		ID:          "article3",
		FeedID:      2,
		Title:       "Test Article 3",
		URL:         "https://example.com/article3",
		Summary:     "Summary of article 3",
		Categories:  []string{"politics"},
		PublishedAt: time.Now(),
		Author:      "Third Author",
	},
}
