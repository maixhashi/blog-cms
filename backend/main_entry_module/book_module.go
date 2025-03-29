package main_entry_module

import (
	"gorm.io/gorm"
	
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

// initBookModule は書籍関連のモジュールを初期化します
// @Summary 書籍モジュールの初期化
// @Description 書籍関連のリポジトリ、ユースケース、コントローラーを初期化します
func (m *MainEntryPackage) initBookModule(db *gorm.DB) {
	bookValidator := validator.NewBookValidator()
	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository, bookValidator)
	m.BookController = controller.NewBookController(bookUsecase)
}
