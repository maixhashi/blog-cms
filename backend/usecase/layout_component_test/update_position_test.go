package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_UpdatePosition(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("コンポーネントの位置を更新できる", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestLayoutComponent(t)
			layout := createTestLayout(t)
			
			// まずレイアウトに割り当てる
			layoutId := layout.ID
			component.LayoutId = &layoutId
			component.X = 10
			component.Y = 20
			db.Save(&component)
			
			// 新しい位置情報
			newPosition := map[string]int{
				"x": 30,
				"y": 40,
				"z": 5,
			}
			
			// テスト実行
			err := layoutComponentUsecase.UpdatePosition(testUserId, component.ID, newPosition)
			
			// 検証
			if err != nil {
				t.Errorf("UpdatePosition() error = %v", err)
			}
			
			// データベースで位置情報が更新されたか確認
			var updatedComponent model.LayoutComponent
			db.First(&updatedComponent, component.ID)
			
			if updatedComponent.X != newPosition["x"] {
				t.Errorf("X が正しく更新されていません。取得: %d, 期待: %d", updatedComponent.X, newPosition["x"])
			}
			
			if updatedComponent.Y != newPosition["y"] {
				t.Errorf("Y が正しく更新されていません。取得: %d, 期待: %d", updatedComponent.Y, newPosition["y"])
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントの位置を更新できない", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistentComponentId := uint(9999)
			
			// 位置情報
			position := map[string]int{
				"x": 30,
				"y": 40,
				"z": 5,
			}
			
			// テスト実行
			err := layoutComponentUsecase.UpdatePosition(testUserId, nonExistentComponentId, position)
			
			// 検証
			if err == nil {
				t.Error("UpdatePosition() did not return error for non-existent component")
			}
		})
		
		t.Run("他のユーザーのコンポーネントの位置を更新できない", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestLayoutComponent(t)
			layout := createTestLayout(t)
			
			// まずレイアウトに割り当てる
			layoutId := layout.ID
			component.LayoutId = &layoutId
			component.X = 10
			component.Y = 20
			db.Save(&component)
			
			// 別のユーザーID
			otherUserId := uint(999)
			
			// 新しい位置情報
			newPosition := map[string]int{
				"x": 30,
				"y": 40,
				"z": 5,
			}
			
			// テスト実行
			err := layoutComponentUsecase.UpdatePosition(otherUserId, component.ID, newPosition)
			
			// 検証
			if err == nil {
				t.Error("UpdatePosition() did not return error for different user")
			}
		})
		
		t.Run("レイアウトに割り当てられていないコンポーネントの位置を更新できない", func(t *testing.T) {
			// テスト用のコンポーネントを作成（レイアウトに割り当てない）
			component := createTestLayoutComponent(t)
			
			// 新しい位置情報
			newPosition := map[string]int{
				"x": 30,
				"y": 40,
				"z": 5,
			}
			
			// テスト実行
			err := layoutComponentUsecase.UpdatePosition(testUserId, component.ID, newPosition)
			
			// 検証
			if err == nil {
				t.Error("UpdatePosition() did not return error for component not assigned to layout")
			}
		})
	})
}