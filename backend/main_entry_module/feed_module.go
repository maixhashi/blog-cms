package main_entry_module

import (
	"gorm.io/gorm"
	
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initFeedModule(db *gorm.DB) {
	feedValidator := validator.NewFeedValidator()
	feedRepository := repository.NewFeedRepository(db)
	feedUsecase := usecase.NewFeedUsecase(feedRepository, feedValidator)
	m.FeedController = controller.NewFeedController(feedUsecase)
}
