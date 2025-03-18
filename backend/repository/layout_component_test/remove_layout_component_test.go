package layout_component_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_RemoveFromLayout(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("コンポーネントをレイアウトから正常に削除できる", func(t *testing.T) {
		// まずコンポーネントを作成してレイアウトに割り当て
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		position := model.PositionRequest{X: 10, Y: 20}
		err = layoutComponentRepo.AssignToLayout(component.ID, testLayoutData.ID, testUserData.ID, position)
		assert.NoError(t, err)
		
		// レイアウトから削除
		err = layoutComponentRepo.RemoveFromLayout(component.ID, testUserData.ID)
		
		assert.NoError(t, err)
		
		// 削除が反映されたか確認
		updatedComponent, err := layoutComponentRepo.GetLayoutComponentById(testUserData.ID, component.ID)
		assert.NoError(t, err)
		assert.Nil(t, updatedComponent.LayoutId)
		assert.Equal(t, 0, updatedComponent.X)
		assert.Equal(t, 0, updatedComponent.Y)
	})

	t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
		err := layoutComponentRepo.RemoveFromLayout(9999, testUserData.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		err = layoutComponentRepo.RemoveFromLayout(component.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})
}
