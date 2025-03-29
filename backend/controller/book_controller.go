package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IBookController interface {
	GetAllBooks(c echo.Context) error
	GetBookById(c echo.Context) error
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

type bookController struct {
	bu usecase.IBookUsecase
}

func NewBookController(bu usecase.IBookUsecase) IBookController {
	return &bookController{bu}
}

// GetAllBooks ユーザーのすべての書籍を取得
// @Summary ユーザーの書籍一覧を取得
// @Description ログインユーザーのすべての書籍を取得する
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} model.BookResponse
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (bc *bookController) GetAllBooks(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	booksRes, err := bc.bu.GetAllBooks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, booksRes)
}

// GetBookById 指定されたIDの書籍を取得
// @Summary 特定の書籍を取得
// @Description 指定されたIDの書籍を取得する
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "書籍ID"
// @Success 200 {object} model.BookResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{bookId} [get]
func (bc *bookController) GetBookById(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("bookId")
	bookId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}
	
	bookRes, err := bc.bu.GetBookById(userId, uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, bookRes)
}

// CreateBook 新しい書籍を作成
// @Summary 新しい書籍を作成
// @Description ユーザーの新しい書籍を作成する
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.BookRequest true "書籍情報"
// @Success 201 {object} model.BookResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (bc *bookController) CreateBook(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	var request model.BookRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	bookRes, err := bc.bu.CreateBook(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, bookRes)
}

// UpdateBook 既存の書籍を更新
// @Summary 書籍を更新
// @Description 指定されたIDの書籍を更新する
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "書籍ID"
// @Param book body model.BookRequest true "更新する書籍情報"
// @Success 200 {object} model.BookResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{bookId} [put]
func (bc *bookController) UpdateBook(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("bookId")
	bookId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}
	
	var request model.BookRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	bookRes, err := bc.bu.UpdateBook(request, userId, uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, bookRes)
}

// DeleteBook 書籍を削除
// @Summary 書籍を削除
// @Description 指定されたIDの書籍を削除する
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "書籍ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{bookId} [delete]
func (bc *bookController) DeleteBook(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("bookId")
	bookId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}
	
	err = bc.bu.DeleteBook(userId, uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
