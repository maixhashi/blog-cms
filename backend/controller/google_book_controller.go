package controller

import (
	"go-react-app/model"
	"go-react-app/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IGoogleBookController interface {
	SearchBooks(c echo.Context) error
	GetBookByID(c echo.Context) error
	ImportBookFromGoogle(c echo.Context) error
}

type googleBookController struct {
	gbu usecase.IGoogleBookUsecase
	bu  usecase.IBookUsecase
}

func NewGoogleBookController(gbu usecase.IGoogleBookUsecase, bu usecase.IBookUsecase) IGoogleBookController {
	return &googleBookController{gbu, bu}
}

// SearchBooks Google Books APIで書籍を検索
// @Summary 書籍を検索
// @Description Google Books APIを使用して書籍を検索する
// @Tags google-books
// @Accept json
// @Produce json
// @Param request body model.GoogleBookSearchRequest true "検索条件"
// @Success 200 {object} model.GoogleBookSearchResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /google-books/search [post]
func (gbc *googleBookController) SearchBooks(c echo.Context) error {
	var request model.GoogleBookSearchRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	result, err := gbc.gbu.SearchBooks(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}

// GetBookByID Google Books APIで特定の書籍を取得
// @Summary 特定の書籍を取得
// @Description Google Books APIを使用して特定のIDの書籍を取得する
// @Tags google-books
// @Accept json
// @Produce json
// @Param id path string true "Google Books ID"
// @Success 200 {object} model.GoogleBook
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /google-books/{id} [get]
func (gbc *googleBookController) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "book ID is required"})
	}
	
	book, err := gbc.gbu.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, book)
}

// ImportBookFromGoogle Google Books APIから書籍をインポート
// @Summary 書籍をインポート
// @Description Google Books APIから取得した書籍をユーザーの蔵書に追加する
// @Tags google-books
// @Accept json
// @Produce json
// @Param id path string true "Google Books ID"
// @Success 201 {object} model.BookResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /google-books/{id}/import [post]
func (gbc *googleBookController) ImportBookFromGoogle(c echo.Context) error {
	userId := getUserIdFromToken(c)
	
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "book ID is required"})
	}
	
	// Google Books APIから書籍情報を取得
	googleBook, err := gbc.gbu.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	// BookRequestに変換
	bookRequest := googleBook.ToBookRequest()
	bookRequest.UserId = userId
	
	// 書籍を作成
	bookRes, err := gbc.bu.CreateBook(bookRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusCreated, bookRes)
}