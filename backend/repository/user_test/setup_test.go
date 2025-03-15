package user_test

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"gorm.io/gorm"
)

var (
	userRepoDb   *gorm.DB
	userRepo     repository.IUserRepository
	testUserData = model.User{
		Email:    "test@example.com",
		Password: "password123",
	}
)

func setupUserRepositoryTest() {
	if userRepoDb != nil {
		testutils.CleanupTestDB(userRepoDb)
	} else {
		userRepoDb = testutils.SetupTestDB()
		userRepo = repository.NewUserRepository(userRepoDb)
	}
}
