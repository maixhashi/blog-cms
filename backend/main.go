package main

import (
	"go-react-app/db"
	"go-react-app/main_entry_module"
	"go-react-app/router"
)

func main() {
	db := db.NewDB()
	
	// モジュール化された初期化処理
	entryPackage := main_entry_module.NewMainEntryPackage(db)
	
	// ルーターの設定
	e := router.NewRouter(
		entryPackage.UserController,
		entryPackage.TaskController,
		entryPackage.FeedController,
		entryPackage.ExternalAPIController,
		entryPackage.QiitaController,
		entryPackage.HatenaController,
		entryPackage.ArticleController,
		entryPackage.FeedArticleController,
		entryPackage.LayoutController,
		entryPackage.LayoutComponentController,
	)
	
	e.Logger.Fatal(e.Start(":8080"))
}