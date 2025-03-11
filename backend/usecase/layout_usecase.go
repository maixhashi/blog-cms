package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type ILayoutUsecase interface {
	GetAllLayouts(userId uint) ([]model.LayoutResponse, error)
	GetLayoutById(userId uint, layoutId uint) (model.LayoutResponse, error)
	CreateLayout(layout model.Layout) (model.LayoutResponse, error)
	UpdateLayout(layout model.Layout, userId uint, layoutId uint) (model.LayoutResponse, error)
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
	layouts := []model.Layout{}
	if err := lu.lr.GetAllLayouts(&layouts, userId); err != nil {
		return nil, err
	}
	resLayouts := []model.LayoutResponse{}
	for _, v := range layouts {
		l := model.LayoutResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resLayouts = append(resLayouts, l)
	}
	return resLayouts, nil
}

func (lu *layoutUsecase) GetLayoutById(userId uint, layoutId uint) (model.LayoutResponse, error) {
	layout := model.Layout{}
	if err := lu.lr.GetLayoutById(&layout, userId, layoutId); err != nil {
		return model.LayoutResponse{}, err
	}
	resLayout := model.LayoutResponse{
		ID:        layout.ID,
		Title:     layout.Title,
		CreatedAt: layout.CreatedAt,
		UpdatedAt: layout.UpdatedAt,
	}
	return resLayout, nil
}

func (lu *layoutUsecase) CreateLayout(layout model.Layout) (model.LayoutResponse, error) {
	if err := lu.lv.LayoutValidate(layout); err != nil {
		return model.LayoutResponse{}, err
	}
	if err := lu.lr.CreateLayout(&layout); err != nil {
		return model.LayoutResponse{}, err
	}
	resLayout := model.LayoutResponse{
		ID:        layout.ID,
		Title:     layout.Title,
		CreatedAt: layout.CreatedAt,
		UpdatedAt: layout.UpdatedAt,
	}
	return resLayout, nil
}

func (lu *layoutUsecase) UpdateLayout(layout model.Layout, userId uint, layoutId uint) (model.LayoutResponse, error) {
	if err := lu.lv.LayoutValidate(layout); err != nil {
		return model.LayoutResponse{}, err
	}
	if err := lu.lr.UpdateLayout(&layout, userId, layoutId); err != nil {
		return model.LayoutResponse{}, err
	}
	resLayout := model.LayoutResponse{
		ID:        layout.ID,
		Title:     layout.Title,
		CreatedAt: layout.CreatedAt,
		UpdatedAt: layout.UpdatedAt,
	}
	return resLayout, nil
}

func (lu *layoutUsecase) DeleteLayout(userId uint, layoutId uint) error {
	if err := lu.lr.DeleteLayout(userId, layoutId); err != nil {
		return err
	}
	return nil
}
