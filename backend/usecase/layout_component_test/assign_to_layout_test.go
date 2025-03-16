package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_AssignToLayout(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("コンポーネントをレイアウトに割り当てることができる", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestLayoutComponent(t)
			layout := createTestLayout(t)
			
			// 位置情報
			position := map[string]int{
				"x": 10,
				"y": 20,
				"z": 0,
			}
			
			// テスト実行
			err := layoutComponentUsecase.AssignToLayout(testUserId, component.ID, layout.ID, position)
			
			// 検証
			if err != nil {
				t.Errorf("AssignToLayout() error = %v", err)
				return // エラーが発生した場合は早期リターン
			}
			
			// データベースで関連付けが正しく設定されたか確認
			var updatedComponent model.LayoutComponent
			result := db.First(&updatedComponent, component.ID)
			if result.Error != nil {
				t.Errorf("コンポーネントの取得に失敗しました: %v", result.Error)
				return
			}
			
			if updatedComponent.LayoutId == nil {
				t.Errorf("コンポーネントがレイアウトに割り当てられていません。LayoutId = nil, want %d", layout.ID)
				return
			}
			
			if *updatedComponent.LayoutId != layout.ID {
				t.Errorf("コンポーネントが正しいレイアウトに割り当てられていません。LayoutId = %d, want %d", 
					*updatedComponent.LayoutId, layout.ID)
				return
			}
			
			// 位置情報が正しく設定されたか確認
			if updatedComponent.X != position["x"] {
				t.Errorf("X が正しく設定されていません。取得: %d, 期待: %d", updatedComponent.X, position["x"])
			}
			
			if updatedComponent.Y != position["y"] {
				t.Errorf("Y が正しく設定されていません。取得: %d, 期待: %d", updatedComponent.Y, position["y"])
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントをレイアウトに割り当てることができない", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistentComponentId := uint(9999)
			
			// テスト用のレイアウトを作成
			layout := createTestLayout(t)
			
			// 位置情報
			position := map[string]int{
				"x": 10,
				"y": 20,
				"z": 0,
			}
			
			// テスト実行
			err := layoutComponentUsecase.AssignToLayout(testUserId, nonExistentComponentId, layout.ID, position)
			
			// 検証
			if err == nil {
				t.Error("AssignToLayout() did not return error for non-existent component")
			}
		})
		
		t.Run("存在しないレイアウトにコンポーネントを割り当てることができない", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// 存在しないレイアウトID（非常に大きな値を使用）
			nonExistentLayoutId := uint(9999999)
			
			// 位置情報
			position := map[string]int{
				"x": 10,
				"y": 20,
				"z": 0,
			}
			
			// テスト実行
			err := layoutComponentUsecase.AssignToLayout(testUserId, component.ID, nonExistentLayoutId, position)
			
			// 検証
			if err == nil {
				t.Error("AssignToLayout() did not return error for non-existent layout")
			}
		})
		
		t.Run("他のユーザーのコンポーネントをレイアウトに割り当てることができない", func(t *testing.T) {
			// テスト用のコンポーネントとレイアウトを作成
			component := createTestLayoutComponent(t)
			layout := createTestLayout(t)
			
			// 別のユーザーID
			otherUserId := uint(999)
			
			// 位置情報
			position := map[string]int{
				"x": 10,
				"y": 20,
				"z": 0,
			}
			
			// テスト実行
			err := layoutComponentUsecase.AssignToLayout(otherUserId, component.ID, layout.ID, position)
			
			// 検証
			if err == nil {
				t.Error("AssignToLayout() did not return error for different user")
			}
		})
	})
}
// 複数のコンポーネントを同時に操作するテスト
func TestLayoutComponentUsecase_MultipleComponentAssignments(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("複数のコンポーネントをレイアウトに割り当てることができる", func(t *testing.T) {
		// テスト用のレイアウトを作成
		layout := createTestLayout(t)
		
		// 複数のコンポーネントを作成
		component1 := createTestLayoutComponent(t)
		component1.Name = "First Component"
		db.Save(&component1)
		
		component2 := createTestLayoutComponent(t)
		component2.Name = "Second Component"
		db.Save(&component2)
		
		// 位置情報
		position1 := map[string]int{
			"x": 10,
			"y": 20,
			"z": 0,
		}
		
		position2 := map[string]int{
			"x": 100,
			"y": 200,
			"z": 1,
		}
		
		// コンポーネント1をレイアウトに割り当て
		err1 := layoutComponentUsecase.AssignToLayout(testUserId, component1.ID, layout.ID, position1)
		
		// コンポーネント2をレイアウトに割り当て
		err2 := layoutComponentUsecase.AssignToLayout(testUserId, component2.ID, layout.ID, position2)
		
		// 検証
		if err1 != nil || err2 != nil {
			t.Errorf("AssignToLayout() errors: %v, %v", err1, err2)
		}
		
		// データベースで関連付けが正しく設定されたか確認
		var updatedComponent1 model.LayoutComponent
		var updatedComponent2 model.LayoutComponent
		db.First(&updatedComponent1, component1.ID)
		db.First(&updatedComponent2, component2.ID)
		
		// コンポーネント1の検証
		if updatedComponent1.LayoutId == nil || *updatedComponent1.LayoutId != layout.ID {
			t.Errorf("コンポーネント1がレイアウトに正しく割り当てられていません。LayoutId = %v, want %d", 
				updatedComponent1.LayoutId, layout.ID)
		}
		
		if updatedComponent1.X != position1["x"] || 
		   updatedComponent1.Y != position1["y"] {
			t.Errorf("コンポーネント1の位置情報が正しくありません。Position = (%d, %d), want (%d, %d)",
				updatedComponent1.X, updatedComponent1.Y,
				position1["x"], position1["y"])
		}
		
		// コンポーネント2の検証
		if updatedComponent2.LayoutId == nil || *updatedComponent2.LayoutId != layout.ID {
			t.Errorf("コンポーネント2がレイアウトに正しく割り当てられていません。LayoutId = %v, want %d", 
				updatedComponent2.LayoutId, layout.ID)
		}
		
		if updatedComponent2.X != position2["x"] || 
		   updatedComponent2.Y != position2["y"] {
			t.Errorf("コンポーネント2の位置情報が正しくありません。Position = (%d, %d), want (%d, %d)",
				updatedComponent2.X, updatedComponent2.Y,
				position2["x"], position2["y"])
		}
	})
}