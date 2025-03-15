package feed_test

import (
    "go-react-app/model"
    "go-react-app/repository"
    "go-react-app/testutils"
    "gorm.io/gorm"
)

var (
    feedDB *gorm.DB
    feedRepo repository.IFeedRepository
    feedTestUser model.User
    feedOtherUser model.User
    nonExistentFeedID uint = 9999
)

func setupFeedTest() {
    feedDB = testutils.SetupTestDB()
    feedRepo = repository.NewFeedRepository(feedDB)
    
    feedTestUser = testutils.CreateTestUser(feedDB)
    feedOtherUser = testutils.CreateOtherUser(feedDB)
}
