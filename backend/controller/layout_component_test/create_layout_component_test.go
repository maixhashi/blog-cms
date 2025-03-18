package layout_component_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLayoutComponent(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Mock data
	requestJSON := `{"name":"New Component","type":"text","content":"Sample content"}`
	
	componentRequest := model.LayoutComponentRequest{
		Name:    "New Component",
		Type:    "text",
		Content: "Sample content",
		UserId:  1,
	}
	
	componentResponse := model.LayoutComponentResponse{
		ID:        1,
		Name:      "New Component",
		Type:      "text",
		Content:   "Sample content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Expectations
	mockUsecase.On("CreateLayoutComponent", componentRequest).Return(componentResponse, nil)
	
	// Test
	c, rec := setupContext(http.MethodPost, "/components", requestJSON)
	
	// Assertions
	if assert.NoError(t, controller.CreateLayoutComponent(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		
		var response model.LayoutComponentResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "New Component", response.Name)
		assert.Equal(t, "text", response.Type)
		assert.Equal(t, "Sample content", response.Content)
	}
	
	mockUsecase.AssertExpectations(t)
}
