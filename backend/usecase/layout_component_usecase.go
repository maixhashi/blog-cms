package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type ILayoutComponentUsecase interface {
	GetAllLayoutComponents(userId uint) ([]model.LayoutComponentResponse, error)
	GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponentResponse, error)
	CreateLayoutComponent(component model.LayoutComponent) (model.LayoutComponentResponse, error)
	UpdateLayoutComponent(component model.LayoutComponent, userId uint, componentId uint) (model.LayoutComponentResponse, error)
	DeleteLayoutComponent(userId uint, componentId uint) error
}

type layoutComponentUsecase struct {
	lcr repository.ILayoutComponentRepository
	lcv validator.ILayoutComponentValidator
}

func NewLayoutComponentUsecase(lcr repository.ILayoutComponentRepository, lcv validator.ILayoutComponentValidator) ILayoutComponentUsecase {
	return &layoutComponentUsecase{lcr, lcv}
}

func (lcu *layoutComponentUsecase) GetAllLayoutComponents(userId uint) ([]model.LayoutComponentResponse, error) {
	components := []model.LayoutComponent{}
	if err := lcu.lcr.GetAllLayoutComponents(&components, userId); err != nil {
		return nil, err
	}
	resComponents := []model.LayoutComponentResponse{}
	for _, v := range components {
		c := model.LayoutComponentResponse{
			ID:        v.ID,
			Name:      v.Name,
			Type:      v.Type,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resComponents = append(resComponents, c)
	}
	return resComponents, nil
}

func (lcu *layoutComponentUsecase) GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponentResponse, error) {
	component := model.LayoutComponent{}
	if err := lcu.lcr.GetLayoutComponentById(&component, userId, componentId); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	resComponent := model.LayoutComponentResponse{
		ID:        component.ID,
		Name:      component.Name,
		Type:      component.Type,
		Content:   component.Content,
		CreatedAt: component.CreatedAt,
		UpdatedAt: component.UpdatedAt,
	}
	return resComponent, nil
}

func (lcu *layoutComponentUsecase) CreateLayoutComponent(component model.LayoutComponent) (model.LayoutComponentResponse, error) {
	if err := lcu.lcv.LayoutComponentValidate(component); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	if err := lcu.lcr.CreateLayoutComponent(&component); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	resComponent := model.LayoutComponentResponse{
		ID:        component.ID,
		Name:      component.Name,
		Type:      component.Type,
		Content:   component.Content,
		CreatedAt: component.CreatedAt,
		UpdatedAt: component.UpdatedAt,
	}
	return resComponent, nil
}

func (lcu *layoutComponentUsecase) UpdateLayoutComponent(component model.LayoutComponent, userId uint, componentId uint) (model.LayoutComponentResponse, error) {
	if err := lcu.lcv.LayoutComponentValidate(component); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	if err := lcu.lcr.UpdateLayoutComponent(&component, userId, componentId); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	resComponent := model.LayoutComponentResponse{
		ID:        component.ID,
		Name:      component.Name,
		Type:      component.Type,
		Content:   component.Content,
		CreatedAt: component.CreatedAt,
		UpdatedAt: component.UpdatedAt,
	}
	return resComponent, nil
}

func (lcu *layoutComponentUsecase) DeleteLayoutComponent(userId uint, componentId uint) error {
	if err := lcu.lcr.DeleteLayoutComponent(userId, componentId); err != nil {
		return err
	}
	return nil
}
