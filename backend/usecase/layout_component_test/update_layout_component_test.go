package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_UpdateLayoutComponent(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("既存のコンポーネントを更新できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			originalComponent := createTestComponent(t, generateUniqueName())

			// 更新用のデータ
			newName := generateUniqueName() + " Updated"
			updateComponentRequest := model.LayoutComponentRequest{
				Name:    newName,
				Type:    "image",
				Content: "更新されたコンテンツ",
				UserId:  testUserId,
			}

			t.Logf("コンポーネント更新: ID=%d, 新Name=%s", originalComponent.ID, newName)

			// テスト実行
			updatedComponent, err := componentUsecase.UpdateLayoutComponent(updateComponentRequest, testUserId, originalComponent.ID)

			// 検証
			if err != nil {
				t.Errorf("UpdateLayoutComponent() error = %v", err)
			}

			if updatedComponent.ID != originalComponent.ID {
				t.Errorf("UpdateLayoutComponent() ID = %v, want %v", updatedComponent.ID, originalComponent.ID)
			}

			if updatedComponent.Name != newName {
				t.Errorf("UpdateLayoutComponent() name = %v, want %v", updatedComponent.Name, newName)
			}

			// データベースから直接確認
			var savedComponent model.LayoutComponent
			componentDb.First(&savedComponent, originalComponent.ID)

			if savedComponent.Name != newName {
				t.Errorf("UpdateLayoutComponent() saved name = %v, want %v", savedComponent.Name, newName)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合はコンポーネント更新に失敗する", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())

			// 空の名前（バリデーションエラーになるはず）
			invalidComponent := model.LayoutComponentRequest{
				Name:    "",
				Type:    "text",
				Content: "テストコンテンツ",
				UserId:  testUserId,
			}

			t.Logf("無効なコンポーネント更新を試行: ID=%d, Name=空文字", component.ID)

			// テスト実行
			_, err := componentUsecase.UpdateLayoutComponent(invalidComponent, testUserId, component.ID)

			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("空の名前でエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})

		t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistComponentId := uint(9999)

			updateComponent := model.LayoutComponentRequest{
				Name:    generateUniqueName(),
				Type:    "text",
				Content: "テストコンテンツ",
				UserId:  testUserId,
			}

			// テスト実行
			_, err := componentUsecase.UpdateLayoutComponent(updateComponent, testUserId, nonExistComponentId)

			// 検証
			if err == nil {
				t.Error("存在しないコンポーネントIDでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})

		t.Run("他のユーザーのコンポーネントを更新しようとするとエラーを返す", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())

			// 別のユーザーID
			otherUserId := uint(999)

			updateComponent := model.LayoutComponentRequest{
				Name:    generateUniqueName(),
				Type:    "text",
				Content: "テストコンテンツ",
				UserId:  otherUserId,
			}

			// テスト実行
			_, err := componentUsecase.UpdateLayoutComponent(updateComponent, otherUserId, component.ID)

			// 検証
			if err == nil {
				t.Error("他のユーザーのコンポーネント更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}