package main_entry_module

import (
	"gorm.io/gorm"
	
	"go-react-app/controller"
	"go-react-app/repository"
	"go-react-app/usecase"
	"go-react-app/validator"
)

// initGoogleBookModule はGoogle Books API関連のモジュールを初期化します
// @Summary Google Booksモジュールの初期化
// @Description Google Books API関連のリポジトリ、ユースケース、コントローラーを初期化します
func (m *MainEntryPackage) initGoogleBookModule(db *gorm.DB) {
	bookValidator := validator.NewBookValidator()
	googleBookRepository := repository.NewGoogleBookRepository()
	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository, bookValidator)
	googleBookUsecase := usecase.NewGoogleBookUsecase(googleBookRepository, bookValidator)
	m.GoogleBookController = controller.NewGoogleBookController(googleBookUsecase, bookUsecase)
}
