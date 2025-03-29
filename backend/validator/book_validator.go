package validator

import (
	"fmt"
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IBookValidator interface {
	ValidateBookRequest(book model.BookRequest) error
	ValidateGoogleBookSearchRequest(request model.GoogleBookSearchRequest) error
}

type bookValidator struct{}

func NewBookValidator() IBookValidator {
	return &bookValidator{}
}

func (bv *bookValidator) ValidateBookRequest(book model.BookRequest) error {
	return validation.ValidateStruct(&book,
		validation.Field(
			&book.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, model.BookTitleMaxLength).Error(
				fmt.Sprintf("limited max %d char", model.BookTitleMaxLength),
			),
		),
		validation.Field(
			&book.Author,
			validation.Required.Error("author is required"),
			validation.RuneLength(1, model.BookAuthorMaxLength).Error(
				fmt.Sprintf("limited max %d char", model.BookAuthorMaxLength),
			),
		),
	)
}

func (bv *bookValidator) ValidateGoogleBookSearchRequest(request model.GoogleBookSearchRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Query,
			validation.Required.Error("query is required"),
		),
	)
}
