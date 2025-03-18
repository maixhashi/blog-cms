package layout_component_test

import (
	"testing"
)

func TestLayoutComponentValidator_ValidateAssignLayoutRequest(t *testing.T) {
	setupLayoutComponentValidatorTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効なレイアウト割り当てリクエストの場合はエラーを返さない", func(t *testing.T) {
			// テスト用の有効なリクエストを作成
			validRequest := createValidAssignLayoutRequest()

			// テスト実行
			err := layoutComponentValidator.ValidateAssignLayoutRequest(validRequest)

			// 検証
			if err != nil {
				t.Errorf("ValidateAssignLayoutRequest() error = %v, want nil", err)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("レイアウトIDが0の場合はエラーを返す", func(t *testing.T) {
			// レイアウトIDが0のリクエスト
			invalidRequest := createValidAssignLayoutRequest()
			invalidRequest.LayoutId = 0

			// テスト実行
			err := layoutComponentValidator.ValidateAssignLayoutRequest(invalidRequest)

			// 検証
			if err == nil {
				t.Error("ValidateAssignLayoutRequest() error = nil, want error for zero layout ID")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}
