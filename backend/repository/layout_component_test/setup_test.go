package layout_component_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"gorm.io/gorm"
)

var (
	lcRepoDb   *gorm.DB
	lcRepo     repository.ILayoutComponentRepository
	userRepo   repository.IUserRepository
	testUser   model.User
	testLayoutComponent = model.LayoutComponent{
		Name:    "Test Component",
		Type:    "header",
		Content: `{"text": "Hello World"}`,
	}
)

func setupLayoutComponentRepositoryTest() {
	if lcRepoDb != nil {
		testutils.CleanupTestDB(lcRepoDb)
	} else {
		lcRepoDb = testutils.SetupTestDB()
		lcRepo = repository.NewLayoutComponentRepository(lcRepoDb)
		userRepo = repository.NewUserRepository(lcRepoDb)
	}
	
	// テスト用ユーザーを作成
	testUser = model.User{
		Email:    "layout-test@example.com",
		Password: "password123",
	}
	userRepo.CreateUser(&testUser)
}