package controller

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTトークンからユーザーIDを取得するヘルパー関数
func getUserIdFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return uint(claims["user_id"].(float64))
}
