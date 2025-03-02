package repository

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-react-app/model"
)

type IHatenaRepository interface {
	GetHatenaArticles() ([]model.HatenaArticle, error)
	GetHatenaArticleByID(id string) (model.HatenaArticle, error)
}

type hatenaRepository struct {
	feedURL string
}

// RSS/Atom用の構造体
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

type Entry struct {
	ID        string    `xml:"id"`
	Title     string    `xml:"title"`
	Link      Link      `xml:"link"`
	Published time.Time `xml:"published"`
	Updated   time.Time `xml:"updated"`
	Summary   string    `xml:"summary"`
	Content   string    `xml:"content"`
	Author    Author    `xml:"author"`
	Categories []Category `xml:"category"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Author struct {
	Name string `xml:"name"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

func NewHatenaRepository() IHatenaRepository {
	return &hatenaRepository{
		feedURL: "https://tech.smarthr.jp/feed", // SmartHRのフィードURL
	}
}

func (hr *hatenaRepository) GetHatenaArticles() ([]model.HatenaArticle, error) {
	resp, err := http.Get(hr.feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("feed request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	var articles []model.HatenaArticle
	for _, entry := range feed.Entries {
		categories := make([]string, len(entry.Categories))
		for i, category := range entry.Categories {
			categories[i] = category.Term
		}

		article := model.HatenaArticle{
			ID:          entry.ID,
			Title:       entry.Title,
			URL:         entry.Link.Href,
			Content:     entry.Content,
			Summary:     entry.Summary,
			Categories:  categories,
			PublishedAt: entry.Published,
			UpdatedAt:   entry.Updated,
			Author:      entry.Author.Name,
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

	return model.HatenaArticle{}, fmt.Errorf("article with ID %s not found", id)
}
