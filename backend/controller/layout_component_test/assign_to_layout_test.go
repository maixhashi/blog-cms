package layout_component_test

import (
	"encoding/json"
	"go-react-app/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignToLayout(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Mock data
	requestJSON := `{"layout_id":2,"position":{"x":10,"y":20,"width":300,"height":200}}`
	
	assignRequest := model.AssignLayoutRequest{
		LayoutId: 2,
		Position: model.PositionRequest{
			X:      10,
			Y:      20,
			Width:  300,
			Height: 200,
		},
	}
	
	// Expectations
	mockUsecase.On("AssignToLayout", uint(1), uint(1), assignRequest).Return(nil)
	
	// Test
	c, rec := setupContext(http.MethodPost, "/components/1/assign", requestJSON)
	c.SetParamNames("componentId")
	c.SetParamValues("1")
	
	// Assertions
	if assert.NoError(t, controller.AssignToLayout(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	
	mockUsecase.AssertExpectations(t)
}

func TestAssignToLayoutInvalidId(t *testing.T) {
	// Setup
	mockUsecase := new(MockLayoutComponentUsecase)
	controller := newMockLayoutComponentController(mockUsecase)
	
	// Test
	c, rec := setupContext(http.MethodPost, "/components/invalid/assign", `{"layout_id":2}`)
	c.SetParamNames("componentId")
	c.SetParamValues("invalid")
	
	// Assertions
	if assert.NoError(t, controller.AssignToLayout(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		
		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "無効なコンポーネントIDです", response["error"])
	}
}
