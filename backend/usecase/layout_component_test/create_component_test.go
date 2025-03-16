package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_CreateLayoutComponent(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいレイアウトコンポーネントを作成できる", func(t *testing.T) {
			// テスト用のコンポーネントデータ
			newComponent := model.LayoutComponent{
				Name:    "New Test Component",
				Type:    "section",
				Content: `{"title": "Test Section", "text": "This is a test section"}`,
				UserId:  testUserId,
			}
			
			// テスト実行
			result, err := layoutComponentUsecase.CreateLayoutComponent(newComponent)
			
			// 検証
			if err != nil {
				t.Errorf("CreateLayoutComponent() error = %v", err)
			}
			
			if result.ID == 0 {
				t.Error("CreateLayoutComponent() returned component with ID = 0")
			}
			
			if result.Name != newComponent.Name {
				t.Errorf("CreateLayoutComponent() returned Name = %s, want %s", result.Name, newComponent.Name)
			}
			
			if result.Type != newComponent.Type {
				t.Errorf("CreateLayoutComponent() returned Type = %s, want %s", result.Type, newComponent.Type)
			}
			
			if result.Content != newComponent.Content {
				t.Errorf("CreateLayoutComponent() returned Content = %s, want %s", result.Content, newComponent.Content)
			}
			
			// データベースに正しく保存されたか確認
			verifyComponentInDB(t, result.ID, newComponent.Name, newComponent.Type, newComponent.Content)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合は作成に失敗する", func(t *testing.T) {
			// 無効なコンポーネント（名前が空）
			invalidComponent := model.LayoutComponent{
				Name:    "", // 空の名前
				Type:    "section",
				Content: `{"title": "Test Section"}`,
				UserId:  testUserId,
			}
			
			// テスト実行
			_, err := layoutComponentUsecase.CreateLayoutComponent(invalidComponent)
			
			// 検証
			if err == nil {
				t.Error("CreateLayoutComponent() did not return error for invalid component")
			}
		})
	})
}