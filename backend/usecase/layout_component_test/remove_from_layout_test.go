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
			
			// テスト実行
			err = componentUsecase.RemoveFromLayout(testUserId, component.ID)
			
			// 検証
			if err != nil {
				t.Errorf("RemoveFromLayout() error = %v", err)
			}
			
			// データベースから直接確認
			var updatedComponent model.LayoutComponent
			componentDb.First(&updatedComponent, component.ID)
			
			if updatedComponent.LayoutId != nil {
				t.Errorf("RemoveFromLayout() failed: component still assigned to layout %d", *updatedComponent.LayoutId)
			}
			
			// 位置情報がリセットされているか確認
			if updatedComponent.X != 0 || updatedComponent.Y != 0 {
				t.Errorf("RemoveFromLayout() position not reset: got (%d,%d), want (0,0)", 
					updatedComponent.X, updatedComponent.Y)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistComponentId := uint(9999)
			
			// テスト実行
			err := componentUsecase.RemoveFromLayout(testUserId, nonExistComponentId)
			
			// 検証
			if err == nil {
				t.Error("存在しないコンポーネントIDでエラーが返されませんでした")
			}
		})
		
		t.Run("他のユーザーのコンポーネントを操作しようとするとエラーを返す", func(t *testing.T) {
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
			
			// テスト実行
			err = componentUsecase.RemoveFromLayout(otherUserId, component.ID)
			
			// 検証
			if err == nil {
				t.Error("他のユーザーのコンポーネント操作でエラーが返されませんでした")
			}
			
			// 元のコンポーネントが変更されていないことを確認
			var savedComponent model.LayoutComponent
			componentDb.First(&savedComponent, component.ID)
			
			if savedComponent.LayoutId == nil || *savedComponent.LayoutId != layoutId {
				t.Errorf("他のユーザーによる操作試行で元のコンポーネントの割り当てが変更されました")
			}
		})
	})
}
