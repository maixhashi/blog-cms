package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_UpdateLayoutComponent(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("既存のレイアウトコンポーネントを更新できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// 更新用データ
			updatedComponent := model.LayoutComponent{
				Name:    "Updated Component",
				Type:    "banner",
				Content: `{"title": "Updated Banner", "image": "banner.jpg"}`,
				UserId:  testUserId,
			}
			
			// テスト実行
			result, err := layoutComponentUsecase.UpdateLayoutComponent(updatedComponent, testUserId, component.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateLayoutComponent() error = %v", err)
			}
			
			if result.ID != component.ID {
				t.Errorf("UpdateLayoutComponent() returned ID = %d, want %d", result.ID, component.ID)
			}
			
			if result.Name != updatedComponent.Name {
				t.Errorf("UpdateLayoutComponent() returned Name = %s, want %s", result.Name, updatedComponent.Name)
			}
			
			if result.Type != updatedComponent.Type {
				t.Errorf("UpdateLayoutComponent() returned Type = %s, want %s", result.Type, updatedComponent.Type)
			}
			
			if result.Content != updatedComponent.Content {
				t.Errorf("UpdateLayoutComponent() returned Content = %s, want %s", result.Content, updatedComponent.Content)
			}
			
			// データベースに正しく保存されたか確認
			verifyComponentInDB(t, component.ID, updatedComponent.Name, updatedComponent.Type, updatedComponent.Content)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントの更新に失敗する", func(t *testing.T) {
			// 存在しないID
			nonExistentId := uint(9999)
			
			// 更新用データ
			updatedComponent := model.LayoutComponent{
				Name:    "Updated Component",
				Type:    "banner",
				Content: `{"title": "Updated Banner"}`,
				UserId:  testUserId,
			}
			
			// テスト実行
			_, err := layoutComponentUsecase.UpdateLayoutComponent(updatedComponent, testUserId, nonExistentId)
			
			// 検証
			if err == nil {
				t.Error("UpdateLayoutComponent() did not return error for non-existent ID")
			}
		})
		
		t.Run("バリデーションエラーが発生する場合は更新に失敗する", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// 無効なコンポーネント（名前が空）
			invalidComponent := model.LayoutComponent{
				Name:    "", // 空の名前
				Type:    "section",
				Content: `{"title": "Test Section"}`,
				UserId:  testUserId,
			}
			
			// テスト実行
			_, err := layoutComponentUsecase.UpdateLayoutComponent(invalidComponent, testUserId, component.ID)
			
			// 検証
			if err == nil {
				t.Error("UpdateLayoutComponent() did not return error for invalid component")
			}
		})
		
		t.Run("他のユーザーのコンポーネントを更新できない", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// 更新用データ
			updatedComponent := model.LayoutComponent{
				Name:    "Updated Component",
				Type:    "banner",
				Content: `{"title": "Updated Banner", "image": "banner.jpg"}`,
				UserId:  testUserId,
			}
			
			// 別のユーザーID
			otherUserId := uint(999)
			
			// テスト実行
			_, err := layoutComponentUsecase.UpdateLayoutComponent(updatedComponent, otherUserId, component.ID)
			
			// 検証
			if err == nil {
				t.Error("UpdateLayoutComponent() did not return error for different user")
			}
			
			// データベースで更新されていないことを確認
			verifyComponentInDB(t, component.ID, component.Name, component.Type, component.Content)
		})
	})
}