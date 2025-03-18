package layout_component_test

import (
	"go-react-app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_CreateLayoutComponent(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("レイアウトコンポーネントを正常に作成できる", func(t *testing.T) {
		component := testLayoutComponentData
		err := layoutComponentRepo.CreateLayoutComponent(&component)
		
		assert.NoError(t, err)
		assert.NotZero(t, component.ID)
		assert.NotZero(t, component.CreatedAt)
		assert.NotZero(t, component.UpdatedAt)
		assert.Equal(t, testUserData.ID, component.UserId)
	})

	t.Run("ユーザーIDが存在しない場合でもエラーにならない", func(t *testing.T) {
		component := model.LayoutComponent{
			Name:    "Invalid User Component",
			Type:    "text",
			Content: "Test Content",
			UserId:  9999, // 存在しないユーザーID
		}
		
		err := layoutComponentRepo.CreateLayoutComponent(&component)
		// 現在の実装ではエラーが発生しないことを確認
		assert.NoError(t, err)
		assert.NotZero(t, component.ID)
	})
}
