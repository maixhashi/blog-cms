package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IExternalAPIController interface {
	GetAllExternalAPIs(c echo.Context) error
	GetExternalAPIById(c echo.Context) error
	CreateExternalAPI(c echo.Context) error
	UpdateExternalAPI(c echo.Context) error
	DeleteExternalAPI(c echo.Context) error
}

type externalAPIController struct {
	au usecase.IExternalAPIUsecase
}

func NewExternalAPIController(au usecase.IExternalAPIUsecase) IExternalAPIController {
	return &externalAPIController{au}
}

func (ac *externalAPIController) GetAllExternalAPIs(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	apisRes, err := ac.au.GetAllExternalAPIs(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, apisRes)
}

func (ac *externalAPIController) GetExternalAPIById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("apiId")
	apiId, _ := strconv.Atoi(id)
	apiRes, err := ac.au.GetExternalAPIById(uint(userId.(float64)), uint(apiId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, apiRes)
}

func (ac *externalAPIController) CreateExternalAPI(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	api := model.ExternalAPI{}
	if err := c.Bind(&api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	api.UserId = uint(userId.(float64))
	apiRes, err := ac.au.CreateExternalAPI(api)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, apiRes)
}

func (ac *externalAPIController) UpdateExternalAPI(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("apiId")
	apiId, _ := strconv.Atoi(id)

	api := model.ExternalAPI{}
	if err := c.Bind(&api); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	apiRes, err := ac.au.UpdateExternalAPI(api, uint(userId.(float64)), uint(apiId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, apiRes)
}

func (ac *externalAPIController) DeleteExternalAPI(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("apiId")
	apiId, _ := strconv.Atoi(id)

	err := ac.au.DeleteExternalAPI(uint(userId.(float64)), uint(apiId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
