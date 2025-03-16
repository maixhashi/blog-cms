package layout_component_test

import (
	"go-react-app/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_UpdateLayoutComponent(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("コンポーネントを正常に更新できる", func(t *testing.T) {
		component, err := createTestLayoutComponent(testUser.ID)
		assert.NoError(t, err)
		
		// 更新用データ
		updatedComponent := *component
		updatedComponent.Name = "Updated Name"
		updatedComponent.Type = "footer"
		updatedComponent.Content = `{"text": "Updated Content"}`
		
		err = lcRepo.UpdateLayoutComponent(&updatedComponent, testUser.ID, component.ID)
		
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updatedComponent.Name)
		assert.Equal(t, "footer", updatedComponent.Type)
		assert.Equal(t, `{"text": "Updated Content"}`, updatedComponent.Content)
		
		// DBから再取得して確認
		foundComponent := &model.LayoutComponent{}
		lcRepo.GetLayoutComponentById(foundComponent, testUser.ID, component.ID)
		assert.Equal(t, "Updated Name", foundComponent.Name)
	})

	t.Run("存在しないコンポーネントは更新できない", func(t *testing.T) {
		nonExistentID := uint(9999)
		component := model.LayoutComponent{
			Name:    "Test",
			Type:    "header",
			Content: `{"text": "Test"}`,
		}
		
		err := lcRepo.UpdateLayoutComponent(&component, testUser.ID, nonExistentID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})

	t.Run("他のユーザーのコンポーネントは更新できない", func(t *testing.T) {
		component, err := createTestLayoutComponent(testUser.ID)
		assert.NoError(t, err)
		
		otherUserID := uint(9999) // 存在しないユーザーID
		
		err = lcRepo.UpdateLayoutComponent(component, otherUserID, component.ID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "layout component does not exist")
	})
}