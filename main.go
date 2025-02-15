package main

import (
	"go-react-app/controller"
	"go-react-app/db"
	"go-react-app/repository"
	"go-react-app/router"
	"go-react-app/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}