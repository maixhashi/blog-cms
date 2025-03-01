package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-react-app/usecase"
)

type IQiitaController interface {
	GetQiitaArticles(c echo.Context) error
	GetQiitaArticleByID(c echo.Context) error
}

type qiitaController struct {
	qu usecase.IQiitaUsecase
}

func NewQiitaController(qu usecase.IQiitaUsecase) IQiitaController {
	return &qiitaController{qu}
}

func (qc *qiitaController) GetQiitaArticles(c echo.Context) error {
	articles, err := qc.qu.GetQiitaArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, articles)
}

func (qc *qiitaController) GetQiitaArticleByID(c echo.Context) error {
	id := c.Param("id")
	article, err := qc.qu.GetQiitaArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, article)
}
