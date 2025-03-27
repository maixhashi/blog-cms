package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupGoogleBookRoutes はGoogle Books API関連のルートを設定します
func SetupGoogleBookRoutes(e *echo.Echo, gbc controller.IGoogleBookController) {
	gb := e.Group("/google-books")
	gb.Use(middleware.GetJWTMiddleware())
	gb.POST("/search", gbc.SearchBooks)
	gb.GET("/:id", gbc.GetBookByID)
	gb.POST("/:id/import", gbc.ImportBookFromGoogle)
}
