package main_entry_module

import (
	"gorm.io/gorm"
	
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

func (m *MainEntryPackage) initTaskModule(db *gorm.DB) {
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	m.TaskController = controller.NewTaskController(taskUsecase)
}
