package layout_component_test

import (
	"go-react-app/model"
)

// テストヘルパー関数
func createTestLayoutComponent(userId uint) (*model.LayoutComponent, error) {
	component := testLayoutComponent
	component.UserId = userId
	err := lcRepo.CreateLayoutComponent(&component)
	return &component, err
}

// テストデータジェネレーター
func generateUniqueTestComponent(name string, componentType string, userId uint) model.LayoutComponent {
	return model.LayoutComponent{
		Name:    name,
		Type:    componentType,
		Content: `{"data": "test content"}`,
		UserId:  userId,
	}
}

// テストアサーション用のヘルパー
func validateComponentFields(component *model.LayoutComponent) bool {
	return component.ID != 0 &&
		component.Name != "" &&
		component.Type != "" &&
		component.UserId != 0 &&
		!component.CreatedAt.IsZero() &&
		!component.UpdatedAt.IsZero()
}
