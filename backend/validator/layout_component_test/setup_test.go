package layout_component_test

import (
	"go-react-app/model"
	"go-react-app/validator"
)

// テスト用の共通変数
var (
	layoutComponentValidator validator.ILayoutComponentValidator
)

// テスト前の共通セットアップ
func setupLayoutComponentValidatorTest() {
	layoutComponentValidator = validator.NewLayoutComponentValidator()
}

// テスト用のレイアウトコンポーネントリクエストを作成
func createValidLayoutComponentRequest() model.LayoutComponentRequest {
	return model.LayoutComponentRequest{
		Name:    "テストコンポーネント",
		Type:    "text",
		Content: "テストコンテンツ",
		X:       0,
		Y:       0,
		Width:   100,
		Height:  100,
		UserId:  1,
	}
}

// テスト用のレイアウト割り当てリクエストを作成
func createValidAssignLayoutRequest() model.AssignLayoutRequest {
	return model.AssignLayoutRequest{
		LayoutId: 1,
		Position: model.PositionRequest{
			X:      10,
			Y:      20,
			Width:  200,
			Height: 150,
		},
	}
}

// テスト用の位置情報リクエストを作成
func createValidPositionRequest() model.PositionRequest {
	return model.PositionRequest{
		X:      10,
		Y:      20,
		Width:  200,
		Height: 150,
	}
}