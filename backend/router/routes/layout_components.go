package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

func SetupLayoutComponentRoutes(e *echo.Echo, lcc controller.ILayoutComponentController) {
	lc := e.Group("/layout-components")
	lc.Use(middleware.GetJWTMiddleware())
	
	// ルートの設定
	lc.GET("", lcc.GetAllLayoutComponents)
	lc.GET("/:componentId", lcc.GetLayoutComponentById)
	lc.POST("", lcc.CreateLayoutComponent)
	lc.PUT("/:componentId", lcc.UpdateLayoutComponent)
	lc.DELETE("/:componentId", lcc.DeleteLayoutComponent)
}
