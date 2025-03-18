package layout_component_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLayoutComponentById(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Mock data
	component := model.LayoutComponentResponse{
		ID:        1,
		Name:      "Test Component",
		Type:      "text",
		Content:   "Test Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Expectations
	mockUsecase.On("GetLayoutComponentById", uint(1), uint(1)).Return(component, nil)
	
	// Test
	c, rec := setupContext(http.MethodGet, "/components/1", "")
	c.SetParamNames("componentId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.GetLayoutComponentById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var response model.LayoutComponentResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "Test Component", response.Name)
		assert.Equal(t, "text", response.Type)
		assert.Equal(t, "Test Content", response.Content)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestGetLayoutComponentByIdInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodGet, "/components/invalid", "")
	c.SetParamNames("componentId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.GetLayoutComponentById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なコンポーネントIDです", response["error"])
	}
}