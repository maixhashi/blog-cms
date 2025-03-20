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

// GetAllLayouts ユーザーのすべてのレイアウトを取得
// @Summary ユーザーのレイアウト一覧を取得
// @Description ログインユーザーのすべてのレイアウトを取得する
// @Tags layouts
// @Accept json
// @Produce json
// @Success 200 {array} model.LayoutResponse
// @Failure 500 {object} map[string]string
// @Router /layouts [get]
func (lc *layoutController) GetAllLayouts(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	layoutsRes, err := lc.lu.GetAllLayouts(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, layoutsRes)
}

// GetLayoutById 指定されたIDのレイアウトを取得
// @Summary 特定のレイアウトを取得
// @Description 指定されたIDのレイアウトを取得する
// @Tags layouts
// @Accept json
// @Produce json
// @Param layoutId path int true "レイアウトID"
// @Success 200 {object} model.LayoutResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layouts/{layoutId} [get]
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

// CreateLayout 新しいレイアウトを作成
// @Summary 新しいレイアウトを作成
// @Description ユーザーの新しいレイアウトを作成する
// @Tags layouts
// @Accept json
// @Produce json
// @Param layout body model.LayoutRequest true "レイアウト情報"
// @Success 201 {object} model.LayoutResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layouts [post]
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

// UpdateLayout 既存のレイアウトを更新
// @Summary レイアウトを更新
// @Description 指定されたIDのレイアウトを更新する
// @Tags layouts
// @Accept json
// @Produce json
// @Param layoutId path int true "レイアウトID"
// @Param layout body model.LayoutRequest true "更新するレイアウト情報"
// @Success 200 {object} model.LayoutResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layouts/{layoutId} [put]
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

// DeleteLayout レイアウトを削除
// @Summary レイアウトを削除
// @Description 指定されたIDのレイアウトを削除する
// @Tags layouts
// @Accept json
// @Produce json
// @Param layoutId path int true "レイアウトID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /layouts/{layoutId} [delete]
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