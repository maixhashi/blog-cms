package layout_component_test

import (
	"go-react-app/model"
)

// テストヘルパー関数
func createTestLayoutComponent() (*model.LayoutComponent, error) {
	component := testLayoutComponentData
	err := layoutComponentRepo.CreateLayoutComponent(&component)
	return &component, err
}

// テストデータジェネレーター
func generateUniqueTestLayoutComponent(name string, componentType string, userId uint) model.LayoutComponent {
	return model.LayoutComponent{
		Name:    name,
		Type:    componentType,
		Content: "Generated Content",
		X:       0,
		Y:       0,
		Width:   100,
		Height:  100,
		UserId:  userId,
	}
}

// テストアサーション用のヘルパー
func validateLayoutComponentFields(component *model.LayoutComponent) bool {
	return component.ID != 0 &&
		component.Name != "" &&
		component.Type != "" &&
		component.UserId != 0 &&
		!component.CreatedAt.IsZero() &&
		!component.UpdatedAt.IsZero()
}
