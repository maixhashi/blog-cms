package feed_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/usecase"
	"go-react-app/validator"

	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	feedDB          *gorm.DB
	feedRepo        repository.IFeedRepository
	feedValidator   validator.IFeedValidator
	feedUsecase     usecase.IFeedUsecase
	feedTestUser    model.User
	feedOtherUser   model.User
)

const nonExistentFeedID uint = 9999

// テスト前の共通セットアップ
func setupFeedTest() {
	// テストごとにデータベースをクリーンアップ
	if feedDB != nil {
		testutils.CleanupTestDB(feedDB)
	} else {
		// 初回のみデータベース接続を作成
		feedDB = testutils.SetupTestDB()
		feedRepo = repository.NewFeedRepository(feedDB)
		feedValidator = validator.NewFeedValidator()
		feedUsecase = usecase.NewFeedUsecase(feedRepo, feedValidator)
	}
	
	// テストユーザーを作成
	feedTestUser = testutils.CreateTestUser(feedDB)
	
	// 別のテストユーザーを作成
	feedOtherUser = testutils.CreateOtherUser(feedDB)
}

// テスト用のフィードを作成するヘルパー関数
func createTestFeed(title string, url string, userId uint) model.Feed {
	feed := model.Feed{
		Title:  title,
		URL:    url,
		UserId: userId,
	}
	feedDB.Create(&feed)
	return feed
}
