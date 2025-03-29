package router

import (
	"go-react-app/controller"
	"github.com/labstack/echo/v4"
	"go-react-app/router/routes"
)

func NewRouter(
	uc controller.IUserController,
	tc controller.ITaskController,
	fc controller.IFeedController,
	ac controller.IExternalAPIController,
	qc controller.IQiitaController,
	hc controller.IHatenaController,
	artc controller.IArticleController,
	fac controller.IFeedArticleController,
	lc controller.ILayoutController,
	lcc controller.ILayoutComponentController,
	bc controller.IBookController,
	gbc controller.IGoogleBookController) *echo.Echo {
	
	e := echo.New()
	
	// ミドルウェアの設定
	setupMiddleware(e)
	
	// 各種ルートの設定
	routes.SetupAuthRoutes(e, uc)
	routes.SetupTaskRoutes(e, tc)
	routes.SetupFeedRoutes(e, fc)
	routes.SetupExternalAPIRoutes(e, ac)
	routes.SetupQiitaRoutes(e, qc)
	routes.SetupHatenaRoutes(e, hc)
	routes.SetupArticleRoutes(e, artc)
	routes.SetupFeedArticleRoutes(e, fac)
	routes.SetupLayoutRoutes(e, lc)
	routes.SetupLayoutComponentRoutes(e, lcc)
	routes.SetupBookRoutes(e, bc)
	routes.SetupGoogleBookRoutes(e, gbc)
	
	return e
}