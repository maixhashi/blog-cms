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

	// フィードのCRUD機能
	feedValidator := validator.NewFeedValidator()
	feedRepository := repository.NewFeedRepository(db)
	feedUsecase := usecase.NewFeedUsecase(feedRepository, feedValidator)
	feedController := controller.NewFeedController(feedUsecase)

	// 外部APIのCRUD機能
	externalAPIValidator := validator.NewExternalAPIValidator()
	externalAPIRepository := repository.NewExternalAPIRepository(db)
	externalAPIUsecase := usecase.NewExternalAPIUsecase(externalAPIRepository, externalAPIValidator)
	externalAPIController := controller.NewExternalAPIController(externalAPIUsecase)

	// 記事のCRUD機能
	articleValidator := validator.NewArticleValidator()
	articleRepository := repository.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository, articleValidator)
	articleController := controller.NewArticleController(articleUsecase)

	qiitaRepository := repository.NewQiitaRepository()
	qiitaUsecase := usecase.NewQiitaUsecase(qiitaRepository)
	qiitaController := controller.NewQiitaController(qiitaUsecase)

	hatenaRepository := repository.NewHatenaRepository("https://tech.smarthr.jp/feed?exclude_body=1")
	hatenaUsecase := usecase.NewHatenaUsecase(hatenaRepository)
	hatenaController := controller.NewHatenaController(hatenaUsecase)

	// FeedArticle関連の依存関係を設定
	feedArticleRepository := repository.NewFeedArticleRepository(feedRepository)
	feedArticleUsecase := usecase.NewFeedArticleUsecase(feedArticleRepository)
	feedArticleController := controller.NewFeedArticleController(feedArticleUsecase)

	e := router.NewRouter(userController, taskController, feedController, externalAPIController, 
		qiitaController, hatenaController, articleController, feedArticleController)
	e.Logger.Fatal(e.Start(":8080"))
}
