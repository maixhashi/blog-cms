package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type IGoogleBookUsecase interface {
	SearchBooks(request model.GoogleBookSearchRequest) (model.GoogleBookSearchResponse, error)
	GetBookByID(id string) (model.GoogleBook, error)
}

type googleBookUsecase struct {
	gbr repository.IGoogleBookRepository
	bv  validator.IBookValidator
}

func NewGoogleBookUsecase(gbr repository.IGoogleBookRepository, bv validator.IBookValidator) IGoogleBookUsecase {
	return &googleBookUsecase{gbr, bv}
}

func (gbu *googleBookUsecase) SearchBooks(request model.GoogleBookSearchRequest) (model.GoogleBookSearchResponse, error) {
	if err := gbu.bv.ValidateGoogleBookSearchRequest(request); err != nil {
		return model.GoogleBookSearchResponse{}, err
	}
	
	return gbu.gbr.SearchBooks(request.Query, request.MaxResults)
}

func (gbu *googleBookUsecase) GetBookByID(id string) (model.GoogleBook, error) {
	return gbu.gbr.GetBookByID(id)
}
