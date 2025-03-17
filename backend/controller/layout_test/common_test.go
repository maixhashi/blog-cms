package layout_test

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// Mock for layout usecase
type MockLayoutUsecase struct {
	mock.Mock
}

func (m *MockLayoutUsecase) GetAllLayouts(userId uint) ([]model.LayoutResponse, error) {
	args := m.Called(userId)
	return args.Get(0).([]model.LayoutResponse), args.Error(1)
}

func (m *MockLayoutUsecase) GetLayoutById(userId uint, layoutId uint) (model.LayoutResponse, error) {
	args := m.Called(userId, layoutId)
	return args.Get(0).(model.LayoutResponse), args.Error(1)
}

func (m *MockLayoutUsecase) CreateLayout(request model.LayoutRequest) (model.LayoutResponse, error) {
	args := m.Called(request)
	return args.Get(0).(model.LayoutResponse), args.Error(1)
}

func (m *MockLayoutUsecase) UpdateLayout(request model.LayoutRequest, userId uint, layoutId uint) (model.LayoutResponse, error) {
	args := m.Called(request, userId, layoutId)
	return args.Get(0).(model.LayoutResponse), args.Error(1)
}

func (m *MockLayoutUsecase) DeleteLayout(userId uint, layoutId uint) error {
	args := m.Called(userId, layoutId)
	return args.Error(0)
}

// Mock the getUserIdFromToken function for testing
type mockLayoutController struct {
	lu usecase.ILayoutUsecase
}

func newMockLayoutController(lu usecase.ILayoutUsecase) *mockLayoutController {
	return &mockLayoutController{lu}
}

func (lc *mockLayoutController) GetAllLayouts(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
	layoutsRes, err := lc.lu.GetAllLayouts(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, layoutsRes)
}

func (lc *mockLayoutController) GetLayoutById(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lc *mockLayoutController) CreateLayout(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lc *mockLayoutController) UpdateLayout(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lc *mockLayoutController) DeleteLayout(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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
