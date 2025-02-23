package main

import (
	"go-react-app/db"

	"go-react-app/validator"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/controller"
	
	"go-react-app/router"
)

func main() {
	db := db.NewDB()

	// ユーザー認証
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	// タスクのCRUD機能
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	// タスクのCRUD機能
	feedValidator := validator.NewFeedValidator()
	feedRepository := repository.NewFeedRepository(db)
	feedUsecase := usecase.NewFeedUsecase(feedRepository, feedValidator)
	feedController := controller.NewFeedController(feedUsecase)

	e := router.NewRouter(userController, taskController, feedController)
	e.Logger.Fatal(e.Start(":8080"))
}