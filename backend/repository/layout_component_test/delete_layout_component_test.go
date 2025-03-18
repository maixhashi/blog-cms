package layout_component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_DeleteLayoutComponent(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("レイアウトコンポーネントを正常に削除できる", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		err = layoutComponentRepo.DeleteLayoutComponent(testUserData.ID, component.ID)
		
		assert.NoError(t, err)
		
		// 削除されたことを確認
		_, err = layoutComponentRepo.GetLayoutComponentById(testUserData.ID, component.ID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
		err := layoutComponentRepo.DeleteLayoutComponent(testUserData.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		err = layoutComponentRepo.DeleteLayoutComponent(9999, component.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})
}
