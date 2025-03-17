package layout_validator_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutValidator_ValidateLayoutRequest_Success(t *testing.T) {
	setupLayoutValidatorTest()

	t.Run("有効なレイアウトリクエストの場合はエラーを返さない", func(t *testing.T) {
		// 有効なレイアウトリクエスト
		validRequest := model.LayoutRequest{
			Title:  "Valid Layout Title",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(validRequest)

		// 検証
		assert.NoError(t, err, "有効なレイアウトリクエストでエラーが発生しました")
	})

	t.Run("タイトルが最小長の場合でも有効", func(t *testing.T) {
		// 最小長のタイトル（1文字）
		minLengthRequest := model.LayoutRequest{
			Title:  "A",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(minLengthRequest)

		// 検証
		assert.NoError(t, err, "最小長のタイトルでエラーが発生しました")
	})

	t.Run("タイトルが長い場合でも有効", func(t *testing.T) {
		// 長いタイトル
		longTitleRequest := model.LayoutRequest{
			Title:  "This is a very long layout title that should still be valid for testing purposes",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(longTitleRequest)

		// 検証
		assert.NoError(t, err, "長いタイトルでエラーが発生しました")
	})
}
