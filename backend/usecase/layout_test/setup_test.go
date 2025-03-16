package layout_test

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
	layoutDb        *gorm.DB
	layoutRepo      repository.ILayoutRepository
	layoutValidator validator.ILayoutValidator
	layoutUsecase   usecase.ILayoutUsecase
	testUserId      uint = 1 // テスト用ユーザーID
)

// テスト前の共通セットアップ
func setupLayoutUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if layoutDb != nil {
		testutils.CleanupTestDB(layoutDb)
	} else {
		// 初回のみデータベース接続を作成
		layoutDb = testutils.SetupTestDB()
		layoutRepo = repository.NewLayoutRepository(layoutDb)
		layoutValidator = validator.NewLayoutValidator()
		layoutUsecase = usecase.NewLayoutUsecase(layoutRepo, layoutValidator)
	}

	// 既存のテストレイアウトを明示的に削除（念のため）
	layoutDb.Exec("DELETE FROM layouts WHERE user_id = ?", testUserId)
}

// テスト用のレイアウトを作成
func createTestLayout(t *testing.T, title string) model.LayoutResponse {
	layout := model.Layout{
		Title:  title,
		UserId: testUserId,
	}

	createdLayout, err := layoutUsecase.CreateLayout(layout)
	if err != nil {
		t.Fatalf("テストレイアウトの作成に失敗しました: %v", err)
	}

	return createdLayout
}
