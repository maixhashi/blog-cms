package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type IBookUsecase interface {
	GetAllBooks(userId uint) ([]model.BookResponse, error)
	GetBookById(userId uint, bookId uint) (model.BookResponse, error)
	CreateBook(request model.BookRequest) (model.BookResponse, error)
	UpdateBook(request model.BookRequest, userId uint, bookId uint) (model.BookResponse, error)
	DeleteBook(userId uint, bookId uint) error
}

type bookUsecase struct {
	br repository.IBookRepository
	bv validator.IBookValidator
}

func NewBookUsecase(br repository.IBookRepository, bv validator.IBookValidator) IBookUsecase {
	return &bookUsecase{br, bv}
}

func (bu *bookUsecase) GetAllBooks(userId uint) ([]model.BookResponse, error) {
	books, err := bu.br.GetAllBooks(userId)
	if err != nil {
		return nil, err
	}
	
	responses := make([]model.BookResponse, len(books))
	for i, book := range books {
		responses[i] = book.ToResponse()
	}
	return responses, nil
}

func (bu *bookUsecase) GetBookById(userId uint, bookId uint) (model.BookResponse, error) {
	book, err := bu.br.GetBookById(userId, bookId)
	if err != nil {
		return model.BookResponse{}, err
	}
	return book.ToResponse(), nil
}

func (bu *bookUsecase) CreateBook(request model.BookRequest) (model.BookResponse, error) {
	if err := bu.bv.ValidateBookRequest(request); err != nil {
		return model.BookResponse{}, err
	}
	
	book := request.ToModel()
	if err := bu.br.CreateBook(&book); err != nil {
		return model.BookResponse{}, err
	}
	
	return book.ToResponse(), nil
}

func (bu *bookUsecase) UpdateBook(request model.BookRequest, userId uint, bookId uint) (model.BookResponse, error) {
	if err := bu.bv.ValidateBookRequest(request); err != nil {
		return model.BookResponse{}, err
	}
	
	book := request.ToModel()
	if err := bu.br.UpdateBook(&book, userId, bookId); err != nil {
		return model.BookResponse{}, err
	}
	
	return book.ToResponse(), nil
}

func (bu *bookUsecase) DeleteBook(userId uint, bookId uint) error {
	return bu.br.DeleteBook(userId, bookId)
}
