package layout_component_test

import (
	"go-react-app/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_DeleteLayoutComponent(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("コンポーネントを正常に削除できる", func(t *testing.T) {
		component, err := createTestLayoutComponent(testUser.ID)
		assert.NoError(t, err)
		
		err = lcRepo.DeleteLayoutComponent(testUser.ID, component.ID)
		
		assert.NoError(t, err)
		
		// 削除されたことを確認
		foundComponent := &model.LayoutComponent{}
		err = lcRepo.GetLayoutComponentById(foundComponent, testUser.ID, component.ID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("存在しないコンポーネントは削除できない", func(t *testing.T) {
		nonExistentID := uint(9999)
		
		err := lcRepo.DeleteLayoutComponent(testUser.ID, nonExistentID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("他のユーザーのコンポーネントは削除できない", func(t *testing.T) {
		component, err := createTestLayoutComponent(testUser.ID)
		assert.NoError(t, err)
		
		otherUserID := uint(9999) // 存在しないユーザーID
		
		err = lcRepo.DeleteLayoutComponent(otherUserID, component.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})
}