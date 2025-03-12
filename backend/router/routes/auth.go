package routes

import (
	"go-react-app/controller"
	"github.com/labstack/echo/v4"
)

// SetupAuthRoutes は認証関連のルートを設定します
func SetupAuthRoutes(e *echo.Echo, uc controller.IUserController) {
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
}
