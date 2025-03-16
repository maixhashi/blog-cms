package layout_component_test

import (
	"testing"
)

func TestLayoutComponentUsecase_GetAllLayoutComponents(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのすべてのレイアウトコンポーネントを取得できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component1 := createTestLayoutComponent(t)
			component2 := createTestLayoutComponent(t)
			component2.Name = "Second Component"
			component2.Type = "footer"
			db.Save(&component2)
			
			// テスト実行
			components, err := layoutComponentUsecase.GetAllLayoutComponents(testUserId)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllLayoutComponents() error = %v", err)
			}
			
			if len(components) != 2 {
				t.Errorf("GetAllLayoutComponents() returned %d components, want 2", len(components))
			}
			
			// 各コンポーネントの検証
			foundComponent1 := false
			foundComponent2 := false
			
			for _, comp := range components {
				if comp.ID == component1.ID {
					if validateLayoutComponentResponse(t, comp, component1) {
						foundComponent1 = true
					}
				} else if comp.ID == component2.ID {
					if validateLayoutComponentResponse(t, comp, component2) {
						foundComponent2 = true
					}
				}
			}
			
			if !foundComponent1 {
				t.Error("最初のコンポーネントが結果に含まれていません")
			}
			
			if !foundComponent2 {
				t.Error("2番目のコンポーネントが結果に含まれていません")
			}
		})
		
		t.Run("コンポーネントが存在しない場合は空の配列を返す", func(t *testing.T) {
			// データベースをクリーンアップ
			db.Exec("DELETE FROM layout_components WHERE user_id = ?", testUserId)
			
			// テスト実行
			components, err := layoutComponentUsecase.GetAllLayoutComponents(testUserId)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllLayoutComponents() error = %v", err)
			}
			
			if len(components) != 0 {
				t.Errorf("GetAllLayoutComponents() returned %d components, want 0", len(components))
			}
		})
	})
}

func TestLayoutComponentUsecase_GetLayoutComponentById(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("指定したIDのレイアウトコンポーネントを取得できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// テスト実行
			result, err := layoutComponentUsecase.GetLayoutComponentById(testUserId, component.ID)
			
			// 検証
			if err != nil {
				t.Errorf("GetLayoutComponentById() error = %v", err)
			}
			
			validateLayoutComponentResponse(t, result, component)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			// 存在しないID
			nonExistentId := uint(9999)
			
			// テスト実行
			_, err := layoutComponentUsecase.GetLayoutComponentById(testUserId, nonExistentId)
			
			// 検証
			if err == nil {
				t.Error("GetLayoutComponentById() did not return error for non-existent ID")
			}
		})
		
		t.Run("他のユーザーのコンポーネントにアクセスできない", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// 別のユーザーIDでアクセス
			otherUserId := uint(999)
			
			// テスト実行
			_, err := layoutComponentUsecase.GetLayoutComponentById(otherUserId, component.ID)
			
			// 検証
			if err == nil {
				t.Error("GetLayoutComponentById() did not return error for different user")
			}
		})
	})
}
