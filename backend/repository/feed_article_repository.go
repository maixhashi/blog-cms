package repository

import (
	"encoding/xml"
	"fmt"
	"go-react-app/model"
	"io"
	"net/http"
)

type IFeedArticleRepository interface {
	GetArticlesByFeedID(userId uint, feedID uint) ([]model.FeedArticle, error)
	GetArticleByID(userId uint, feedID uint, articleID string) (model.FeedArticle, error)
}

type feedArticleRepository struct {
	feedRepository IFeedRepository
}

func NewFeedArticleRepository(fr IFeedRepository) IFeedArticleRepository {
	return &feedArticleRepository{feedRepository: fr}
}

func (far *feedArticleRepository) GetArticlesByFeedID(userId uint, feedID uint) ([]model.FeedArticle, error) {
	// ユーザーIDを引数から受け取る
	var feed model.Feed
	err := far.feedRepository.GetFeedById(&feed, userId, feedID)
	if err != nil {
		return nil, fmt.Errorf("フィードの取得に失敗しました: %w", err)
	}

	// フィードのURLからRSSを取得
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, fmt.Errorf("フィードの取得に失敗しました: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスボディの読み込みに失敗しました: %w", err)
	}

	var xmlFeed Feed
	if err := xml.Unmarshal(body, &xmlFeed); err != nil {
		return nil, fmt.Errorf("XMLのパースに失敗しました: %w", err)
	}

	var articles []model.FeedArticle
	for _, entry := range xmlFeed.Entries {
		// カテゴリの配列を作成
		categories := make([]string, len(entry.Category))
		for i, category := range entry.Category {
			categories[i] = category.Term
		}

		// 記事URLを取得
		var articleURL string
		for _, link := range entry.Links {
			// rel属性がない、またはrel="alternate"のリンクを記事URLとして選択
			if link.Rel == "" || link.Rel == "alternate" {
				articleURL = link.Href
				break
			}
		}

		article := model.FeedArticle{
			ID:          entry.ID,
			FeedID:      feedID,
			Title:       entry.Title,
			URL:         articleURL,
			Summary:     entry.Summary.Content,
			Categories:  categories,
			PublishedAt: entry.Published,
			UpdatedAt:   entry.Updated,
			Author:      entry.Author.Name,
			Content:     entry.Content.Content,
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (far *feedArticleRepository) GetArticleByID(userId uint, feedID uint, articleID string) (model.FeedArticle, error) {
	articles, err := far.GetArticlesByFeedID(userId, feedID)
	if err != nil {
		return model.FeedArticle{}, err
	}

	for _, article := range articles {
		if article.ID == articleID {
			return article, nil
		}
	}

	return model.FeedArticle{}, fmt.Errorf("記事が見つかりません: %s", articleID)
}
