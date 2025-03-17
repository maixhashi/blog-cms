package layout_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteLayout(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Expectations
	mockUsecase.On("DeleteLayout", uint(1), uint(1)).Return(nil)
	
	// Test
	c, rec := setupContext(http.MethodDelete, "/layouts/1", "")
	c.SetParamNames("layoutId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.DeleteLayout(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestDeleteLayoutInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutUsecase)
	controller := newMockLayoutController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodDelete, "/layouts/invalid", "")
	c.SetParamNames("layoutId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.DeleteLayout(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なレイアウトIDです", response["error"])
	}
}