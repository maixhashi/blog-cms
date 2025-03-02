package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go-react-app/usecase"
)

type IFeedArticleController interface {
	GetArticlesByFeedID(c echo.Context) error
	GetArticleByID(c echo.Context) error
}

type feedArticleController struct {
	fau usecase.IFeedArticleUsecase
}

func NewFeedArticleController(fau usecase.IFeedArticleUsecase) IFeedArticleController {
	return &feedArticleController{fau}
}

func (fac *feedArticleController) GetArticlesByFeedID(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	feedIDStr := c.Param("feedId")
	feedID, err := strconv.ParseUint(feedIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "無効なフィードIDです",
		})
	}

	// ユーザーIDを引数に追加
	articles, err := fac.fau.GetArticlesByFeedID(userId, uint(feedID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, articles)
}

func (fac *feedArticleController) GetArticleByID(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	feedIDStr := c.Param("feedId")
	feedID, err := strconv.ParseUint(feedIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "無効なフィードIDです",
		})
	}

	articleID := c.Param("articleId")
	// ユーザーIDを引数に追加
	article, err := fac.fau.GetArticleByID(userId, uint(feedID), articleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, article)
}
