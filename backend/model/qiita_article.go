package model

import "time"

type QiitaArticle struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	Body         string    `json:"body"`
	LikesCount   int       `json:"likes_count"`
	ReactionsCount int     `json:"reactions_count"`
	CommentsCount int      `json:"comments_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Tags         []QiitaTag `json:"tags"`
	User         QiitaUser  `json:"user"`
}

type QiitaTag struct {
	Name string `json:"name"`
}

type QiitaUser struct {
	ID              string `json:"id"`
	ProfileImageURL string `json:"profile_image_url"`
	Name            string `json:"name"`
}

type QiitaArticleResponse struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	LikesCount   int       `json:"likes_count"`
	Tags         []QiitaTag `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	User         QiitaUser  `json:"user"`
}
