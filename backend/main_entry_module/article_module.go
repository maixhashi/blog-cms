package main_entry_module

import (
	"gorm.io/gorm"

	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initArticleModule(db *gorm.DB) {
	articleValidator := validator.NewArticleValidator()
	articleRepository := repository.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository, articleValidator)
	m.ArticleController = controller.NewArticleController(articleUsecase)
}
