package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type IArticleUsecase interface {
	GetAllArticles(userId uint) ([]model.ArticleResponse, error)
	GetArticleById(userId uint, articleId uint) (model.ArticleResponse, error)
	CreateArticle(article model.Article) (model.ArticleResponse, error)
	UpdateArticle(article model.Article, userId uint, articleId uint) (model.ArticleResponse, error)
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
	resArticles := []model.ArticleResponse{}
	for _, v := range articles {
		a := model.ArticleResponse{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			Published: v.Published,
			Tags:      v.Tags,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resArticles = append(resArticles, a)
	}
	return resArticles, nil
}

func (au *articleUsecase) GetArticleById(userId uint, articleId uint) (model.ArticleResponse, error) {
	article := model.Article{}
	if err := au.ar.GetArticleById(&article, userId, articleId); err != nil {
		return model.ArticleResponse{}, err
	}
	resArticle := model.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Published: article.Published,
		Tags:      article.Tags,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	return resArticle, nil
}

func (au *articleUsecase) CreateArticle(article model.Article) (model.ArticleResponse, error) {
	if err := au.av.ArticleValidate(article); err != nil {
		return model.ArticleResponse{}, err
	}
	if err := au.ar.CreateArticle(&article); err != nil {
		return model.ArticleResponse{}, err
	}
	resArticle := model.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Published: article.Published,
		Tags:      article.Tags,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	return resArticle, nil
}

func (au *articleUsecase) UpdateArticle(article model.Article, userId uint, articleId uint) (model.ArticleResponse, error) {
	if err := au.av.ArticleValidate(article); err != nil {
		return model.ArticleResponse{}, err
	}
	if err := au.ar.UpdateArticle(&article, userId, articleId); err != nil {
		return model.ArticleResponse{}, err
	}
	resArticle := model.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Published: article.Published,
		Tags:      article.Tags,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	return resArticle, nil
}

func (au *articleUsecase) DeleteArticle(userId uint, articleId uint) error {
	if err := au.ar.DeleteArticle(userId, articleId); err != nil {
		return err
	}
	return nil
}
