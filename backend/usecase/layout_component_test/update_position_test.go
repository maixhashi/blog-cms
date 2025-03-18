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
			component := createTestComponent(t, generateUniqueName())
			layoutId := createTestLayout(t)
			
			// まずレイアウトに割り当て
			assignRequest := model.AssignLayoutRequest{
				LayoutId: layoutId,
				Position: model.PositionRequest{X: 10, Y: 20},
			}
			
			err := componentUsecase.AssignToLayout(testUserId, component.ID, assignRequest)
			if err != nil {
				t.Fatalf("テスト準備中にエラーが発生しました: %v", err)
			}
			
			// 新しい位置情報
			newPosition := model.PositionRequest{
				X: 30,
				Y: 40,
				Width: 300,
				Height: 250,
			}
			
			// テスト実行
			err = componentUsecase.UpdatePosition(testUserId, component.ID, newPosition)
			
			// 検証
			if err != nil {
				t.Errorf("UpdatePosition() error = %v", err)
			}
			
			// データベースから直接確認
			var updatedComponent model.LayoutComponent
			componentDb.First(&updatedComponent, component.ID)
			
			if updatedComponent.X != 30 || updatedComponent.Y != 40 {
				t.Errorf("UpdatePosition() position not updated correctly: got (%d,%d), want (30,40)", 
					updatedComponent.X, updatedComponent.Y)
			}
			
			if updatedComponent.Width != 300 || updatedComponent.Height != 250 {
				t.Errorf("UpdatePosition() size not updated correctly: got (%d,%d), want (300,250)", 
					updatedComponent.Width, updatedComponent.Height)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistComponentId := uint(9999)
			
			position := model.PositionRequest{X: 30, Y: 40}
			
			// テスト実行
			err := componentUsecase.UpdatePosition(testUserId, nonExistComponentId, position)
			
			// 検証
			if err == nil {
				t.Error("存在しないコンポーネントIDでエラーが返されませんでした")
			}
		})
		
		t.Run("レイアウトに割り当てられていないコンポーネントの位置更新はエラーを返す", func(t *testing.T) {
			// レイアウトに割り当てていないコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())
			
			position := model.PositionRequest{X: 30, Y: 40}
			
			// テスト実行
			err := componentUsecase.UpdatePosition(testUserId, component.ID, position)
			
			// 検証
			if err == nil {
				t.Error("レイアウトに割り当てられていないコンポーネントの位置更新でエラーが返されませんでした")
			}
		})
		
		t.Run("他のユーザーのコンポーネントの位置を更新しようとするとエラーを返す", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestComponent(t, generateUniqueName())
			layoutId := createTestLayout(t)
			
			// まずレイアウトに割り当て
			assignRequest := model.AssignLayoutRequest{
				LayoutId: layoutId,
				Position: model.PositionRequest{X: 10, Y: 20},
			}
			
			err := componentUsecase.AssignToLayout(testUserId, component.ID, assignRequest)
			if err != nil {
				t.Fatalf("テスト準備中にエラーが発生しました: %v", err)
			}
			
			// 別のユーザーID
			otherUserId := uint(999)
			
			position := model.PositionRequest{X: 30, Y: 40}
			
			// テスト実行
			err = componentUsecase.UpdatePosition(otherUserId, component.ID, position)
			
			// 検証
			if err == nil {
				t.Error("他のユーザーのコンポーネント位置更新でエラーが返されませんでした")
			}
			
			// 元のコンポーネントが変更されていないことを確認
			var savedComponent model.LayoutComponent
			componentDb.First(&savedComponent, component.ID)
			
			if savedComponent.X != 10 || savedComponent.Y != 20 {
				t.Errorf("他のユーザーによる位置更新試行で元のコンポーネントの位置が変更されました: got (%d,%d), want (10,20)",
					savedComponent.X, savedComponent.Y)
			}
		})
	})
}