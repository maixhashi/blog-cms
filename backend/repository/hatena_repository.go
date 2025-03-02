package repository

import (
	"encoding/xml"
	"fmt"
	"go-react-app/model"
	"io"
	"net/http"
	"time"
)

type IHatenaRepository interface {
	GetHatenaArticles() ([]model.HatenaArticle, error)
	GetHatenaArticleByID(id string) (model.HatenaArticle, error)
}

type hatenaRepository struct {
	feedURL string
}

func NewHatenaRepository(feedURL string) IHatenaRepository {
	return &hatenaRepository{feedURL: feedURL}
}

// XMLフィードの構造体
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

type Entry struct {
	ID        string    `xml:"id"`
	Title     string    `xml:"title"`
	Links     []struct {
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr,omitempty"`
	} `xml:"link"`
	Summary   struct {
		Type    string `xml:"type,attr"`
		Content string `xml:",chardata"`
	} `xml:"summary"`
	Published time.Time `xml:"published"`
	Updated   time.Time `xml:"updated"`
	Author    struct {
		Name string `xml:"name"`
	} `xml:"author"`
	Content   struct {
		Type    string `xml:"type,attr"`
		Content string `xml:",chardata"`
	} `xml:"content"`
	Category  []struct {
		Term string `xml:"term,attr"`
	} `xml:"category"`
}

func (hr *hatenaRepository) GetHatenaArticles() ([]model.HatenaArticle, error) {
	resp, err := http.Get(hr.feedURL)
	if err != nil {
		return nil, fmt.Errorf("はてなフィードの取得に失敗しました: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスボディの読み込みに失敗しました: %w", err)
	}

	var feed Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("XMLのパースに失敗しました: %w", err)
	}

	var articles []model.HatenaArticle
	for _, entry := range feed.Entries {
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

		article := model.HatenaArticle{
			ID:          entry.ID,
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

func (hr *hatenaRepository) GetHatenaArticleByID(id string) (model.HatenaArticle, error) {
	articles, err := hr.GetHatenaArticles()
	if err != nil {
		return model.HatenaArticle{}, err
	}

	for _, article := range articles {
		if article.ID == id {
			return article, nil
		}
	}

	return model.HatenaArticle{}, fmt.Errorf("記事が見つかりません: %s", id)
}