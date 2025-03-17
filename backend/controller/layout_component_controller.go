package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ILayoutComponentController interface {
	GetAllLayoutComponents(c echo.Context) error
	GetLayoutComponentById(c echo.Context) error
	CreateLayoutComponent(c echo.Context) error
	UpdateLayoutComponent(c echo.Context) error
	DeleteLayoutComponent(c echo.Context) error
	
	// 新しいメソッド
	AssignToLayout(c echo.Context) error
	RemoveFromLayout(c echo.Context) error
	UpdatePosition(c echo.Context) error
}

type layoutComponentController struct {
	lcu usecase.ILayoutComponentUsecase
}

func NewLayoutComponentController(lcu usecase.ILayoutComponentUsecase) ILayoutComponentController {
	return &layoutComponentController{lcu}
}

func (lcc *layoutComponentController) GetAllLayoutComponents(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	componentsRes, err := lcc.lcu.GetAllLayoutComponents(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, componentsRes)
}

func (lcc *layoutComponentController) GetLayoutComponentById(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("componentId")
	componentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	componentRes, err := lcc.lcu.GetLayoutComponentById(userId, uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, componentRes)
}

func (lcc *layoutComponentController) CreateLayoutComponent(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	var request model.LayoutComponentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	componentRes, err := lcc.lcu.CreateLayoutComponent(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, componentRes)
}

func (lcc *layoutComponentController) UpdateLayoutComponent(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("componentId")
	componentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	var request model.LayoutComponentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	componentRes, err := lcc.lcu.UpdateLayoutComponent(request, userId, uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, componentRes)
}

func (lcc *layoutComponentController) DeleteLayoutComponent(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("componentId")
	componentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	err = lcc.lcu.DeleteLayoutComponent(userId, uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (lcc *layoutComponentController) AssignToLayout(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	componentId, err := strconv.ParseUint(c.Param("componentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	var request model.AssignLayoutRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	err = lcc.lcu.AssignToLayout(userId, uint(componentId), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.NoContent(http.StatusOK)
}

func (lcc *layoutComponentController) RemoveFromLayout(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	componentId, err := strconv.ParseUint(c.Param("componentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	err = lcc.lcu.RemoveFromLayout(userId, uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.NoContent(http.StatusOK)
}

func (lcc *layoutComponentController) UpdatePosition(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	componentId, err := strconv.ParseUint(c.Param("componentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	var position model.PositionRequest
	if err := c.Bind(&position); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	err = lcc.lcu.UpdatePosition(userId, uint(componentId), position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.NoContent(http.StatusOK)
}
