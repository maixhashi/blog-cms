package layout_component_test

import (
	"go-react-app/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_CreateLayoutComponent(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("コンポーネントを正常に作成できる", func(t *testing.T) {
		component := testLayoutComponent
		component.UserId = testUser.ID
		
		err := lcRepo.CreateLayoutComponent(&component)
		
		assert.NoError(t, err)
		assert.NotZero(t, component.ID)
		assert.NotZero(t, component.CreatedAt)
		assert.NotZero(t, component.UpdatedAt)
		assert.Equal(t, testUser.ID, component.UserId)
	})

	// テストケースの名前と期待値を変更
	t.Run("リポジトリレベルではバリデーションが行われない", func(t *testing.T) {
		invalidComponent := model.LayoutComponent{
			// Name が欠けている
			Type:    "header",
			Content: `{"text": "Test"}`,
			UserId:  testUser.ID,
		}
		
		err := lcRepo.CreateLayoutComponent(&invalidComponent)
		
		// エラーを期待しない
		assert.NoError(t, err)
	})
}