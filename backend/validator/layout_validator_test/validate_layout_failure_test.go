package layout_validator_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutValidator_ValidateLayoutRequest_Failure(t *testing.T) {
	setupLayoutValidatorTest()

	t.Run("タイトルが空の場合はエラーを返す", func(t *testing.T) {
		// 空のタイトル
		emptyTitleRequest := model.LayoutRequest{
			Title:  "",
			UserId: 1,
		}

		// テスト実行
		err := layoutValidator.ValidateLayoutRequest(emptyTitleRequest)

		// 検証
		assert.Error(t, err, "空のタイトルでエラーが返されませんでした")
		assert.Contains(t, err.Error(), "タイトルは必須です", "期待されるエラーメッセージが含まれていません")
	})
}
