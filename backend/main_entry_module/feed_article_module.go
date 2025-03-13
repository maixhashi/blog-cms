package main_entry_module

import (
	"gorm.io/gorm"

	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
)

func (m *MainEntryPackage) initFeedArticleModule(db *gorm.DB) {
	feedRepository := repository.NewFeedRepository(db)
	feedArticleRepository := repository.NewFeedArticleRepository(feedRepository)
	feedArticleUsecase := usecase.NewFeedArticleUsecase(feedArticleRepository)
	m.FeedArticleController = controller.NewFeedArticleController(feedArticleUsecase)
}
