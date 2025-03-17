package layout_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"gorm.io/gorm"
)

var (
	layoutRepoDb   *gorm.DB
	layoutRepo     repository.ILayoutRepository
	userRepo       repository.IUserRepository
	testUserData   model.User
	testLayoutData model.Layout
)

func setupLayoutRepositoryTest() {
	if layoutRepoDb != nil {
		testutils.CleanupTestDB(layoutRepoDb)
	} else {
		layoutRepoDb = testutils.SetupTestDB()
		layoutRepo = repository.NewLayoutRepository(layoutRepoDb)
		userRepo = repository.NewUserRepository(layoutRepoDb)
	}

	// テスト用ユーザーを作成
	testUserData = model.User{
		Email:    "layout-test@example.com",
		Password: "password123",
	}
	userRepo.CreateUser(&testUserData)
	// テスト用レイアウトデータを初期化
	testLayoutData = model.Layout{
		Title:  "Test Layout",
		UserId: testUserData.ID,
	}
}