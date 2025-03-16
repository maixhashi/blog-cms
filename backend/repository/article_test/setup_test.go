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
    // テスト用DBの設定
    articleDB = testutils.SetupTestDB()
    articleRepo = repository.NewArticleRepository(articleDB)
    
    // テストユーザーの作成
    articleTestUser = testutils.CreateTestUser(articleDB)
    articleOtherUser = testutils.CreateOtherUser(articleDB)
}
