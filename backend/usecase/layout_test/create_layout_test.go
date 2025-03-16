package layout_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutUsecase_CreateLayout(t *testing.T) {
	setupLayoutUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("新規レイアウトを作成できる", func(t *testing.T) {
			// テスト用のレイアウト
			title := generateUniqueTitle()
			layout := model.Layout{
				Title:  title,
				UserId: testUserId,
			}

			t.Logf("レイアウト作成: Title=%s", layout.Title)

			// テスト実行
			createdLayout, err := layoutUsecase.CreateLayout(layout)

			// 検証
			if err != nil {
				t.Errorf("CreateLayout() error = %v", err)
			}

			if createdLayout.ID == 0 || createdLayout.Title != title {
				t.Errorf("CreateLayout() = %v, want title=%s and ID > 0", createdLayout, title)
			} else {
				t.Logf("生成されたレイアウトID: %d", createdLayout.ID)
			}

			// データベースから直接確認
			var savedLayout model.Layout
			layoutDb.First(&savedLayout, createdLayout.ID)

			if savedLayout.Title != title {
				t.Errorf("CreateLayout() saved title = %v, want %v", savedLayout.Title, title)
			}

			if savedLayout.UserId != testUserId {
				t.Errorf("CreateLayout() saved userId = %v, want %v", savedLayout.UserId, testUserId)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合はレイアウト作成に失敗する", func(t *testing.T) {
			// 空のタイトル（バリデーションエラーになるはず）
			invalidLayout := model.Layout{
				Title:  "",
				UserId: testUserId,
			}

			t.Log("無効なレイアウト作成を試行: Title=空文字")

			// テスト実行
			_, err := layoutUsecase.CreateLayout(invalidLayout)

			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("空のタイトルでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})
	})
}