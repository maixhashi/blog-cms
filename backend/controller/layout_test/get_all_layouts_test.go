package layout_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllLayouts(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Mock data
	layouts := []model.LayoutResponse{
		{
			ID:        1,
			Title:     "Test Layout 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Test Layout 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	
	// Expectations
	mockUsecase.On("GetAllLayouts", uint(1)).Return(layouts, nil)
	
	// Test
	c, rec := setupContext(http.MethodGet, "/layouts", "")
	
	// Assertions
	if assert.NoError(t, controller.GetAllLayouts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var response []model.LayoutResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
		assert.Equal(t, "Test Layout 1", response[0].Title)
		assert.Equal(t, "Test Layout 2", response[1].Title)
	}
	
	mockUsecase.AssertExpectations(t)
}
