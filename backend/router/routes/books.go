package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupBookRoutes は書籍関連のルートを設定します
func SetupBookRoutes(e *echo.Echo, bc controller.IBookController) {
	b := e.Group("/books")
	b.Use(middleware.GetJWTMiddleware())
	b.GET("", bc.GetAllBooks)
	b.GET("/:bookId", bc.GetBookById)
	b.POST("", bc.CreateBook)
	b.PUT("/:bookId", bc.UpdateBook)
	b.DELETE("/:bookId", bc.DeleteBook)
}
