package article_test

import (
    "go-react-app/model"
    "go-react-app/repository"
    "go-react-app/testutils"
    "gorm.io/gorm"
)

var (
    articleDB *gorm.DB
    articleRepo repository.IArticleRepository
    articleTestUser model.User
    articleOtherUser model.User
    nonExistentArticleID uint = 9999
)

func setupArticleTest() {
    articleDB = testutils.SetupTestDB()
    articleRepo = repository.NewArticleRepository(articleDB)
    
    articleTestUser = testutils.CreateTestUser(articleDB)
    articleOtherUser = testutils.CreateOtherUser(articleDB)
}
