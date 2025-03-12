package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupHatenaRoutes はHatena関連のルートを設定します
func SetupHatenaRoutes(e *echo.Echo, hc controller.IHatenaController) {
	hatena := e.Group("/hatena")
	hatena.Use(middleware.GetJWTMiddleware())
	hatena.GET("", hc.GetHatenaArticles)
	hatena.GET("/:id", hc.GetHatenaArticleByID)
}