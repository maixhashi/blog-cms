package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupFeedRoutes はフィード関連のルートを設定します
func SetupFeedRoutes(e *echo.Echo, fc controller.IFeedController) {
	f := e.Group("/feeds")
	f.Use(middleware.GetJWTMiddleware())
	f.GET("", fc.GetAllFeeds)
	f.GET("/:feedId", fc.GetFeedById)
	f.POST("", fc.CreateFeed)
	f.PUT("/:feedId", fc.UpdateFeed)
	f.DELETE("/:feedId", fc.DeleteFeed)
}