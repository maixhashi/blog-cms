package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupLayoutRoutes はレイアウト関連のルートを設定します
func SetupLayoutRoutes(e *echo.Echo, lc controller.ILayoutController) {
	l := e.Group("/layouts")
	l.Use(middleware.GetJWTMiddleware())
	l.GET("", lc.GetAllLayouts)
	l.GET("/:layoutId", lc.GetLayoutById)
	l.POST("", lc.CreateLayout)
	l.PUT("/:layoutId", lc.UpdateLayout)
	l.DELETE("/:layoutId", lc.DeleteLayout)
}