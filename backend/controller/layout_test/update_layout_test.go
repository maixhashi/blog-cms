package layout_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateLayout(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Mock data
	requestJSON := `{"title":"Updated Layout"}`
	
	layoutRequest := model.LayoutRequest{
		Title:  "Updated Layout",
		UserId: 1,
	}
	
	layoutResponse := model.LayoutResponse{
		ID:        1,
		Title:     "Updated Layout",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Expectations
	mockUsecase.On("UpdateLayout", layoutRequest, uint(1), uint(1)).Return(layoutResponse, nil)
	
	// Test
	c, rec := setupContext(http.MethodPut, "/layouts/1", requestJSON)
	c.SetParamNames("layoutId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.UpdateLayout(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		
		var response model.LayoutResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "Updated Layout", response.Title)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestUpdateLayoutInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodPut, "/layouts/invalid", `{"title":"Updated Layout"}`)
	c.SetParamNames("layoutId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.UpdateLayout(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なレイアウトIDです", response["error"])
	}
}
