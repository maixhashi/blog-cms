package layout_component_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/usecase"
	"go-react-app/validator"
	"testing"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	componentDb        *gorm.DB
	componentRepo      repository.ILayoutComponentRepository
	componentValidator validator.ILayoutComponentValidator
	componentUsecase   usecase.ILayoutComponentUsecase
	testUserId         uint = 1 // テスト用ユーザーID
)

// テスト前の共通セットアップ
func setupLayoutComponentUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if componentDb != nil {
		testutils.CleanupTestDB(componentDb)
	} else {
		// 初回のみデータベース接続を作成
		componentDb = testutils.SetupTestDB()
		componentRepo = repository.NewLayoutComponentRepository(componentDb)
		componentValidator = validator.NewLayoutComponentValidator()
		componentUsecase = usecase.NewLayoutComponentUsecase(componentRepo, componentValidator)
	}

	// 既存のテストコンポーネントを明示的に削除（念のため）
	componentDb.Exec("DELETE FROM layout_components WHERE user_id = ?", testUserId)
}

// テスト用のレイアウトコンポーネントを作成
func createTestComponent(t *testing.T, name string) model.LayoutComponentResponse {
	componentRequest := model.LayoutComponentRequest{
		Name:    name,
		Type:    "text",
		Content: "テストコンテンツ",
		UserId:  testUserId,
	}

	createdComponent, err := componentUsecase.CreateLayoutComponent(componentRequest)
	if err != nil {
		t.Fatalf("テストコンポーネントの作成に失敗しました: %v", err)
	}

	return createdComponent
}
