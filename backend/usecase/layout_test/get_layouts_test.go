package layout_test

import (
	"testing"
)

func TestLayoutUsecase_GetAllLayouts(t *testing.T) {
	setupLayoutUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのレイアウト一覧を取得できる", func(t *testing.T) {
			// テスト用のレイアウトを複数作成
			layout1 := createTestLayout(t, generateUniqueTitle())
			layout2 := createTestLayout(t, generateUniqueTitle())
			layout3 := createTestLayout(t, generateUniqueTitle())

			// テスト実行
			layouts, err := layoutUsecase.GetAllLayouts(testUserId)

			// 検証
			if err != nil {
				t.Errorf("GetAllLayouts() error = %v", err)
			}

			// 少なくとも作成した3つのレイアウトが含まれているか確認
			if len(layouts) < 3 {
				t.Errorf("GetAllLayouts() returned %d layouts, want at least 3", len(layouts))
			}

			// 作成したレイアウトが含まれているか確認
			foundLayout1 := false
			foundLayout2 := false
			foundLayout3 := false

			for _, layout := range layouts {
				if layout.ID == layout1.ID {
					foundLayout1 = true
				}
				if layout.ID == layout2.ID {
					foundLayout2 = true
				}
				if layout.ID == layout3.ID {
					foundLayout3 = true
				}
			}

			if !foundLayout1 || !foundLayout2 || !foundLayout3 {
				t.Errorf("GetAllLayouts() did not return all created layouts")
			}
		})

		t.Run("レイアウトが存在しない場合は空の配列を返す", func(t *testing.T) {
			// 別のユーザーIDを使用
			nonExistUserId := uint(999)

			// テスト実行
			layouts, err := layoutUsecase.GetAllLayouts(nonExistUserId)

			// 検証
			if err != nil {
				t.Errorf("GetAllLayouts() error = %v", err)
			}

			if len(layouts) != 0 {
				t.Errorf("GetAllLayouts() returned %d layouts, want 0", len(layouts))
			}
		})
	})
}

func TestLayoutUsecase_GetLayoutById(t *testing.T) {
	setupLayoutUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("指定したIDのレイアウトを取得できる", func(t *testing.T) {
			// テスト用のレイアウトを作成
			expectedLayout := createTestLayout(t, generateUniqueTitle())

			// テスト実行
			layout, err := layoutUsecase.GetLayoutById(testUserId, expectedLayout.ID)

			// 検証
			if err != nil {
				t.Errorf("GetLayoutById() error = %v", err)
			}

			if layout.ID != expectedLayout.ID {
				t.Errorf("GetLayoutById() = %v, want %v", layout.ID, expectedLayout.ID)
			}

			if layout.Title != expectedLayout.Title {
				t.Errorf("GetLayoutById() title = %v, want %v", layout.Title, expectedLayout.Title)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないレイアウトIDの場合はエラーを返す", func(t *testing.T) {
			// 存在しないレイアウトID
			nonExistLayoutId := uint(9999)

			// テスト実行
			_, err := layoutUsecase.GetLayoutById(testUserId, nonExistLayoutId)

			// 検証
			if err == nil {
				t.Error("GetLayoutById() error = nil, want error for non-existent layout")
			}
		})

		t.Run("他のユーザーのレイアウトにアクセスするとエラーを返す", func(t *testing.T) {
			// テスト用のレイアウトを作成
			layout := createTestLayout(t, generateUniqueTitle())

			// 別のユーザーIDを使用
			otherUserId := uint(999)

			// テスト実行
			_, err := layoutUsecase.GetLayoutById(otherUserId, layout.ID)

			// 検証
			if err == nil {
				t.Error("GetLayoutById() error = nil, want error for accessing other user's layout")
			}
		})
	})
}