package layout_component_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_UpdatePosition(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("コンポーネントの位置を正常に更新できる", func(t *testing.T) {
		// まずコンポーネントを作成してレイアウトに割り当て
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		initialPosition := model.PositionRequest{X: 10, Y: 20}
		err = layoutComponentRepo.AssignToLayout(component.ID, testLayoutData.ID, testUserData.ID, initialPosition)
		assert.NoError(t, err)
		
		// 位置を更新
		newPosition := model.PositionRequest{
			X:      30,
			Y:      40,
			Width:  250,
			Height: 180,
		}
		
		err = layoutComponentRepo.UpdatePosition(component.ID, testUserData.ID, newPosition)
		
		assert.NoError(t, err)
		
		// 更新が反映されたか確認
		updatedComponent, err := layoutComponentRepo.GetLayoutComponentById(testUserData.ID, component.ID)
		assert.NoError(t, err)
		assert.Equal(t, newPosition.X, updatedComponent.X)
		assert.Equal(t, newPosition.Y, updatedComponent.Y)
		assert.Equal(t, newPosition.Width, updatedComponent.Width)
		assert.Equal(t, newPosition.Height, updatedComponent.Height)
	})

	t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
		position := model.PositionRequest{X: 10, Y: 20}
		
		err := layoutComponentRepo.UpdatePosition(9999, testUserData.ID, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		position := model.PositionRequest{X: 10, Y: 20}
		
		err = layoutComponentRepo.UpdatePosition(component.ID, 9999, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("レイアウトに割り当てられていないコンポーネントの場合はエラーを返す", func(t *testing.T) {
		// レイアウトに割り当てていないコンポーネントを作成
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		position := model.PositionRequest{X: 10, Y: 20}
		
		err = layoutComponentRepo.UpdatePosition(component.ID, testUserData.ID, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "component is not assigned to any layout")
	})
}
