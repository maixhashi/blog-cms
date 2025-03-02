package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
)

type IHatenaUsecase interface {
	GetHatenaArticles() ([]model.HatenaArticleResponse, error)
	GetHatenaArticleByID(id string) (model.HatenaArticleResponse, error)
}

type hatenaUsecase struct {
	hr repository.IHatenaRepository
}

func NewHatenaUsecase(hr repository.IHatenaRepository) IHatenaUsecase {
	return &hatenaUsecase{hr}
}

func (hu *hatenaUsecase) GetHatenaArticles() ([]model.HatenaArticleResponse, error) {
	articles, err := hu.hr.GetHatenaArticles()
	if err != nil {
		return nil, err
	}

	var response []model.HatenaArticleResponse
	for _, article := range articles {
		response = append(response, model.HatenaArticleResponse{
			ID:          article.ID,
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

func (hu *hatenaUsecase) GetHatenaArticleByID(id string) (model.HatenaArticleResponse, error) {
	article, err := hu.hr.GetHatenaArticleByID(id)
	if err != nil {
		return model.HatenaArticleResponse{}, err
	}

	response := model.HatenaArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		URL:         article.URL,
		Summary:     article.Summary,
		Categories:  article.Categories,
		PublishedAt: article.PublishedAt,
		Author:      article.Author,
	}

	return response, nil
}
