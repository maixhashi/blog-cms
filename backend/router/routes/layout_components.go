package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

func SetupLayoutComponentRoutes(e *echo.Echo, lcc controller.ILayoutComponentController) {
	lc := e.Group("/layout-components")
	lc.Use(middleware.GetJWTMiddleware())
	
	// 既存のルート
	lc.GET("", lcc.GetAllLayoutComponents)
	lc.GET("/:componentId", lcc.GetLayoutComponentById)
	lc.POST("", lcc.CreateLayoutComponent)
	lc.PUT("/:componentId", lcc.UpdateLayoutComponent)
	lc.DELETE("/:componentId", lcc.DeleteLayoutComponent)
	
	lc.POST("/:componentId/assign/:layoutId", lcc.AssignToLayout)
	lc.DELETE("/:componentId/assign", lcc.RemoveFromLayout)
	lc.PUT("/:componentId/position", lcc.UpdatePosition)
}