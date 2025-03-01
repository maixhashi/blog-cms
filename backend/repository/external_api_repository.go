package repository

import (
	"fmt"
	"go-react-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IExternalAPIRepository interface {
	GetAllExternalAPIs(apis *[]model.ExternalAPI, userId uint) error
	GetExternalAPIById(api *model.ExternalAPI, userId uint, apiId uint) error
	CreateExternalAPI(api *model.ExternalAPI) error
	UpdateExternalAPI(api *model.ExternalAPI, userId uint, apiId uint) error
	DeleteExternalAPI(userId uint, apiId uint) error
}

type externalAPIRepository struct {
	db *gorm.DB
}

func NewExternalAPIRepository(db *gorm.DB) IExternalAPIRepository {
	return &externalAPIRepository{db}
}

func (ar *externalAPIRepository) GetAllExternalAPIs(apis *[]model.ExternalAPI, userId uint) error {
	if err := ar.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(apis).Error; err != nil {
		return err
	}
	return nil
}

func (ar *externalAPIRepository) GetExternalAPIById(api *model.ExternalAPI, userId uint, apiId uint) error {
	if err := ar.db.Joins("User").Where("user_id=?", userId).First(api, apiId).Error; err != nil {
		return err
	}
	return nil
}

func (ar *externalAPIRepository) CreateExternalAPI(api *model.ExternalAPI) error {
	if err := ar.db.Create(api).Error; err != nil {
		return err
	}
	return nil
}

func (ar *externalAPIRepository) UpdateExternalAPI(api *model.ExternalAPI, userId uint, apiId uint) error {
	result := ar.db.Model(api).Clauses(clause.Returning{}).Where("id=? AND user_id=?", apiId, userId).Updates(map[string]interface{}{
		"name":        api.Name,
		"base_url":    api.BaseURL,
		"description": api.Description,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (ar *externalAPIRepository) DeleteExternalAPI(userId uint, apiId uint) error {
	result := ar.db.Where("id=? AND user_id=?", apiId, userId).Delete(&model.ExternalAPI{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
