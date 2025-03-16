package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IArticleController interface {
	GetAllArticles(c echo.Context) error
	GetArticleById(c echo.Context) error
	CreateArticle(c echo.Context) error
	UpdateArticle(c echo.Context) error
	DeleteArticle(c echo.Context) error
}

type articleController struct {
	au usecase.IArticleUsecase
}

func NewArticleController(au usecase.IArticleUsecase) IArticleController {
	return &articleController{au}
}

func (ac *articleController) GetAllArticles(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	articlesRes, err := ac.au.GetAllArticles(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articlesRes)
}

func (ac *articleController) GetArticleById(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("articleId")
	articleId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効な記事IDです"})
	}
	
	articleRes, err := ac.au.GetArticleById(userId, uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articleRes)
}

func (ac *articleController) CreateArticle(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	var request model.ArticleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	articleRes, err := ac.au.CreateArticle(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, articleRes)
}

func (ac *articleController) UpdateArticle(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("articleId")
	articleId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効な記事IDです"})
	}
	
	var request model.ArticleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	articleRes, err := ac.au.UpdateArticle(request, userId, uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articleRes)
}

func (ac *articleController) DeleteArticle(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("articleId")
	articleId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効な記事IDです"})
	}
	
	err = ac.au.DeleteArticle(userId, uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
