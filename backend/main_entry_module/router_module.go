package main_entry_module

import (
	"github.com/labstack/echo/v4"
	"go-react-app/router"
	
	echoSwagger "github.com/swaggo/echo-swagger" // Swaggerのインポート
)

// InitRouter はルーターを初期化し、設定済みのEchoインスタンスを返します
func (m *MainEntryPackage) InitRouter() *echo.Echo {
	// 既存のrouter.NewRouterを利用
	e := router.NewRouter(
		m.UserController,
		m.TaskController,
		m.FeedController,
		m.ExternalAPIController,
		m.QiitaController,
		m.HatenaController,
		m.ArticleController,
		m.FeedArticleController,
		m.LayoutController,
		m.LayoutComponentController,
	)
	
	// Swaggerのエンドポイントを追加
	if m.SwaggerEnabled {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	
	return e
}