package layout_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutRepository_GetAllLayouts(t *testing.T) {
	setupLayoutRepositoryTest()

	t.Run("ユーザーのすべてのレイアウトを取得できる", func(t *testing.T) {
		// テスト用のレイアウトを複数作成
		layout1 := generateUniqueTestLayout("Layout 1", testUserData.ID)
		layout2 := generateUniqueTestLayout("Layout 2", testUserData.ID)
		
		err1 := layoutRepo.CreateLayout(&layout1)
		err2 := layoutRepo.CreateLayout(&layout2)
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		
		var layouts []model.Layout
		err := layoutRepo.GetAllLayouts(&layouts, testUserData.ID)
		
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(layouts), 2)
	})

	t.Run("存在しないユーザーIDの場合でも結果が返される", func(t *testing.T) {
		// まず存在しないユーザーIDでレイアウトを作成
		nonExistentUserId := uint(9999)
		layout := generateUniqueTestLayout("Invalid User Layout", nonExistentUserId)
		err := layoutRepo.CreateLayout(&layout)
		assert.NoError(t, err)
		
		var layouts []model.Layout
		err = layoutRepo.GetAllLayouts(&layouts, nonExistentUserId)
		
		assert.NoError(t, err)
		// 現在の実装では空ではなく結果が返されることを確認
		assert.NotEmpty(t, layouts)
		assert.Equal(t, nonExistentUserId, layouts[0].UserId)
	})
}

func TestLayoutRepository_GetLayoutById(t *testing.T) {
	setupLayoutRepositoryTest()

	t.Run("IDで特定のレイアウトを取得できる", func(t *testing.T) {
		layout, err := createTestLayout()
		assert.NoError(t, err)
		
		var foundLayout model.Layout
		err = layoutRepo.GetLayoutById(&foundLayout, testUserData.ID, layout.ID)
		
		assert.NoError(t, err)
		assert.Equal(t, layout.ID, foundLayout.ID)
		assert.Equal(t, layout.Title, foundLayout.Title)
		assert.Equal(t, layout.UserId, foundLayout.UserId)
	})

	t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
		var layout model.Layout
		err := layoutRepo.GetLayoutById(&layout, testUserData.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		layout, err := createTestLayout()
		assert.NoError(t, err)
		
		var foundLayout model.Layout
		err = layoutRepo.GetLayoutById(&foundLayout, 9999, layout.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})
}