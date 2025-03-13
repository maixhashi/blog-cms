package main_entry_module

import (
	"gorm.io/gorm"

	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initLayoutComponentModule(db *gorm.DB) {
	layoutComponentValidator := validator.NewLayoutComponentValidator()
	layoutComponentRepository := repository.NewLayoutComponentRepository(db)
	layoutComponentUsecase := usecase.NewLayoutComponentUsecase(layoutComponentRepository, layoutComponentValidator)
	m.LayoutComponentController = controller.NewLayoutComponentController(layoutComponentUsecase)
}
