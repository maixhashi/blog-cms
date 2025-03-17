package layout_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLayout(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Mock data
	requestJSON := `{"title":"New Layout"}`
	
	layoutRequest := model.LayoutRequest{
		Title:  "New Layout",
		UserId: 1,
	}
	
	layoutResponse := model.LayoutResponse{
		ID:        1,
		Title:     "New Layout",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Expectations
	mockUsecase.On("CreateLayout", layoutRequest).Return(layoutResponse, nil)
	
	// Test
	c, rec := setupContext(http.MethodPost, "/layouts", requestJSON)
	
	// Assertions
	if assert.NoError(t, controller.CreateLayout(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		
		var response model.LayoutResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "New Layout", response.Title)
	}
	
	mockUsecase.AssertExpectations(t)
}
