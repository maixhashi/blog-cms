package testutils

import (
	"go-react-app/model"
)

// MockLayoutUsecase はレイアウトユースケースのモック
type MockLayoutUsecase struct {
	// モックメソッドの呼び出し結果を保存
	GetAllLayoutsFunc      func(userId uint) ([]model.LayoutResponse, error)
	GetLayoutByIdFunc      func(userId uint, layoutId uint) (model.LayoutResponse, error)
	CreateLayoutFunc       func(request model.LayoutRequest) (model.LayoutResponse, error)
	UpdateLayoutFunc       func(request model.LayoutRequest, userId uint, layoutId uint) (model.LayoutResponse, error)
	DeleteLayoutFunc       func(userId uint, layoutId uint) error
}

// GetAllLayouts はモックメソッド
func (m *MockLayoutUsecase) GetAllLayouts(userId uint) ([]model.LayoutResponse, error) {
	return m.GetAllLayoutsFunc(userId)
}

// GetLayoutById はモックメソッド
func (m *MockLayoutUsecase) GetLayoutById(userId uint, layoutId uint) (model.LayoutResponse, error) {
	return m.GetLayoutByIdFunc(userId, layoutId)
}

// CreateLayout はモックメソッド
func (m *MockLayoutUsecase) CreateLayout(request model.LayoutRequest) (model.LayoutResponse, error) {
	return m.CreateLayoutFunc(request)
}

// UpdateLayout はモックメソッド
func (m *MockLayoutUsecase) UpdateLayout(request model.LayoutRequest, userId uint, layoutId uint) (model.LayoutResponse, error) {
	return m.UpdateLayoutFunc(request, userId, layoutId)
}

// DeleteLayout はモックメソッド
func (m *MockLayoutUsecase) DeleteLayout(userId uint, layoutId uint) error {
	return m.DeleteLayoutFunc(userId, layoutId)
}
