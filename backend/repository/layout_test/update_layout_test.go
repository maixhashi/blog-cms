package layout_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutRepository_UpdateLayout(t *testing.T) {
	setupLayoutRepositoryTest()

	t.Run("レイアウトを正常に更新できる", func(t *testing.T) {
		layout, err := createTestLayout()
		assert.NoError(t, err)
		
		updatedTitle := "Updated Layout Title"
		updatedLayout := *layout
		updatedLayout.Title = updatedTitle
		
		err = layoutRepo.UpdateLayout(&updatedLayout, testUserData.ID, layout.ID)
		
		assert.NoError(t, err)
		assert.Equal(t, updatedTitle, updatedLayout.Title)
	})

	t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
		layout := testLayoutData
		layout.Title = "Non-existent Layout"
		
		err := layoutRepo.UpdateLayout(&layout, testUserData.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout does not exist")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		layout, err := createTestLayout()
		assert.NoError(t, err)
		
		updatedLayout := *layout
		updatedLayout.Title = "Unauthorized Update"
		
		err = layoutRepo.UpdateLayout(&updatedLayout, 9999, layout.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout does not exist")
	})
}