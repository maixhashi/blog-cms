package layout_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLayoutById(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Mock data
	layout := model.LayoutResponse{
		ID:        1,
		Title:     "Test Layout",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Expectations
	mockUsecase.On("GetLayoutById", uint(1), uint(1)).Return(layout, nil)
	
	// Test
	c, rec := setupContext(http.MethodGet, "/layouts/1", "")
	c.SetParamNames("layoutId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.GetLayoutById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var response model.LayoutResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "Test Layout", response.Title)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestGetLayoutByIdInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodGet, "/layouts/invalid", "")
	c.SetParamNames("layoutId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.GetLayoutById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なレイアウトIDです", response["error"])
	}
}
