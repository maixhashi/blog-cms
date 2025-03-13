package main_entry_module

import (
	"gorm.io/gorm"

	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initExternalAPIModule(db *gorm.DB) {
	externalAPIValidator := validator.NewExternalAPIValidator()
	externalAPIRepository := repository.NewExternalAPIRepository(db)
	externalAPIUsecase := usecase.NewExternalAPIUsecase(externalAPIRepository, externalAPIValidator)
	m.ExternalAPIController = controller.NewExternalAPIController(externalAPIUsecase)
}
