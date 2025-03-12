package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupFeedArticleRoutes はフィード記事関連のルートを設定します
func SetupFeedArticleRoutes(e *echo.Echo, fac controller.IFeedArticleController) {
	fa := e.Group("/feed-articles")
	fa.Use(middleware.GetJWTMiddleware())
	fa.GET("/:feedId", fac.GetArticlesByFeedID)
	fa.GET("/:feedId/:articleId", fac.GetArticleByID)
	fa.GET("", fac.GetAllArticles)
}