package article_test

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
	articleDb         *gorm.DB
	articleRepo       repository.IArticleRepository
	articleValidator  validator.IArticleValidator
	articleUsecase    usecase.IArticleUsecase
	articleTestUser   model.User
	articleOtherUser  model.User
)

const nonExistentArticleID uint = 9999

// テスト前の共通セットアップ
func setupArticleUsecaseTest() {
	// テストごとにデータベースをクリーンアップ
	if articleDb != nil {
		testutils.CleanupTestDB(articleDb)
	} else {
		// 初回のみデータベース接続を作成
		articleDb = testutils.SetupTestDB()
		articleRepo = repository.NewArticleRepository(articleDb)
		articleValidator = validator.NewArticleValidator()
		articleUsecase = usecase.NewArticleUsecase(articleRepo, articleValidator)
	}
	
	// テストユーザーを作成
	articleTestUser = testutils.CreateTestUser(articleDb)
	
	// 別のテストユーザーを作成
	articleOtherUser = testutils.CreateOtherUser(articleDb)
}

// テスト用の記事を作成するヘルパー関数
func createTestArticle(t *testing.T, title string, content string, userId uint) model.Article {
	// ArticleRequestを使用
	request := model.ArticleRequest{
		Title:   title,
		Content: content,
		UserId:  userId,
	}
	
	article := request.ToModel()
	result := articleDb.Create(&article)
	if result.Error != nil {
		t.Fatalf("テスト記事の作成に失敗しました: %v", result.Error)
	}
	
	return article
}
