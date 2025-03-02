package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-react-app/usecase"
)

type IHatenaController interface {
	GetHatenaArticles(c echo.Context) error
	GetHatenaArticleByID(c echo.Context) error
}

type hatenaController struct {
	hu usecase.IHatenaUsecase
}

func NewHatenaController(hu usecase.IHatenaUsecase) IHatenaController {
	return &hatenaController{hu}
}

func (hc *hatenaController) GetHatenaArticles(c echo.Context) error {
	articles, err := hc.hu.GetHatenaArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, articles)
}

func (hc *hatenaController) GetHatenaArticleByID(c echo.Context) error {
	id := c.Param("id")
	article, err := hc.hu.GetHatenaArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, article)
}
