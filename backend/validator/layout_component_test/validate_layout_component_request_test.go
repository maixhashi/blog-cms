package layout_component_test

import (
	"testing"
)

func TestLayoutComponentValidator_ValidateLayoutComponentRequest(t *testing.T) {
	setupLayoutComponentValidatorTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効なレイアウトコンポーネントリクエストの場合はエラーを返さない", func(t *testing.T) {
			// テスト用の有効なリクエストを作成
			validRequest := createValidLayoutComponentRequest()

			// テスト実行
			err := layoutComponentValidator.ValidateLayoutComponentRequest(validRequest)

			// 検証
			if err != nil {
				t.Errorf("ValidateLayoutComponentRequest() error = %v, want nil", err)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("名前が空の場合はエラーを返す", func(t *testing.T) {
			// 名前が空のリクエスト
			invalidRequest := createValidLayoutComponentRequest()
			invalidRequest.Name = ""

			// テスト実行
			err := layoutComponentValidator.ValidateLayoutComponentRequest(invalidRequest)

			// 検証
			if err == nil {
				t.Error("ValidateLayoutComponentRequest() error = nil, want error for empty name")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})

		t.Run("タイプが空の場合はエラーを返す", func(t *testing.T) {
			// タイプが空のリクエスト
			invalidRequest := createValidLayoutComponentRequest()
			invalidRequest.Type = ""

			// テスト実行
			err := layoutComponentValidator.ValidateLayoutComponentRequest(invalidRequest)

			// 検証
			if err == nil {
				t.Error("ValidateLayoutComponentRequest() error = nil, want error for empty type")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
