package layout_component_test

import (
	"testing"
)

func TestLayoutComponentUsecase_GetAllLayoutComponents(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのコンポーネント一覧を取得できる", func(t *testing.T) {
			// テスト用のコンポーネントを複数作成
			component1 := createTestComponent(t, generateUniqueName())
			component2 := createTestComponent(t, generateUniqueName())
			component3 := createTestComponent(t, generateUniqueName())

			// テスト実行
			components, err := componentUsecase.GetAllLayoutComponents(testUserId)

			// 検証
			if err != nil {
				t.Errorf("GetAllLayoutComponents() error = %v", err)
			}

			// 少なくとも作成した3つのコンポーネントが含まれているか確認
			if len(components) < 3 {
				t.Errorf("GetAllLayoutComponents() returned %d components, want at least 3", len(components))
			}

			// 作成したコンポーネントが含まれているか確認
			foundComponent1 := false
			foundComponent2 := false
			foundComponent3 := false

			for _, component := range components {
				if component.ID == component1.ID {
					foundComponent1 = true
				}
				if component.ID == component2.ID {
					foundComponent2 = true
				}
				if component.ID == component3.ID {
					foundComponent3 = true
				}
			}

			if !foundComponent1 || !foundComponent2 || !foundComponent3 {
				t.Errorf("GetAllLayoutComponents() did not return all created components")
			}
		})

		t.Run("コンポーネントが存在しない場合は空の配列を返す", func(t *testing.T) {
			// 別のユーザーIDを使用
			nonExistUserId := uint(999)

			// テスト実行
			components, err := componentUsecase.GetAllLayoutComponents(nonExistUserId)

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
		t.Run("指定したIDのコンポーネントを取得できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			expectedComponent := createTestComponent(t, generateUniqueName())

			// テスト実行
			component, err := componentUsecase.GetLayoutComponentById(testUserId, expectedComponent.ID)

			// 検証
			if err != nil {
				t.Errorf("GetLayoutComponentById() error = %v", err)
			}

			if component.ID != expectedComponent.ID {
				t.Errorf("GetLayoutComponentById() = %v, want %v", component.ID, expectedComponent.ID)
			}

			if component.Name != expectedComponent.Name {
				t.Errorf("GetLayoutComponentById() name = %v, want %v", component.Name, expectedComponent.Name)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistComponentId := uint(9999)

			// テスト実行
			_, err := componentUsecase.GetLayoutComponentById(testUserId, nonExistComponentId)

			// 検証
			if err == nil {
				t.Error("GetLayoutComponentById() error = nil, want error for non-existent component")
			}
		})

		t.Run("他のユーザーのコンポーネントにアクセスするとエラーを返す", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())

			// 別のユーザーIDを使用
			otherUserId := uint(999)

			// テスト実行
			_, err := componentUsecase.GetLayoutComponentById(otherUserId, component.ID)

			// 検証
			if err == nil {
				t.Error("GetLayoutComponentById() error = nil, want error for accessing other user's component")
			}
		})
	})
}
