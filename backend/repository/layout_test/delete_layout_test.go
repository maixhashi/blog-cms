package layout_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutRepository_DeleteLayout(t *testing.T) {
	setupLayoutRepositoryTest()

	t.Run("レイアウトを正常に削除できる", func(t *testing.T) {
		layout, err := createTestLayout()
		assert.NoError(t, err)
		
		err = layoutRepo.DeleteLayout(testUserData.ID, layout.ID)
		
		assert.NoError(t, err)
		
		// 削除されたことを確認
		_, err = layoutRepo.GetLayoutById(testUserData.ID, layout.ID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
		err := layoutRepo.DeleteLayout(testUserData.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout does not exist")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		layout, err := createTestLayout()
		assert.NoError(t, err)
		
		err = layoutRepo.DeleteLayout(9999, layout.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout does not exist")
	})
}