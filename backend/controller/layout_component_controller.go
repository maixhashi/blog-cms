package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILayoutComponentController interface {
	GetAllLayoutComponents(c echo.Context) error
	GetLayoutComponentById(c echo.Context) error
	CreateLayoutComponent(c echo.Context) error
	UpdateLayoutComponent(c echo.Context) error
	DeleteLayoutComponent(c echo.Context) error
}

type layoutComponentController struct {
	lcu usecase.ILayoutComponentUsecase
}

func NewLayoutComponentController(lcu usecase.ILayoutComponentUsecase) ILayoutComponentController {
	return &layoutComponentController{lcu}
}

func (lcc *layoutComponentController) GetAllLayoutComponents(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	componentsRes, err := lcc.lcu.GetAllLayoutComponents(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, componentsRes)
}

func (lcc *layoutComponentController) GetLayoutComponentById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("componentId")
	componentId, _ := strconv.Atoi(id)

	componentRes, err := lcc.lcu.GetLayoutComponentById(uint(userId.(float64)), uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, componentRes)
}

func (lcc *layoutComponentController) CreateLayoutComponent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	component := model.LayoutComponent{}
	if err := c.Bind(&component); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	component.UserId = uint(userId.(float64))

	componentRes, err := lcc.lcu.CreateLayoutComponent(component)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, componentRes)
}

func (lcc *layoutComponentController) UpdateLayoutComponent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("componentId")
	componentId, _ := strconv.Atoi(id)

	component := model.LayoutComponent{}
	if err := c.Bind(&component); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	componentRes, err := lcc.lcu.UpdateLayoutComponent(component, uint(userId.(float64)), uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, componentRes)
}

func (lcc *layoutComponentController) DeleteLayoutComponent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("componentId")
	componentId, _ := strconv.Atoi(id)

	err := lcc.lcu.DeleteLayoutComponent(uint(userId.(float64)), uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
