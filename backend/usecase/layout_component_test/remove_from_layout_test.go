package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_RemoveFromLayout(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("コンポーネントをレイアウトから削除できる", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestLayoutComponent(t)
			layout := createTestLayout(t)
			
			// まずレイアウトに割り当てる
			layoutId := layout.ID
			component.LayoutId = &layoutId
			component.X = 10
			component.Y = 20
			db.Save(&component)
			
			// テスト実行
			err := layoutComponentUsecase.RemoveFromLayout(testUserId, component.ID)
			
			// 検証
			if err != nil {
				t.Errorf("RemoveFromLayout() error = %v", err)
			}
			
			// データベースで関連付けが解除されたか確認
			var updatedComponent model.LayoutComponent
			db.First(&updatedComponent, component.ID)
			
			if updatedComponent.LayoutId != nil {
				t.Errorf("コンポーネントがレイアウトから削除されていません。LayoutId = %v, want nil", updatedComponent.LayoutId)
			}
			
			// 位置情報がリセットされたか確認
			if updatedComponent.X != 0 || updatedComponent.Y != 0 {
				t.Errorf("位置情報がリセットされていません。Position = (%d, %d), want (0, 0)",
					updatedComponent.X, updatedComponent.Y)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントをレイアウトから削除できない", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistentComponentId := uint(9999)
			
			// テスト実行
			err := layoutComponentUsecase.RemoveFromLayout(testUserId, nonExistentComponentId)
			
			// 検証
			if err == nil {
				t.Error("RemoveFromLayout() did not return error for non-existent component")
			}
		})
		
		t.Run("他のユーザーのコンポーネントをレイアウトから削除できない", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestLayoutComponent(t)
			layout := createTestLayout(t)
			
			// まずレイアウトに割り当てる
			layoutId := layout.ID
			component.LayoutId = &layoutId
			db.Save(&component)
			
			// 別のユーザーID
			otherUserId := uint(999)
			
			// テスト実行
			err := layoutComponentUsecase.RemoveFromLayout(otherUserId, component.ID)
			
			// 検証
			if err == nil {
				t.Error("RemoveFromLayout() did not return error for different user")
			}
		})
	})
}
