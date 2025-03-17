package layout_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutUsecase_UpdateLayout(t *testing.T) {
	setupLayoutUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("既存のレイアウトを更新できる", func(t *testing.T) {
			// テスト用のレイアウトを作成
			originalLayout := createTestLayout(t, generateUniqueTitle())

			// 更新用のデータ
			newTitle := generateUniqueTitle() + " Updated"
			updateLayoutRequest := model.LayoutRequest{
				Title:  newTitle,
				UserId: testUserId,
			}

			t.Logf("レイアウト更新: ID=%d, 新Title=%s", originalLayout.ID, newTitle)

			// テスト実行
			updatedLayout, err := layoutUsecase.UpdateLayout(updateLayoutRequest, testUserId, originalLayout.ID)

			// 検証
			if err != nil {
				t.Errorf("UpdateLayout() error = %v", err)
			}

			if updatedLayout.ID != originalLayout.ID {
				t.Errorf("UpdateLayout() ID = %v, want %v", updatedLayout.ID, originalLayout.ID)
			}

			if updatedLayout.Title != newTitle {
				t.Errorf("UpdateLayout() title = %v, want %v", updatedLayout.Title, newTitle)
			}

			// データベースから直接確認
			var savedLayout model.Layout
			layoutDb.First(&savedLayout, originalLayout.ID)

			if savedLayout.Title != newTitle {
				t.Errorf("UpdateLayout() saved title = %v, want %v", savedLayout.Title, newTitle)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合はレイアウト更新に失敗する", func(t *testing.T) {
			// テスト用のレイアウトを作成
			layout := createTestLayout(t, generateUniqueTitle())

			// 空のタイトル（バリデーションエラーになるはず）
			invalidLayout := model.LayoutRequest{
				Title:  "",
				UserId: testUserId,
			}

			t.Logf("無効なレイアウト更新を試行: ID=%d, Title=空文字", layout.ID)

			// テスト実行
			_, err := layoutUsecase.UpdateLayout(invalidLayout, testUserId, layout.ID)

			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("空のタイトルでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})

		t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないレイアウトID
			nonExistLayoutId := uint(9999)

			updateLayout := model.LayoutRequest{
				Title:  generateUniqueTitle(),
				UserId: testUserId,
			}

			// テスト実行
			_, err := layoutUsecase.UpdateLayout(updateLayout, testUserId, nonExistLayoutId)

			// 検証
			if err == nil {
				t.Error("存在しないレイアウトIDでエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})

		t.Run("他のユーザーのレイアウトを更新しようとするとエラーを返す", func(t *testing.T) {
			// テスト用のレイアウトを作成
			layout := createTestLayout(t, generateUniqueTitle())

			// 別のユーザーID
			otherUserId := uint(999)

			updateLayout := model.LayoutRequest{
				Title:  generateUniqueTitle(),
				UserId: otherUserId,
			}

			// テスト実行
			_, err := layoutUsecase.UpdateLayout(updateLayout, otherUserId, layout.ID)

			// 検証
			if err == nil {
				t.Error("他のユーザーのレイアウト更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}