package layout_component_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllLayoutComponents(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Mock data
	components := []model.LayoutComponentResponse{
		{
			ID:        1,
			Name:      "Component 1",
			Type:      "text",
			Content:   "Content 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Component 2",
			Type:      "image",
			Content:   "Content 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	
	// Expectations
	mockUsecase.On("GetAllLayoutComponents", uint(1)).Return(components, nil)
	
	// Test
	c, rec := setupContext(http.MethodGet, "/components", "")
	
	// Assertions
	if assert.NoError(t, controller.GetAllLayoutComponents(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var response []model.LayoutComponentResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, "Component 1", response[0].Name)
		assert.Equal(t, "Component 2", response[1].Name)
	}
	
	mockUsecase.AssertExpectations(t)
}
