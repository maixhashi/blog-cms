package layout_component_test

import (
	"testing"
)

func TestLayoutComponentValidator_ValidatePositionRequest(t *testing.T) {
	setupLayoutComponentValidatorTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効な位置情報リクエストの場合はエラーを返さない", func(t *testing.T) {
			// テスト用の有効なリクエストを作成
			validRequest := createValidPositionRequest()

			// テスト実行
			err := layoutComponentValidator.ValidatePositionRequest(validRequest)

			// 検証
			if err != nil {
				t.Errorf("ValidatePositionRequest() error = %v, want nil", err)
			}
		})

		t.Run("幅と高さが0でも有効", func(t *testing.T) {
			// 幅と高さが0のリクエスト
			request := createValidPositionRequest()
			request.Width = 0
			request.Height = 0

			// テスト実行
			err := layoutComponentValidator.ValidatePositionRequest(request)

			// 検証
			if err != nil {
				t.Errorf("ValidatePositionRequest() error = %v, want nil", err)
			}
		})

		t.Run("負の座標でも有効", func(t *testing.T) {
			// 負の座標を持つリクエスト
			request := createValidPositionRequest()
			request.X = -10
			request.Y = -20

			// テスト実行
			err := layoutComponentValidator.ValidatePositionRequest(request)

			// 検証
			if err != nil {
				t.Errorf("ValidatePositionRequest() error = %v, want nil", err)
			}
		})
	})

	// 現在の実装では位置情報に対するバリデーションは行われていないため、
	// 異常系のテストケースは追加していません。
	// もし将来的に位置情報のバリデーションが追加された場合は、
	// ここに異常系のテストケースを追加してください。
}
