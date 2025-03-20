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

// GetAllLayoutComponents ユーザーのすべてのレイアウトコンポーネントを取得
// @Summary ユーザーのレイアウトコンポーネント一覧を取得
// @Description ログインユーザーのすべてのレイアウトコンポーネントを取得する
// @Tags layout-components
// @Accept json
// @Produce json
// @Success 200 {array} model.LayoutComponentResponse
// @Failure 500 {object} map[string]string
// @Router /layout-components [get]
func (lcc *layoutComponentController) GetAllLayoutComponents(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	componentsRes, err := lcc.lcu.GetAllLayoutComponents(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, componentsRes)
}

// GetLayoutComponentById 指定されたIDのレイアウトコンポーネントを取得
// @Summary 特定のレイアウトコンポーネントを取得
// @Description 指定されたIDのレイアウトコンポーネントを取得する
// @Tags layout-components
// @Accept json
// @Produce json
// @Param componentId path int true "コンポーネントID"
// @Success 200 {object} model.LayoutComponentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components/{componentId} [get]
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

// CreateLayoutComponent 新しいレイアウトコンポーネントを作成
// @Summary 新しいレイアウトコンポーネントを作成
// @Description ユーザーの新しいレイアウトコンポーネントを作成する
// @Tags layout-components
// @Accept json
// @Produce json
// @Param component body model.LayoutComponentRequest true "コンポーネント情報"
// @Success 201 {object} model.LayoutComponentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components [post]
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

// UpdateLayoutComponent 既存のレイアウトコンポーネントを更新
// @Summary レイアウトコンポーネントを更新
// @Description 指定されたIDのレイアウトコンポーネントを更新する
// @Tags layout-components
// @Accept json
// @Produce json
// @Param componentId path int true "コンポーネントID"
// @Param component body model.LayoutComponentRequest true "更新するコンポーネント情報"
// @Success 200 {object} model.LayoutComponentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components/{componentId} [put]
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

// DeleteLayoutComponent レイアウトコンポーネントを削除
// @Summary レイアウトコンポーネントを削除
// @Description 指定されたIDのレイアウトコンポーネントを削除する
// @Tags layout-components
// @Accept json
// @Produce json
// @Param componentId path int true "コンポーネントID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components/{componentId} [delete]
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

// AssignToLayout レイアウトコンポーネントをレイアウトに割り当て
// @Summary コンポーネントをレイアウトに割り当て
// @Description 指定されたコンポーネントを特定のレイアウトに割り当てる
// @Tags layout-components
// @Accept json
// @Produce json
// @Param componentId path int true "コンポーネントID"
// @Param layoutId path int true "レイアウトID"
// @Param position body model.AssignLayoutRequest true "割り当て情報"
// @Success 200 "OK"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components/{componentId}/assign/{layoutId} [post]
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

// RemoveFromLayout レイアウトコンポーネントをレイアウトから削除
// @Summary コンポーネントをレイアウトから削除
// @Description 指定されたコンポーネントをレイアウトから削除する
// @Tags layout-components
// @Accept json
// @Produce json
// @Param componentId path int true "コンポーネントID"
// @Success 200 "OK"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components/{componentId}/assign [delete]
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

// UpdatePosition レイアウトコンポーネントの位置を更新
// @Summary コンポーネントの位置を更新
// @Description 指定されたコンポーネントの位置情報を更新する
// @Tags layout-components
// @Accept json
// @Produce json
// @Param componentId path int true "コンポーネントID"
// @Param position body model.PositionRequest true "位置情報"
// @Success 200 "OK"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layout-components/{componentId}/position [put]
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