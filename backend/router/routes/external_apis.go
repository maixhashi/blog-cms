package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupExternalAPIRoutes は外部API関連のルートを設定します
func SetupExternalAPIRoutes(e *echo.Echo, ac controller.IExternalAPIController) {
	a := e.Group("/external-apis")
	a.Use(middleware.GetJWTMiddleware())
	a.GET("", ac.GetAllExternalAPIs)
	a.GET("/:apiId", ac.GetExternalAPIById)
	a.POST("", ac.CreateExternalAPI)
	a.PUT("/:apiId", ac.UpdateExternalAPI)
	a.DELETE("/:apiId", ac.DeleteExternalAPI)
}