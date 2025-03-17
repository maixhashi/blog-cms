package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type ILayoutComponentUsecase interface {
	GetAllLayoutComponents(userId uint) ([]model.LayoutComponentResponse, error)
	GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponentResponse, error)
	CreateLayoutComponent(request model.LayoutComponentRequest) (model.LayoutComponentResponse, error)
	UpdateLayoutComponent(request model.LayoutComponentRequest, userId uint, componentId uint) (model.LayoutComponentResponse, error)
	DeleteLayoutComponent(userId uint, componentId uint) error
	
	AssignToLayout(userId uint, componentId uint, request model.AssignLayoutRequest) error
	RemoveFromLayout(userId uint, componentId uint) error
	UpdatePosition(userId uint, componentId uint, position model.PositionRequest) error
}

type layoutComponentUsecase struct {
	lcr repository.ILayoutComponentRepository
	lcv validator.ILayoutComponentValidator
}

func NewLayoutComponentUsecase(lcr repository.ILayoutComponentRepository, lcv validator.ILayoutComponentValidator) ILayoutComponentUsecase {
	return &layoutComponentUsecase{lcr, lcv}
}

func (lcu *layoutComponentUsecase) GetAllLayoutComponents(userId uint) ([]model.LayoutComponentResponse, error) {
	components, err := lcu.lcr.GetAllLayoutComponents(userId)
	if err != nil {
		return nil, err
	}
	
	responses := make([]model.LayoutComponentResponse, len(components))
	for i, component := range components {
		responses[i] = component.ToResponse()
	}
	return responses, nil
}

func (lcu *layoutComponentUsecase) GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponentResponse, error) {
	component, err := lcu.lcr.GetLayoutComponentById(userId, componentId)
	if err != nil {
		return model.LayoutComponentResponse{}, err
	}
	return component.ToResponse(), nil
}

func (lcu *layoutComponentUsecase) CreateLayoutComponent(request model.LayoutComponentRequest) (model.LayoutComponentResponse, error) {
	if err := lcu.lcv.ValidateLayoutComponentRequest(request); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	
	component := request.ToModel()
	if err := lcu.lcr.CreateLayoutComponent(&component); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	
	return component.ToResponse(), nil
}

func (lcu *layoutComponentUsecase) UpdateLayoutComponent(request model.LayoutComponentRequest, userId uint, componentId uint) (model.LayoutComponentResponse, error) {
	if err := lcu.lcv.ValidateLayoutComponentRequest(request); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	
	component := request.ToModel()
	if err := lcu.lcr.UpdateLayoutComponent(&component, userId, componentId); err != nil {
		return model.LayoutComponentResponse{}, err
	}
	
	return component.ToResponse(), nil
}

func (lcu *layoutComponentUsecase) DeleteLayoutComponent(userId uint, componentId uint) error {
	return lcu.lcr.DeleteLayoutComponent(userId, componentId)
}

func (lcu *layoutComponentUsecase) AssignToLayout(userId uint, componentId uint, request model.AssignLayoutRequest) error {
	return lcu.lcr.AssignToLayout(componentId, request.LayoutId, userId, request.Position)
}

func (lcu *layoutComponentUsecase) RemoveFromLayout(userId uint, componentId uint) error {
	return lcu.lcr.RemoveFromLayout(componentId, userId)
}

func (lcu *layoutComponentUsecase) UpdatePosition(userId uint, componentId uint, position model.PositionRequest) error {
	return lcu.lcr.UpdatePosition(componentId, userId, position)
}
