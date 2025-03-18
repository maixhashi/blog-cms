package layout_component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_UpdateLayoutComponent(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("レイアウトコンポーネントを正常に更新できる", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		updatedName := "Updated Component Name"
		updatedComponent := *component
		updatedComponent.Name = updatedName
		
		err = layoutComponentRepo.UpdateLayoutComponent(&updatedComponent, testUserData.ID, component.ID)
		
		assert.NoError(t, err)
		assert.Equal(t, updatedName, updatedComponent.Name)
	})

	t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
		component := testLayoutComponentData
		component.Name = "Non-existent Component"
		
		err := layoutComponentRepo.UpdateLayoutComponent(&component, testUserData.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		updatedComponent := *component
		updatedComponent.Name = "Unauthorized Update"
		
		err = layoutComponentRepo.UpdateLayoutComponent(&updatedComponent, 9999, component.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})
}