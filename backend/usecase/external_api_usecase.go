package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type IExternalAPIUsecase interface {
	GetAllExternalAPIs(userId uint) ([]model.ExternalAPIResponse, error)
	GetExternalAPIById(userId uint, apiId uint) (model.ExternalAPIResponse, error)
	CreateExternalAPI(api model.ExternalAPI) (model.ExternalAPIResponse, error)
	UpdateExternalAPI(api model.ExternalAPI, userId uint, apiId uint) (model.ExternalAPIResponse, error)
	DeleteExternalAPI(userId uint, apiId uint) error
}

type externalAPIUsecase struct {
	ar repository.IExternalAPIRepository
	av validator.IExternalAPIValidator
}

func NewExternalAPIUsecase(ar repository.IExternalAPIRepository, av validator.IExternalAPIValidator) IExternalAPIUsecase {
	return &externalAPIUsecase{ar, av}
}

func (au *externalAPIUsecase) GetAllExternalAPIs(userId uint) ([]model.ExternalAPIResponse, error) {
	apis := []model.ExternalAPI{}
	if err := au.ar.GetAllExternalAPIs(&apis, userId); err != nil {
		return nil, err
	}
	resApis := []model.ExternalAPIResponse{}
	for _, v := range apis {
		a := model.ExternalAPIResponse{
			ID:          v.ID,
			Name:        v.Name,
			BaseURL:     v.BaseURL,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		resApis = append(resApis, a)
	}
	return resApis, nil
}

func (au *externalAPIUsecase) GetExternalAPIById(userId uint, apiId uint) (model.ExternalAPIResponse, error) {
	api := model.ExternalAPI{}
	if err := au.ar.GetExternalAPIById(&api, userId, apiId); err != nil {
		return model.ExternalAPIResponse{}, err
	}
	resApi := model.ExternalAPIResponse{
		ID:          api.ID,
		Name:        api.Name,
		BaseURL:     api.BaseURL,
		Description: api.Description,
		CreatedAt:   api.CreatedAt,
		UpdatedAt:   api.UpdatedAt,
	}
	return resApi, nil
}

func (au *externalAPIUsecase) CreateExternalAPI(api model.ExternalAPI) (model.ExternalAPIResponse, error) {
	if err := au.av.ExternalAPIValidate(api); err != nil {
		return model.ExternalAPIResponse{}, err
	}
	if err := au.ar.CreateExternalAPI(&api); err != nil {
		return model.ExternalAPIResponse{}, err
	}
	resApi := model.ExternalAPIResponse{
		ID:          api.ID,
		Name:        api.Name,
		BaseURL:     api.BaseURL,
		Description: api.Description,
		CreatedAt:   api.CreatedAt,
		UpdatedAt:   api.UpdatedAt,
	}
	return resApi, nil
}

func (au *externalAPIUsecase) UpdateExternalAPI(api model.ExternalAPI, userId uint, apiId uint) (model.ExternalAPIResponse, error) {
	if err := au.av.ExternalAPIValidate(api); err != nil {
		return model.ExternalAPIResponse{}, err
	}
	if err := au.ar.UpdateExternalAPI(&api, userId, apiId); err != nil {
		return model.ExternalAPIResponse{}, err
	}
	resApi := model.ExternalAPIResponse{
		ID:          api.ID,
		Name:        api.Name,
		BaseURL:     api.BaseURL,
		Description: api.Description,
		CreatedAt:   api.CreatedAt,
		UpdatedAt:   api.UpdatedAt,
	}
	return resApi, nil
}

func (au *externalAPIUsecase) DeleteExternalAPI(userId uint, apiId uint) error {
	if err := au.ar.DeleteExternalAPI(userId, apiId); err != nil {
		return err
	}
	return nil
}
