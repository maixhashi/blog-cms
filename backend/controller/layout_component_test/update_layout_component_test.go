package layout_component_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateLayoutComponent(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Mock data
	requestJSON := `{"name":"Updated Component","type":"text","content":"Updated content"}`
	
	componentRequest := model.LayoutComponentRequest{
		Name:    "Updated Component",
		Type:    "text",
		Content: "Updated content",
		UserId:  1,
	}
	
	componentResponse := model.LayoutComponentResponse{
		ID:        1,
		Name:      "Updated Component",
		Type:      "text",
		Content:   "Updated content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Expectations
	mockUsecase.On("UpdateLayoutComponent", componentRequest, uint(1), uint(1)).Return(componentResponse, nil)
	
	// Test
	c, rec := setupContext(http.MethodPut, "/components/1", requestJSON)
	c.SetParamNames("componentId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.UpdateLayoutComponent(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var response model.LayoutComponentResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "Updated Component", response.Name)
		assert.Equal(t, "text", response.Type)
		assert.Equal(t, "Updated content", response.Content)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestUpdateLayoutComponentInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodPut, "/components/invalid", `{"name":"Updated Component","type":"text"}`)
	c.SetParamNames("componentId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.UpdateLayoutComponent(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なコンポーネントIDです", response["error"])
	}
}
