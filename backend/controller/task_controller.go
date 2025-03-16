package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	tasksRes, err := tc.tu.GetAllTasks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("taskId")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
	}
	
	taskRes, err := tc.tu.GetTaskById(userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	var request model.TaskRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	taskRes, err := tc.tu.CreateTask(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("taskId")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
	}
	
	var request model.TaskRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	taskRes, err := tc.tu.UpdateTask(request, userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("taskId")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
	}
	
	err = tc.tu.DeleteTask(userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

