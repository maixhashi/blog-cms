package layout_component_test

import (
	"go-react-app/model" // modelパッケージをインポート
	"testing"
)

func TestLayoutComponentUsecase_DeleteLayoutComponent(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("レイアウトコンポーネントを削除できる", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// テスト実行
			err := layoutComponentUsecase.DeleteLayoutComponent(testUserId, component.ID)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteLayoutComponent() error = %v", err)
			}
			
			// データベースから削除されたことを確認
			var count int64
			db.Model(&model.LayoutComponent{}).Where("id = ?", component.ID).Count(&count)
			
			if count != 0 {
				t.Errorf("コンポーネントがデータベースから削除されていません。count = %d", count)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないコンポーネントの削除に失敗する", func(t *testing.T) {
			// 存在しないID
			nonExistentId := uint(9999)
			
			// テスト実行
			err := layoutComponentUsecase.DeleteLayoutComponent(testUserId, nonExistentId)
			
			// 検証
			if err == nil {
				t.Error("DeleteLayoutComponent() did not return error for non-existent ID")
			}
		})
		
		t.Run("他のユーザーのコンポーネントを削除できない", func(t *testing.T) {
			// テスト用のコンポーネントを作成
			component := createTestLayoutComponent(t)
			
			// 別のユーザーIDで削除を試みる
			otherUserId := uint(999)
			
			// テスト実行
			err := layoutComponentUsecase.DeleteLayoutComponent(otherUserId, component.ID)
			
			// 検証
			if err == nil {
				t.Error("DeleteLayoutComponent() did not return error for different user")
			}
			
			// データベースから削除されていないことを確認
			var count int64
			db.Model(&model.LayoutComponent{}).Where("id = ?", component.ID).Count(&count)
			
			if count != 1 {
				t.Errorf("コンポーネントがデータベースから誤って削除されました。count = %d", count)
			}
		})
	})
}