package layout_component_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePosition(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Mock data
	requestJSON := `{"x":15,"y":25,"width":400,"height":300}`
	
	positionRequest := model.PositionRequest{
		X:      15,
		Y:      25,
		Width:  400,
		Height: 300,
	}
	
	// Expectations
	mockUsecase.On("UpdatePosition", uint(1), uint(1), positionRequest).Return(nil)
	
	// Test
	c, rec := setupContext(http.MethodPut, "/components/1/position", requestJSON)
	c.SetParamNames("componentId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.UpdatePosition(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestUpdatePositionInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodPut, "/components/invalid/position", `{"x":15,"y":25}`)
	c.SetParamNames("componentId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.UpdatePosition(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なコンポーネントIDです", response["error"])
	}
}
