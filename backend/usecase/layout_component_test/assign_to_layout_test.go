package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

// レイアウト作成のヘルパー関数
func createTestLayout(t *testing.T) uint {
	// レイアウトテーブルに直接挿入
	layout := model.Layout{
		Title:  "Test Layout for Component",
		UserId: testUserId,
	}
	
	result := componentDb.Create(&layout)
	if result.Error != nil {
		t.Fatalf("テストレイアウトの作成に失敗しました: %v", result.Error)
	}
	
	return layout.ID
}

func TestLayoutComponentUsecase_AssignToLayout(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("コンポーネントをレイアウトに割り当てることができる", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestComponent(t, generateUniqueName())
			layoutId := createTestLayout(t)
			
			// 割り当て用のリクエスト
			assignRequest := model.AssignLayoutRequest{
				LayoutId: layoutId,
				Position: model.PositionRequest{
					X: 10,
					Y: 20,
					Width: 200,
					Height: 150,
				},
			}
			
			// テスト実行
			err := componentUsecase.AssignToLayout(testUserId, component.ID, assignRequest)
			
			// 検証
			if err != nil {
				t.Errorf("AssignToLayout() error = %v", err)
			}
			
			// データベースから直接確認
			var updatedComponent model.LayoutComponent
			componentDb.First(&updatedComponent, component.ID)
			
			if updatedComponent.LayoutId == nil || *updatedComponent.LayoutId != layoutId {
				t.Errorf("AssignToLayout() failed: component not assigned to layout %d", layoutId)
			}
			
			if updatedComponent.X != 10 || updatedComponent.Y != 20 {
				t.Errorf("AssignToLayout() position not set correctly: got (%d,%d), want (10,20)", 
					updatedComponent.X, updatedComponent.Y)
			}
			
			if updatedComponent.Width != 200 || updatedComponent.Height != 150 {
				t.Errorf("AssignToLayout() size not set correctly: got (%d,%d), want (200,150)", 
					updatedComponent.Width, updatedComponent.Height)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistComponentId := uint(9999)
			layoutId := createTestLayout(t)
			
			assignRequest := model.AssignLayoutRequest{
				LayoutId: layoutId,
				Position: model.PositionRequest{X: 10, Y: 20},
			}
			
			// テスト実行
			err := componentUsecase.AssignToLayout(testUserId, nonExistComponentId, assignRequest)
			
			// 検証
			if err == nil {
				t.Error("存在しないコンポーネントIDでエラーが返されませんでした")
			}
		})
		
		t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())
			// 存在しないレイアウトID
			nonExistLayoutId := uint(9999)
			
			assignRequest := model.AssignLayoutRequest{
				LayoutId: nonExistLayoutId,
				Position: model.PositionRequest{X: 10, Y: 20},
			}
			
			// テスト実行
			err := componentUsecase.AssignToLayout(testUserId, component.ID, assignRequest)
			
			// 検証
			if err == nil {
				t.Error("存在しないレイアウトIDでエラーが返されませんでした")
			}
		})
		
		t.Run("他のユーザーのコンポーネントを割り当てようとするとエラーを返す", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestComponent(t, generateUniqueName())
			layoutId := createTestLayout(t)
			
			// 別のユーザーID
			otherUserId := uint(999)
			
			assignRequest := model.AssignLayoutRequest{
				LayoutId: layoutId,
				Position: model.PositionRequest{X: 10, Y: 20},
			}
			
			// テスト実行
			err := componentUsecase.AssignToLayout(otherUserId, component.ID, assignRequest)
			
			// 検証
			if err == nil {
				t.Error("他のユーザーのコンポーネント割り当てでエラーが返されませんでした")
			}
		})
	})
}
