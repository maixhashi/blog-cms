package repository

import (
	"fmt"
	"go-react-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ILayoutRepository interface {
	GetAllLayouts(userId uint) ([]model.Layout, error)
	GetLayoutById(userId uint, layoutId uint) (model.Layout, error)
	CreateLayout(layout *model.Layout) error
	UpdateLayout(layout *model.Layout, userId uint, layoutId uint) error
	DeleteLayout(userId uint, layoutId uint) error
}

type layoutRepository struct {
	db *gorm.DB
}

func NewLayoutRepository(db *gorm.DB) ILayoutRepository {
	return &layoutRepository{db}
}

func (lr *layoutRepository) GetAllLayouts(userId uint) ([]model.Layout, error) {
	var layouts []model.Layout
	if err := lr.db.Where("user_id=?", userId).Order("created_at DESC").Find(&layouts).Error; err != nil {
		return nil, err
	}
	return layouts, nil
}

func (lr *layoutRepository) GetLayoutById(userId uint, layoutId uint) (model.Layout, error) {
	var layout model.Layout
	if err := lr.db.Preload("Components").Where("user_id=?", userId).First(&layout, layoutId).Error; err != nil {
		return model.Layout{}, err
	}
	return layout, nil
}

func (lr *layoutRepository) CreateLayout(layout *model.Layout) error {
	if err := lr.db.Create(layout).Error; err != nil {
		return err
	}
	return nil
}

func (lr *layoutRepository) UpdateLayout(layout *model.Layout, userId uint, layoutId uint) error {
	result := lr.db.Model(&model.Layout{}).Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", layoutId, userId).
		Updates(map[string]interface{}{
			"title": layout.Title,
		}).First(layout)
	
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("layout does not exist")
	}
	return nil
}

func (lr *layoutRepository) DeleteLayout(userId uint, layoutId uint) error {
	result := lr.db.Where("id=? AND user_id=?", layoutId, userId).Delete(&model.Layout{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("layout does not exist")
	}
	return nil
}