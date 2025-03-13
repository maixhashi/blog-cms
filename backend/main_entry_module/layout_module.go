package main_entry_module

import (
	"gorm.io/gorm"

	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initLayoutModule(db *gorm.DB) {
	layoutValidator := validator.NewLayoutValidator()
	layoutRepository := repository.NewLayoutRepository(db)
	layoutUsecase := usecase.NewLayoutUsecase(layoutRepository, layoutValidator)
	m.LayoutController = controller.NewLayoutController(layoutUsecase)
}
