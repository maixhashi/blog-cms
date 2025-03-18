package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_DeleteLayoutComponent(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("既存のコンポーネントを削除できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())

			t.Logf("コンポーネント削除: ID=%d", component.ID)

			// テスト実行
			err := componentUsecase.DeleteLayoutComponent(testUserId, component.ID)

			// 検証
			if err != nil {
				t.Errorf("DeleteLayoutComponent() error = %v", err)
			}

			// データベースから直接確認
			var deletedComponent model.LayoutComponent
			result := componentDb.First(&deletedComponent, component.ID)

			// レコードが見つからないはず
			if result.Error == nil {
				t.Errorf("DeleteLayoutComponent() failed: component with ID %d still exists", component.ID)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないコンポーネントID
			nonExistComponentId := uint(9999)

			// テスト実行
			err := componentUsecase.DeleteLayoutComponent(testUserId, nonExistComponentId)

			// 検証
			if err == nil {
				t.Error("存在しないコンポーネントIDでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})

		t.Run("他のユーザーのコンポーネントを削除しようとするとエラーを返す", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestComponent(t, generateUniqueName())

			// 別のユーザーID
			otherUserId := uint(999)

			// テスト実行
			err := componentUsecase.DeleteLayoutComponent(otherUserId, component.ID)

			// 検証
			if err == nil {
				t.Error("他のユーザーのコンポーネント削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}

			// 元のコンポーネントが削除されていないことを確認
			var savedComponent model.LayoutComponent
			result := componentDb.First(&savedComponent, component.ID)

			if result.Error != nil {
				t.Errorf("他のユーザーによる削除試行で元のコンポーネントが削除されました: %v", result.Error)
			}
		})
	})
}
