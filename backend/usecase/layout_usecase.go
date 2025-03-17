package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type ILayoutUsecase interface {
	GetAllLayouts(userId uint) ([]model.LayoutResponse, error)
	GetLayoutById(userId uint, layoutId uint) (model.LayoutResponse, error)
	CreateLayout(request model.LayoutRequest) (model.LayoutResponse, error)
	UpdateLayout(request model.LayoutRequest, userId uint, layoutId uint) (model.LayoutResponse, error)
	DeleteLayout(userId uint, layoutId uint) error
}

type layoutUsecase struct {
	lr repository.ILayoutRepository
	lv validator.ILayoutValidator
}

func NewLayoutUsecase(lr repository.ILayoutRepository, lv validator.ILayoutValidator) ILayoutUsecase {
	return &layoutUsecase{lr, lv}
}

func (lu *layoutUsecase) GetAllLayouts(userId uint) ([]model.LayoutResponse, error) {
	layouts, err := lu.lr.GetAllLayouts(userId)
	if err != nil {
		return nil, err
	}
	
	responses := make([]model.LayoutResponse, len(layouts))
	for i, layout := range layouts {
		responses[i] = layout.ToResponse()
	}
	return responses, nil
}

func (lu *layoutUsecase) GetLayoutById(userId uint, layoutId uint) (model.LayoutResponse, error) {
	layout, err := lu.lr.GetLayoutById(userId, layoutId)
	if err != nil {
		return model.LayoutResponse{}, err
	}
	return layout.ToResponse(), nil
}

func (lu *layoutUsecase) CreateLayout(request model.LayoutRequest) (model.LayoutResponse, error) {
	if err := lu.lv.ValidateLayoutRequest(request); err != nil {
		return model.LayoutResponse{}, err
	}
	
	layout := request.ToModel()
	if err := lu.lr.CreateLayout(&layout); err != nil {
		return model.LayoutResponse{}, err
	}
	
	return layout.ToResponse(), nil
}

func (lu *layoutUsecase) UpdateLayout(request model.LayoutRequest, userId uint, layoutId uint) (model.LayoutResponse, error) {
	if err := lu.lv.ValidateLayoutRequest(request); err != nil {
		return model.LayoutResponse{}, err
	}
	
	layout := request.ToModel()
	if err := lu.lr.UpdateLayout(&layout, userId, layoutId); err != nil {
		return model.LayoutResponse{}, err
	}
	
	return layout.ToResponse(), nil
}

func (lu *layoutUsecase) DeleteLayout(userId uint, layoutId uint) error {
	return lu.lr.DeleteLayout(userId, layoutId)
}
