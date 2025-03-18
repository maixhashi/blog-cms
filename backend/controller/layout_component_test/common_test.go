package layout_component_test

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// Mock for layout component usecase
type MockLayoutComponentUsecase struct {
	mock.Mock
}

func (m *MockLayoutComponentUsecase) GetAllLayoutComponents(userId uint) ([]model.LayoutComponentResponse, error) {
	args := m.Called(userId)
	return args.Get(0).([]model.LayoutComponentResponse), args.Error(1)
}

func (m *MockLayoutComponentUsecase) GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponentResponse, error) {
	args := m.Called(userId, componentId)
	return args.Get(0).(model.LayoutComponentResponse), args.Error(1)
}

func (m *MockLayoutComponentUsecase) CreateLayoutComponent(request model.LayoutComponentRequest) (model.LayoutComponentResponse, error) {
	args := m.Called(request)
	return args.Get(0).(model.LayoutComponentResponse), args.Error(1)
}

func (m *MockLayoutComponentUsecase) UpdateLayoutComponent(request model.LayoutComponentRequest, userId uint, componentId uint) (model.LayoutComponentResponse, error) {
	args := m.Called(request, userId, componentId)
	return args.Get(0).(model.LayoutComponentResponse), args.Error(1)
}

func (m *MockLayoutComponentUsecase) DeleteLayoutComponent(userId uint, componentId uint) error {
	args := m.Called(userId, componentId)
	return args.Error(0)
}

func (m *MockLayoutComponentUsecase) AssignToLayout(userId uint, componentId uint, request model.AssignLayoutRequest) error {
	args := m.Called(userId, componentId, request)
	return args.Error(0)
}

func (m *MockLayoutComponentUsecase) RemoveFromLayout(userId uint, componentId uint) error {
	args := m.Called(userId, componentId)
	return args.Error(0)
}

func (m *MockLayoutComponentUsecase) UpdatePosition(userId uint, componentId uint, position model.PositionRequest) error {
	args := m.Called(userId, componentId, position)
	return args.Error(0)
}

// Mock the getUserIdFromToken function for testing
type mockLayoutComponentController struct {
	lcu usecase.ILayoutComponentUsecase
}

func newMockLayoutComponentController(lcu usecase.ILayoutComponentUsecase) *mockLayoutComponentController {
	return &mockLayoutComponentController{lcu}
}

func (lcc *mockLayoutComponentController) GetAllLayoutComponents(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
	componentsRes, err := lcc.lcu.GetAllLayoutComponents(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, componentsRes)
}

func (lcc *mockLayoutComponentController) GetLayoutComponentById(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lcc *mockLayoutComponentController) CreateLayoutComponent(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lcc *mockLayoutComponentController) UpdateLayoutComponent(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lcc *mockLayoutComponentController) DeleteLayoutComponent(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
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

func (lcc *mockLayoutComponentController) AssignToLayout(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
	id := c.Param("componentId")
	componentId, err := strconv.ParseUint(id, 10, 32)
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

func (lcc *mockLayoutComponentController) RemoveFromLayout(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
	id := c.Param("componentId")
	componentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なコンポーネントIDです"})
	}
	
	err = lcc.lcu.RemoveFromLayout(userId, uint(componentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (lcc *mockLayoutComponentController) UpdatePosition(c echo.Context) error {
	userId := uint(1) // Hardcoded for testing
	
	id := c.Param("componentId")
	componentId, err := strconv.ParseUint(id, 10, 32)
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
