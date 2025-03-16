package layout_component_test

import (
	"go-react-app/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLayoutComponentRepository_GetLayoutComponentById(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("存在するコンポーネントを取得できる", func(t *testing.T) {
		component, err := createTestLayoutComponent(testUser.ID)
		assert.NoError(t, err)
		
		foundComponent := &model.LayoutComponent{}
		err = lcRepo.GetLayoutComponentById(foundComponent, testUser.ID, component.ID)
		
		assert.NoError(t, err)
		assert.Equal(t, component.ID, foundComponent.ID)
		assert.Equal(t, component.Name, foundComponent.Name)
		assert.Equal(t, component.Type, foundComponent.Type)
		assert.Equal(t, component.Content, foundComponent.Content)
		assert.Equal(t, testUser.ID, foundComponent.UserId)
	})

	t.Run("存在しないコンポーネントを取得するとエラーになる", func(t *testing.T) {
		nonExistentID := uint(9999)
		foundComponent := &model.LayoutComponent{}
		
		err := lcRepo.GetLayoutComponentById(foundComponent, testUser.ID, nonExistentID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})

	t.Run("他のユーザーのコンポーネントは取得できない", func(t *testing.T) {
		component, err := createTestLayoutComponent(testUser.ID)
		assert.NoError(t, err)
		
		otherUserID := uint(9999) // 存在しないユーザーID
		foundComponent := &model.LayoutComponent{}
		
		err = lcRepo.GetLayoutComponentById(foundComponent, otherUserID, component.ID)
		
		assert.Error(t, err)
	})
}

func TestLayoutComponentRepository_GetAllLayoutComponents(t *testing.T) {
	setupLayoutComponentRepositoryTest()

	t.Run("ユーザーの全コンポーネントを取得できる", func(t *testing.T) {
		// テスト用のコンポーネントを複数作成
		component1, _ := createTestLayoutComponent(testUser.ID)
		component2 := generateUniqueTestComponent("Second Component", "footer", testUser.ID)
		lcRepo.CreateLayoutComponent(&component2)
		
		var components []model.LayoutComponent
		err := lcRepo.GetAllLayoutComponents(&components, testUser.ID)
		
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(components), 2)
		
		// 最新のコンポーネントが最初に来ることを確認（降順）
		assert.Equal(t, component2.ID, components[0].ID)
		assert.Equal(t, component1.ID, components[1].ID)
	})

	t.Run("コンポーネントがない場合は空の配列を返す", func(t *testing.T) {
		newUser := model.User{
			Email:    "no-components@example.com",
			Password: "password123",
		}
		userRepo.CreateUser(&newUser)
		
		var components []model.LayoutComponent
		err := lcRepo.GetAllLayoutComponents(&components, newUser.ID)
		
		assert.NoError(t, err)
		assert.Empty(t, components)
	})
}