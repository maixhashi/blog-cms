package layout_component_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"gorm.io/gorm"
)

var (
	layoutComponentRepoDb   *gorm.DB
	layoutComponentRepo     repository.ILayoutComponentRepository
	layoutRepo              repository.ILayoutRepository
	userRepo                repository.IUserRepository
	testUserData            model.User
	testLayoutData          model.Layout
	testLayoutComponentData model.LayoutComponent
)

func setupLayoutComponentRepositoryTest() {
	if layoutComponentRepoDb != nil {
		testutils.CleanupTestDB(layoutComponentRepoDb)
	} else {
		layoutComponentRepoDb = testutils.SetupTestDB()
		layoutComponentRepo = repository.NewLayoutComponentRepository(layoutComponentRepoDb)
		layoutRepo = repository.NewLayoutRepository(layoutComponentRepoDb)
		userRepo = repository.NewUserRepository(layoutComponentRepoDb)
	}

	// テスト用ユーザーを作成
	testUserData = model.User{
		Email:    "layout-component-test@example.com",
		Password: "password123",
	}
	userRepo.CreateUser(&testUserData)
	
	// テスト用レイアウトを作成
	testLayoutData = model.Layout{
		Title:  "Test Layout",
		UserId: testUserData.ID,
	}
	layoutRepo.CreateLayout(&testLayoutData)
	
	// テスト用レイアウトコンポーネントデータを初期化
	testLayoutComponentData = model.LayoutComponent{
		Name:    "Test Component",
		Type:    "text",
		Content: "Test Content",
		X:       0,
		Y:       0,
		Width:   100,
		Height:  100,
		UserId:  testUserData.ID,
	}
}
