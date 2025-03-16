package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type IArticleUsecase interface {
	GetAllArticles(userId uint) ([]model.ArticleResponse, error)
	GetArticleById(userId uint, articleId uint) (model.ArticleResponse, error)
	CreateArticle(request model.ArticleRequest) (model.ArticleResponse, error)
	UpdateArticle(request model.ArticleRequest, userId uint, articleId uint) (model.ArticleResponse, error)
	DeleteArticle(userId uint, articleId uint) error
}

type articleUsecase struct {
	ar repository.IArticleRepository
	av validator.IArticleValidator
}

func NewArticleUsecase(ar repository.IArticleRepository, av validator.IArticleValidator) IArticleUsecase {
	return &articleUsecase{ar, av}
}

func (au *articleUsecase) GetAllArticles(userId uint) ([]model.ArticleResponse, error) {
	articles := []model.Article{}
	if err := au.ar.GetAllArticles(&articles, userId); err != nil {
		return nil, err
	}
	
	resArticles := make([]model.ArticleResponse, len(articles))
	for i, article := range articles {
		resArticles[i] = article.ToResponse()
	}
	return resArticles, nil
}

func (au *articleUsecase) GetArticleById(userId uint, articleId uint) (model.ArticleResponse, error) {
	article := model.Article{}
	if err := au.ar.GetArticleById(&article, userId, articleId); err != nil {
		return model.ArticleResponse{}, err
	}
	return article.ToResponse(), nil
}

func (au *articleUsecase) CreateArticle(request model.ArticleRequest) (model.ArticleResponse, error) {
	if err := au.av.ValidateArticleRequest(request); err != nil {
		return model.ArticleResponse{}, err
	}
	
	article := request.ToModel()
	if err := au.ar.CreateArticle(&article); err != nil {
		return model.ArticleResponse{}, err
	}
	
	return article.ToResponse(), nil
}

func (au *articleUsecase) UpdateArticle(request model.ArticleRequest, userId uint, articleId uint) (model.ArticleResponse, error) {
	if err := au.av.ValidateArticleRequest(request); err != nil {
		return model.ArticleResponse{}, err
	}
	
	article := request.ToModel()
	if err := au.ar.UpdateArticle(&article, userId, articleId); err != nil {
		return model.ArticleResponse{}, err
	}
	
	return article.ToResponse(), nil
}

func (au *articleUsecase) DeleteArticle(userId uint, articleId uint) error {
	return au.ar.DeleteArticle(userId, articleId)
}