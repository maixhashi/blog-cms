package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupArticleRoutes は記事関連のルートを設定します
func SetupArticleRoutes(e *echo.Echo, artc controller.IArticleController) {
	art := e.Group("/articles")
	art.Use(middleware.GetJWTMiddleware())
	art.GET("", artc.GetAllArticles)
	art.GET("/:articleId", artc.GetArticleById)
	art.POST("", artc.CreateArticle)
	art.PUT("/:articleId", artc.UpdateArticle)
	art.DELETE("/:articleId", artc.DeleteArticle)
}