package layout_component_test

import (
	"go-react-app/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_LayoutOperations(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	// テスト用のレイアウトを作成する関数
	createTestLayout := func() *model.Layout {
		layout := model.Layout{
			Title:  "Test Layout", // Name を Title に変更
			UserId: testUser.ID,
		}
		lcRepoDb.Create(&layout)
		return &layout
	}

	t.Run("コンポーネントをレイアウトに割り当てられる", func(t *testing.T) {
		component, _ := createTestLayoutComponent(testUser.ID)
		layout := createTestLayout()
		
		position := map[string]int{
			"x":      10,
			"y":      20,
			"width":  300,
			"height": 200,
		}
		
		err := lcRepo.AssignToLayout(component.ID, layout.ID, testUser.ID, position)
		
		assert.NoError(t, err)
		
		// 割り当てられたことを確認
		updatedComponent := &model.LayoutComponent{}
		lcRepo.GetLayoutComponentById(updatedComponent, testUser.ID, component.ID)
		
		assert.Equal(t, layout.ID, *updatedComponent.LayoutId)
		assert.Equal(t, 10, updatedComponent.X)
		assert.Equal(t, 20, updatedComponent.Y)
		assert.Equal(t, 300, updatedComponent.Width)
		assert.Equal(t, 200, updatedComponent.Height)
	})

	t.Run("コンポーネントをレイアウトから削除できる", func(t *testing.T) {
		component, _ := createTestLayoutComponent(testUser.ID)
		layout := createTestLayout()
		
		// まずレイアウトに割り当て
		position := map[string]int{"x": 10, "y": 20, "width": 300, "height": 200}
		lcRepo.AssignToLayout(component.ID, layout.ID, testUser.ID, position)
		
		// レイアウトから削除
		err := lcRepo.RemoveFromLayout(component.ID, testUser.ID)
		
		assert.NoError(t, err)
		
		// 削除されたことを確認
		updatedComponent := &model.LayoutComponent{}
		lcRepo.GetLayoutComponentById(updatedComponent, testUser.ID, component.ID)
		
		assert.Nil(t, updatedComponent.LayoutId)
	})

	t.Run("他のユーザーのコンポーネントはレイアウトに割り当てられない", func(t *testing.T) {
		component, _ := createTestLayoutComponent(testUser.ID)
		layout := createTestLayout()
		
		otherUserID := uint(9999) // 存在しないユーザーID
		position := map[string]int{"x": 10, "y": 20, "width": 300, "height": 200}
		
		err := lcRepo.AssignToLayout(component.ID, layout.ID, otherUserID, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("存在しないコンポーネントはレイアウトに割り当てられない", func(t *testing.T) {
		nonExistentID := uint(9999)
		layout := createTestLayout()
		
		position := map[string]int{"x": 10, "y": 20, "width": 300, "height": 200}
		
		err := lcRepo.AssignToLayout(nonExistentID, layout.ID, testUser.ID, position)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("他のユーザーのコンポーネントはレイアウトから削除できない", func(t *testing.T) {
		component, _ := createTestLayoutComponent(testUser.ID)
		layout := createTestLayout()
		
		// まずレイアウトに割り当て
		position := map[string]int{"x": 10, "y": 20, "width": 300, "height": 200}
		lcRepo.AssignToLayout(component.ID, layout.ID, testUser.ID, position)
		
		otherUserID := uint(9999) // 存在しないユーザーID
		
		err := lcRepo.RemoveFromLayout(component.ID, otherUserID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})
}