package layout_test

import (
	"go-react-app/model"
)

// テストヘルパー関数
func createTestLayout() (*model.Layout, error) {
	layout := testLayoutData
	err := layoutRepo.CreateLayout(&layout)
	return &layout, err
}

// テストデータジェネレーター
func generateUniqueTestLayout(title string, userId uint) model.Layout {
	return model.Layout{
		Title:  title,
		UserId: userId,
	}
}

// テストアサーション用のヘルパー
func validateLayoutFields(layout *model.Layout) bool {
	return layout.ID != 0 &&
		layout.Title != "" &&
		layout.UserId != 0 &&
		!layout.CreatedAt.IsZero() &&
		!layout.UpdatedAt.IsZero()
}
