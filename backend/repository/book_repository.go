package repository

import (
	"fmt"
	"go-react-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBookRepository interface {
	GetAllBooks(userId uint) ([]model.Book, error)
	GetBookById(userId uint, bookId uint) (model.Book, error)
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book, userId uint, bookId uint) error
	DeleteBook(userId uint, bookId uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) GetAllBooks(userId uint) ([]model.Book, error) {
	var books []model.Book
	if err := br.db.Where("user_id=?", userId).Order("created_at").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (br *bookRepository) GetBookById(userId uint, bookId uint) (model.Book, error) {
	var book model.Book
	result := br.db.Where("user_id=?", userId).First(&book, bookId)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Book{}, fmt.Errorf("book not found")
	}
	return book, nil
}

func (br *bookRepository) CreateBook(book *model.Book) error {
	return br.db.Create(book).Error
}

func (br *bookRepository) UpdateBook(book *model.Book, userId uint, bookId uint) error {
	result := br.db.Model(&model.Book{}).Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", bookId, userId).
		Updates(map[string]interface{}{
			"title":         book.Title,
			"author":        book.Author,
			"description":   book.Description,
			"isbn":          book.ISBN,
			"image_url":     book.ImageURL,
			"published_date": book.PublishedDate,
		}).First(book)
	
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (br *bookRepository) DeleteBook(userId uint, bookId uint) error {
	result := br.db.Where("id=? AND user_id=?", bookId, userId).Delete(&model.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
