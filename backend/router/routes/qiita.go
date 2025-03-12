package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupQiitaRoutes はQiita関連のルートを設定します
func SetupQiitaRoutes(e *echo.Echo, qc controller.IQiitaController) {
	q := e.Group("/qiita")
	q.Use(middleware.GetJWTMiddleware())
	q.GET("/articles", qc.GetQiitaArticles)
	q.GET("/articles/:id", qc.GetQiitaArticleByID)
}
