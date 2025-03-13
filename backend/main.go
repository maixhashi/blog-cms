package main

import (
	"go-react-app/db"
	"go-react-app/main_entry_module"
	"log"
)

func main() {
	// データベース接続
	db := db.NewDB()
	
	// アプリケーション初期化
	app := main_entry_module.NewMainEntryPackage(db)
	
	// サーバー起動（エラーがあれば終了）
	if err := app.StartServer(); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}