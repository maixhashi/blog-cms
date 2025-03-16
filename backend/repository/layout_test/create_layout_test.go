package layout_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutRepository_CreateLayout(t *testing.T) {
	setupLayoutRepositoryTest()

	t.Run("レイアウトを正常に作成できる", func(t *testing.T) {
		layout := testLayoutData
		err := layoutRepo.CreateLayout(&layout)
		
		assert.NoError(t, err)
		assert.NotZero(t, layout.ID)
		assert.NotZero(t, layout.CreatedAt)
		assert.NotZero(t, layout.UpdatedAt)
		assert.Equal(t, testUserData.ID, layout.UserId)
	})

	t.Run("ユーザーIDが存在しない場合でもエラーにならない", func(t *testing.T) {
		layout := model.Layout{
			Title:  "Invalid User Layout",
			UserId: 9999, // 存在しないユーザーID
		}
		
		err := layoutRepo.CreateLayout(&layout)
		// 現在の実装ではエラーが発生しないことを確認
		assert.NoError(t, err)
		assert.NotZero(t, layout.ID)
	})
}