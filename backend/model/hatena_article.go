package model

import "time"

type HatenaArticle struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Content     string    `json:"content"`
	Summary     string    `json:"summary"`
	Categories  []string  `json:"categories"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      string    `json:"author"`
}

type HatenaArticleResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Summary     string    `json:"summary"`
	Categories  []string  `json:"categories"`
	PublishedAt time.Time `json:"published_at"`
	Author      string    `json:"author"`
}
