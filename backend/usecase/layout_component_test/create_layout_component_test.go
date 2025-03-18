package layout_component_test

import (
	"go-react-app/model"
	"testing"
)

func TestLayoutComponentUsecase_CreateLayoutComponent(t *testing.T) {
	setupLayoutComponentUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("新規コンポーネントを作成できる", func(t *testing.T) {
			// テスト用のコンポーネント
			name := generateUniqueName()
			componentRequest := model.LayoutComponentRequest{
				Name:    name,
				Type:    "text",
				Content: "テストコンテンツ",
				UserId:  testUserId,
			}

			t.Logf("コンポーネント作成: Name=%s", componentRequest.Name)

			// テスト実行
			createdComponent, err := componentUsecase.CreateLayoutComponent(componentRequest)

			// 検証
			if err != nil {
				t.Errorf("CreateLayoutComponent() error = %v", err)
			}

			if createdComponent.ID == 0 || createdComponent.Name != name {
				t.Errorf("CreateLayoutComponent() = %v, want name=%s and ID > 0", createdComponent, name)
			} else {
				t.Logf("生成されたコンポーネントID: %d", createdComponent.ID)
			}

			// データベースから直接確認
			var savedComponent model.LayoutComponent
			componentDb.First(&savedComponent, createdComponent.ID)

			if savedComponent.Name != name {
				t.Errorf("CreateLayoutComponent() saved name = %v, want %v", savedComponent.Name, name)
			}

			if savedComponent.UserId != testUserId {
				t.Errorf("CreateLayoutComponent() saved userId = %v, want %v", savedComponent.UserId, testUserId)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する場合はコンポーネント作成に失敗する", func(t *testing.T) {
			// 空の名前（バリデーションエラーになるはず）
			invalidComponent := model.LayoutComponentRequest{
				Name:    "",
				Type:    "text",
				Content: "テストコンテンツ",
				UserId:  testUserId,
			}

			t.Log("無効なコンポーネント作成を試行: Name=空文字")

			// テスト実行
			_, err := componentUsecase.CreateLayoutComponent(invalidComponent)

			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("空の名前でエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
		})
	})
}
