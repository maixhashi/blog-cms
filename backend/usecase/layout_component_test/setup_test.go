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
	db                      *gorm.DB
	layoutComponentRepo     repository.ILayoutComponentRepository
	layoutComponentValidator validator.ILayoutComponentValidator
	layoutComponentUsecase  usecase.ILayoutComponentUsecase
	testUserId              uint = 1 // テスト用ユーザーID
)

// テスト前の共通セットアップ
func setupLayoutComponentUsecaseTest() {
	// テストデータベースのセットアップ（初回のみ）
	if db == nil {
		db = testutils.SetupTestDB()
		
		// リポジトリとバリデータの初期化
		layoutComponentRepo = repository.NewLayoutComponentRepository(db)
		layoutComponentValidator = validator.NewLayoutComponentValidator()
		
		// ユースケースの初期化
		layoutComponentUsecase = usecase.NewLayoutComponentUsecase(layoutComponentRepo, layoutComponentValidator)
	}
	
	// テストデータベースのクリーンアップ - より徹底的に
	db.Exec("DELETE FROM layout_components")
	db.Exec("DELETE FROM layouts")
	db.Exec("DELETE FROM users")
	
	// テスト用ユーザーを作成
	testUser := model.User{
		ID:       testUserId,
		Email:    "test@example.com",
		Password: "password123",
	}
	db.Create(&testUser)
}

// テスト用のレイアウトコンポーネントを作成
func createTestLayoutComponent(t *testing.T) model.LayoutComponent {
	component := model.LayoutComponent{
		Name:    "Test Component",
		Type:    "section",
		Content: `{"title": "Test Section", "text": "This is a test section"}`,
		UserId:  testUserId,
	}
	
	if err := db.Create(&component).Error; err != nil {
		t.Fatalf("テスト用コンポーネントの作成に失敗しました: %v", err)
	}
	
	return component
}

// テスト用のレイアウトを作成
func createTestLayout(t *testing.T) model.Layout {
	layout := model.Layout{
		Title:  "Test Layout",
		UserId: testUserId,
	}
	
	if err := db.Create(&layout).Error; err != nil {
		t.Fatalf("テスト用レイアウトの作成に失敗しました: %v", err)
	}
	
	return layout
}