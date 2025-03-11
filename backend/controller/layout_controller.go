package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILayoutController interface {
	GetAllLayouts(c echo.Context) error
	GetLayoutById(c echo.Context) error
	CreateLayout(c echo.Context) error
	UpdateLayout(c echo.Context) error
	DeleteLayout(c echo.Context) error
}

type layoutController struct {
	lu usecase.ILayoutUsecase
}

func NewLayoutController(lu usecase.ILayoutUsecase) ILayoutController {
	return &layoutController{lu}
}

func (lc *layoutController) GetAllLayouts(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	layoutsRes, err := lc.lu.GetAllLayouts(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, layoutsRes)
}

func (lc *layoutController) GetLayoutById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("layoutId")
	layoutId, _ := strconv.Atoi(id)

	layoutRes, err := lc.lu.GetLayoutById(uint(userId.(float64)), uint(layoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, layoutRes)
}

func (lc *layoutController) CreateLayout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	layout := model.Layout{}
	if err := c.Bind(&layout); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	layout.UserId = uint(userId.(float64))

	layoutRes, err := lc.lu.CreateLayout(layout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, layoutRes)
}

func (lc *layoutController) UpdateLayout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("layoutId")
	layoutId, _ := strconv.Atoi(id)

	layout := model.Layout{}
	if err := c.Bind(&layout); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	layoutRes, err := lc.lu.UpdateLayout(layout, uint(userId.(float64)), uint(layoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, layoutRes)
}

func (lc *layoutController) DeleteLayout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("layoutId")
	layoutId, _ := strconv.Atoi(id)

	err := lc.lu.DeleteLayout(uint(userId.(float64)), uint(layoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
