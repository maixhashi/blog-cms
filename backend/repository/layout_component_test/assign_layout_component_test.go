package layout_component_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_AssignToLayout(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("コンポーネントをレイアウトに正常に割り当てできる", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		position := model.PositionRequest{
			X:      10,
			Y:      20,
			Width:  200,
			Height: 150,
		}
		
		err = layoutComponentRepo.AssignToLayout(component.ID, testLayoutData.ID, testUserData.ID, position)
		
		assert.NoError(t, err)
		
		// 割り当てが反映されたか確認
		updatedComponent, err := layoutComponentRepo.GetLayoutComponentById(testUserData.ID, component.ID)
		assert.NoError(t, err)
		assert.NotNil(t, updatedComponent.LayoutId)
		assert.Equal(t, testLayoutData.ID, *updatedComponent.LayoutId)
		assert.Equal(t, position.X, updatedComponent.X)
		assert.Equal(t, position.Y, updatedComponent.Y)
		assert.Equal(t, position.Width, updatedComponent.Width)
		assert.Equal(t, position.Height, updatedComponent.Height)
	})

	t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
		position := model.PositionRequest{X: 10, Y: 20}
		
		err := layoutComponentRepo.AssignToLayout(9999, testLayoutData.ID, testUserData.ID, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		position := model.PositionRequest{X: 10, Y: 20}
		
		err = layoutComponentRepo.AssignToLayout(component.ID, 9999, testUserData.ID, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout does not exist")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		position := model.PositionRequest{X: 10, Y: 20}
		
		err = layoutComponentRepo.AssignToLayout(component.ID, testLayoutData.ID, 9999, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})
}