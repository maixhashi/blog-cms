package routes

import (
	"go-react-app/controller"
	"go-react-app/utils/middleware"
	"github.com/labstack/echo/v4"
)

// SetupTaskRoutes はタスク関連のルートを設定します
func SetupTaskRoutes(e *echo.Echo, tc controller.ITaskController) {
	t := e.Group("/tasks")
	t.Use(middleware.GetJWTMiddleware())
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
}