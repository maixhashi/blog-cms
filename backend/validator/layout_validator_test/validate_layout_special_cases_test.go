package layout_validator_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutValidator_ValidateLayoutRequest_SpecialCases(t *testing.T) {
	setupLayoutValidatorTest()

	t.Run("特殊文字を含むタイトルの場合も有効", func(t *testing.T) {
		// 特殊文字を含むタイトル
		specialCharsRequest := model.LayoutRequest{
			Title:  "Layout with special chars: !@#$%^&*()_+",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(specialCharsRequest)

		// 検証
		assert.NoError(t, err, "特殊文字を含むタイトルでエラーが発生しました")
	})

	t.Run("日本語のタイトルの場合も有効", func(t *testing.T) {
		// 日本語のタイトル
		japaneseRequest := model.LayoutRequest{
			Title:  "日本語のレイアウトタイトル",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(japaneseRequest)

		// 検証
		assert.NoError(t, err, "日本語のタイトルでエラーが発生しました")
	})

	t.Run("絵文字を含むタイトルの場合も有効", func(t *testing.T) {
		// 絵文字を含むタイトル
		emojiRequest := model.LayoutRequest{
			Title:  "Layout with emoji 😊🎉👍",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(emojiRequest)

		// 検証
		assert.NoError(t, err, "絵文字を含むタイトルでエラーが発生しました")
	})
}
