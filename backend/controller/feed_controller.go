package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IFeedController interface {
	GetAllFeeds(c echo.Context) error
	GetFeedById(c echo.Context) error
	CreateFeed(c echo.Context) error
	UpdateFeed(c echo.Context) error
	DeleteFeed(c echo.Context) error
}

type feedController struct {
	fu usecase.IFeedUsecase
}

func NewFeedController(fu usecase.IFeedUsecase) IFeedController {
	return &feedController{fu}
}

func (fc *feedController) GetAllFeeds(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	feedsRes, err := fc.fu.GetAllFeeds(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, feedsRes)
}

func (fc *feedController) GetFeedById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("feedId")
	feedId, _ := strconv.Atoi(id)

	feedRes, err := fc.fu.GetFeedById(uint(userId.(float64)), uint(feedId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, feedRes)
}

func (fc *feedController) CreateFeed(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	feed := model.Feed{}
	if err := c.Bind(&feed); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	feed.UserId = uint(userId.(float64))

	feedRes, err := fc.fu.CreateFeed(feed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, feedRes)
}

func (fc *feedController) UpdateFeed(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("feedId")
	feedId, _ := strconv.Atoi(id)

	feed := model.Feed{}
	if err := c.Bind(&feed); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	feedRes, err := fc.fu.UpdateFeed(feed, uint(userId.(float64)), uint(feedId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, feedRes)
}

func (fc *feedController) DeleteFeed(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("feedId")
	feedId, _ := strconv.Atoi(id)

	err := fc.fu.DeleteFeed(uint(userId.(float64)), uint(feedId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
