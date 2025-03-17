package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

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
	userId := getUserIdFromToken(c)
	
	layoutsRes, err := lc.lu.GetAllLayouts(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, layoutsRes)
}

func (lc *layoutController) GetLayoutById(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("layoutId")
	layoutId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なレイアウトIDです"})
	}
	
	layoutRes, err := lc.lu.GetLayoutById(userId, uint(layoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, layoutRes)
}

func (lc *layoutController) CreateLayout(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	var request model.LayoutRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	layoutRes, err := lc.lu.CreateLayout(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, layoutRes)
}

func (lc *layoutController) UpdateLayout(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("layoutId")
	layoutId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なレイアウトIDです"})
	}
	
	var request model.LayoutRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	layoutRes, err := lc.lu.UpdateLayout(request, userId, uint(layoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, layoutRes)
}

func (lc *layoutController) DeleteLayout(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("layoutId")
	layoutId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なレイアウトIDです"})
	}
	
	err = lc.lu.DeleteLayout(userId, uint(layoutId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}