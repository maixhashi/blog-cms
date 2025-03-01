package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
)

type IQiitaUsecase interface {
	GetQiitaArticles() ([]model.QiitaArticleResponse, error)
	GetQiitaArticleByID(id string) (model.QiitaArticleResponse, error)
}

type qiitaUsecase struct {
	qr repository.IQiitaRepository
}

func NewQiitaUsecase(qr repository.IQiitaRepository) IQiitaUsecase {
	return &qiitaUsecase{qr}
}

func (qu *qiitaUsecase) GetQiitaArticles() ([]model.QiitaArticleResponse, error) {
	articles, err := qu.qr.GetQiitaArticles()
	if err != nil {
		return nil, err
	}

	var response []model.QiitaArticleResponse
	for _, article := range articles {
		response = append(response, model.QiitaArticleResponse{
			ID:           article.ID,
			Title:        article.Title,
			URL:          article.URL,
			LikesCount:   article.LikesCount,
			Tags:         article.Tags,
			CreatedAt:    article.CreatedAt,
			User:         article.User,
		})
	}

	return response, nil
}

func (qu *qiitaUsecase) GetQiitaArticleByID(id string) (model.QiitaArticleResponse, error) {
	article, err := qu.qr.GetQiitaArticleByID(id)
	if err != nil {
		return model.QiitaArticleResponse{}, err
	}

	response := model.QiitaArticleResponse{
		ID:           article.ID,
		Title:        article.Title,
		URL:          article.URL,
		LikesCount:   article.LikesCount,
		Tags:         article.Tags,
		CreatedAt:    article.CreatedAt,
		User:         article.User,
	}

	return response, nil
}
