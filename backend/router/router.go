package router

import (
	"go-react-app/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, fc controller.IFeedController, ac controller.IExternalAPIController, qc controller.IQiitaController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
	}))

	// ユーザー認証関連
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	// タスク関連のルート
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	// フィード関連のルート
	f := e.Group("/feeds")
	f.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	f.GET("", fc.GetAllFeeds)
	f.GET("/:feedId", fc.GetFeedById)
	f.POST("", fc.CreateFeed)
	f.PUT("/:feedId", fc.UpdateFeed)
	f.DELETE("/:feedId", fc.DeleteFeed)

	// 外部API関連のルート
	a := e.Group("/external-apis")
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	a.GET("", ac.GetAllExternalAPIs)
	a.GET("/:apiId", ac.GetExternalAPIById)
	a.POST("", ac.CreateExternalAPI)
	a.PUT("/:apiId", ac.UpdateExternalAPI)
	a.DELETE("/:apiId", ac.DeleteExternalAPI)
	
	// 外部API関連のルート
	q := e.Group("/qiita")
	q.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	q.GET("/articles", qc.GetQiitaArticles)
	q.GET("/articles/:id", qc.GetQiitaArticleByID)

	return e
}

