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

// GetAllArticles ユーザーのすべての記事を取得
// @Summary ユーザーの記事一覧を取得
// @Description ログインユーザーのすべての記事を取得する
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {array} model.ArticleResponse
// @Failure 500 {object} map[string]string
// @Router /articles [get]
func (ac *articleController) GetAllArticles(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	articlesRes, err := ac.au.GetAllArticles(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articlesRes)
}

// GetArticleById 指定されたIDの記事を取得
// @Summary 特定の記事を取得
// @Description 指定されたIDの記事を取得する
// @Tags articles
// @Accept json
// @Produce json
// @Param articleId path int true "記事ID"
// @Success 200 {object} model.ArticleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{articleId} [get]
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

// CreateArticle 新しい記事を作成
// @Summary 新しい記事を作成
// @Description ユーザーの新しい記事を作成する
// @Tags articles
// @Accept json
// @Produce json
// @Param article body model.ArticleRequest true "記事情報"
// @Success 201 {object} model.ArticleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles [post]
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

// UpdateArticle 既存の記事を更新
// @Summary 記事を更新
// @Description 指定されたIDの記事を更新する
// @Tags articles
// @Accept json
// @Produce json
// @Param articleId path int true "記事ID"
// @Param article body model.ArticleRequest true "更新する記事情報"
// @Success 200 {object} model.ArticleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{articleId} [put]
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

// DeleteArticle 記事を削除
// @Summary 記事を削除
// @Description 指定されたIDの記事を削除する
// @Tags articles
// @Accept json
// @Produce json
// @Param articleId path int true "記事ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{articleId} [delete]
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