package layout_component_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_GetAllLayoutComponents(t *testing.T) {
	setupLayoutComponentRepositoryTest() // この関数がデータベースをクリーンアップしていることを確認

	t.Run("ユーザーのすべてのレイアウトコンポーネントを取得できる", func(t *testing.T) {
		// テスト用のコンポーネントを複数作成
		component1 := generateUniqueTestLayoutComponent("Component 1", "text", testUserData.ID)
		component2 := generateUniqueTestLayoutComponent("Component 2", "image", testUserData.ID)
		
		err1 := layoutComponentRepo.CreateLayoutComponent(&component1)
		err2 := layoutComponentRepo.CreateLayoutComponent(&component2)
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		
		components, err := layoutComponentRepo.GetAllLayoutComponents(testUserData.ID)
		
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(components), 2)
	})

	t.Run("存在しないユーザーIDの場合は空の結果が返される", func(t *testing.T) {
		// 非常に大きな値を使用して、他のテストと衝突しないようにする
		nonExistentUserId := uint(99999)
		
		components, err := layoutComponentRepo.GetAllLayoutComponents(nonExistentUserId)
		
		assert.NoError(t, err)
		assert.Empty(t, components)
	})
}

func TestLayoutComponentRepository_GetLayoutComponentById(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("IDで特定のレイアウトコンポーネントを取得できる", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		foundComponent, err := layoutComponentRepo.GetLayoutComponentById(testUserData.ID, component.ID)
		
		assert.NoError(t, err)
		assert.Equal(t, component.ID, foundComponent.ID)
		assert.Equal(t, component.Name, foundComponent.Name)
		assert.Equal(t, component.Type, foundComponent.Type)
		assert.Equal(t, component.UserId, foundComponent.UserId)
	})

	t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
		_, err := layoutComponentRepo.GetLayoutComponentById(testUserData.ID, 9999)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("ユーザーIDが一致しない場合はエラーを返す", func(t *testing.T) {
		component, err := createTestLayoutComponent()
		assert.NoError(t, err)
		
		_, err = layoutComponentRepo.GetLayoutComponentById(9999, component.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})
}