package layout_component_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteLayoutComponent(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Expectations
	mockUsecase.On("DeleteLayoutComponent", uint(1), uint(1)).Return(nil)
	
	// Test
	c, rec := setupContext(http.MethodDelete, "/components/1", "")
	c.SetParamNames("componentId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.DeleteLayoutComponent(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestDeleteLayoutComponentInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodDelete, "/components/invalid", "")
	c.SetParamNames("componentId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.DeleteLayoutComponent(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なコンポーネントIDです", response["error"])
	}
}
