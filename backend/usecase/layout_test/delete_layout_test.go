package layout_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutUsecase_DeleteLayout(t *testing.T) {
	setupLayoutUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("既存のレイアウトを削除できる", func(t *testing.T) {
			// テスト用のレイアウトを作成
			layout := createTestLayout(t, generateUniqueTitle())

			t.Logf("レイアウト削除: ID=%d", layout.ID)

			// テスト実行
			err := layoutUsecase.DeleteLayout(testUserId, layout.ID)

			// 検証
			if err != nil {
				t.Errorf("DeleteLayout() error = %v", err)
			}

			// データベースから直接確認
			var deletedLayout model.Layout
			result := layoutDb.First(&deletedLayout, layout.ID)

			// レコードが見つからないはず
			if result.Error == nil {
				t.Errorf("DeleteLayout() failed: layout with ID %d still exists", layout.ID)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないレイアウトID
			nonExistLayoutId := uint(9999)

			// テスト実行
			err := layoutUsecase.DeleteLayout(testUserId, nonExistLayoutId)

			// 検証
			if err == nil {
				t.Error("存在しないレイアウトIDでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})

		t.Run("他のユーザーのレイアウトを削除しようとするとエラーを返す", func(t *testing.T) {
			// テスト用のレイアウトを作成
			layout := createTestLayout(t, generateUniqueTitle())

			// 別のユーザーID
			otherUserId := uint(999)

			// テスト実行
			err := layoutUsecase.DeleteLayout(otherUserId, layout.ID)

			// 検証
			if err == nil {
				t.Error("他のユーザーのレイアウト削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}

			// 元のレイアウトが削除されていないことを確認
			var savedLayout model.Layout
			result := layoutDb.First(&savedLayout, layout.ID)

			if result.Error != nil {
				t.Errorf("他のユーザーによる削除試行で元のレイアウトが削除されました: %v", result.Error)
			}
		})
	})
}
