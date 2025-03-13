package main_entry_module

import (
	"gorm.io/gorm"
	
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initUserModule(db *gorm.DB) {
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	m.UserController = controller.NewUserController(userUsecase)
}
