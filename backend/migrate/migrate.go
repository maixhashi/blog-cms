package main

import (
	"fmt"
	"go-react-app/db"
	"go-react-app/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.Feed{},
		&model.ExternalAPI{},
		&model.Article{},
		&model.Layout{},
		&model.LayoutComponent{},
		&model.Book{},
	)
}