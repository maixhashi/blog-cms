package model

import "time"

type FeedArticle struct {
	ID          string    `json:"id"`
	FeedID      uint      `json:"feed_id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Content     string    `json:"content"`
	Summary     string    `json:"summary"`
	Categories  []string  `json:"categories"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      string    `json:"author"`
}

type FeedArticleResponse struct {
	ID          string    `json:"id"`
	FeedID      uint      `json:"feed_id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Summary     string    `json:"summary"`
	Categories  []string  `json:"categories"`
	PublishedAt time.Time `json:"published_at"`
	Author      string    `json:"author"`
}
