package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
)

type IFeedArticleUsecase interface {
	GetArticlesByFeedID(userId uint, feedID uint) ([]model.FeedArticleResponse, error)
	GetArticleByID(userId uint, feedID uint, articleID string) (model.FeedArticleResponse, error)
}

type feedArticleUsecase struct {
	far repository.IFeedArticleRepository
}

func NewFeedArticleUsecase(far repository.IFeedArticleRepository) IFeedArticleUsecase {
	return &feedArticleUsecase{far}
}

func (fau *feedArticleUsecase) GetArticlesByFeedID(userId uint, feedID uint) ([]model.FeedArticleResponse, error) {
	// リポジトリ層にユーザーIDを渡す
	articles, err := fau.far.GetArticlesByFeedID(userId, feedID)
	if err != nil {
		return nil, err
	}

	var response []model.FeedArticleResponse
	for _, article := range articles {
		response = append(response, model.FeedArticleResponse{
			ID:          article.ID,
			FeedID:      article.FeedID,
			Title:       article.Title,
			URL:         article.URL,
			Summary:     article.Summary,
			Categories:  article.Categories,
			PublishedAt: article.PublishedAt,
			Author:      article.Author,
		})
	}

	return response, nil
}

func (fau *feedArticleUsecase) GetArticleByID(userId uint, feedID uint, articleID string) (model.FeedArticleResponse, error) {
	// リポジトリ層にユーザーIDを渡す
	article, err := fau.far.GetArticleByID(userId, feedID, articleID)
	if err != nil {
		return model.FeedArticleResponse{}, err
	}

	response := model.FeedArticleResponse{
		ID:          article.ID,
		FeedID:      article.FeedID,
		Title:       article.Title,
		URL:         article.URL,
		Summary:     article.Summary,
		Categories:  article.Categories,
		PublishedAt: article.PublishedAt,
		Author:      article.Author,
	}

	return response, nil
}
