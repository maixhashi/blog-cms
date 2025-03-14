package repository

import (
	"fmt"
	"go-react-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ILayoutComponentRepository interface {
	GetAllLayoutComponents(components *[]model.LayoutComponent, userId uint) error
	GetLayoutComponentById(component *model.LayoutComponent, userId uint, componentId uint) error
	CreateLayoutComponent(component *model.LayoutComponent) error
	UpdateLayoutComponent(component *model.LayoutComponent, userId uint, componentId uint) error
	DeleteLayoutComponent(userId uint, componentId uint) error
	AssignToLayout(componentId uint, layoutId uint, userId uint, position map[string]int) error
	RemoveFromLayout(componentId uint, userId uint) error
	UpdatePosition(componentId uint, userId uint, position map[string]int) error
}

type layoutComponentRepository struct {
	db *gorm.DB
}

func NewLayoutComponentRepository(db *gorm.DB) ILayoutComponentRepository {
	return &layoutComponentRepository{db}
}

func (lcr *layoutComponentRepository) GetAllLayoutComponents(components *[]model.LayoutComponent, userId uint) error {
	if err := lcr.db.Joins("User").Where("user_id=?", userId).Order("layout_components.created_at DESC").Find(components).Error; err != nil {
		return err
	}
	return nil
}

func (lcr *layoutComponentRepository) GetLayoutComponentById(component *model.LayoutComponent, userId uint, componentId uint) error {
	if err := lcr.db.Joins("User").Where("user_id=?", userId).First(component, componentId).Error; err != nil {
		return err
	}
	return nil
}

func (lcr *layoutComponentRepository) CreateLayoutComponent(component *model.LayoutComponent) error {
	if err := lcr.db.Create(component).Error; err != nil {
		return err
	}
	return nil
}

func (lcr *layoutComponentRepository) UpdateLayoutComponent(component *model.LayoutComponent, userId uint, componentId uint) error {
	result := lcr.db.Model(component).Clauses(clause.Returning{}).Where("id=? AND user_id=?", componentId, userId).Updates(map[string]interface{}{
		"name":    component.Name,
		"type":    component.Type,
		"content": component.Content,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("layout component does not exist")
	}
	return nil
}

func (lcr *layoutComponentRepository) DeleteLayoutComponent(userId uint, componentId uint) error {
	result := lcr.db.Where("id=? AND user_id=?", componentId, userId).Delete(&model.LayoutComponent{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("layout component does not exist")
	}
	return nil
}

func (lcr *layoutComponentRepository) AssignToLayout(componentId uint, layoutId uint, userId uint, position map[string]int) error {
    result := lcr.db.Model(&model.LayoutComponent{}).
        Where("id = ? AND user_id = ?", componentId, userId).
        Updates(map[string]interface{}{
            "layout_id": layoutId,
            "x": position["x"],
            "y": position["y"],
            "width": position["width"],
            "height": position["height"],
        })
    
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected < 1 {
        return fmt.Errorf("layout component does not exist or already assigned")
    }
    return nil
}

func (lcr *layoutComponentRepository) RemoveFromLayout(componentId uint, userId uint) error {
    result := lcr.db.Model(&model.LayoutComponent{}).
        Where("id = ? AND user_id = ?", componentId, userId).
        Update("layout_id", nil)
    
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected < 1 {
        return fmt.Errorf("layout component does not exist")
    }
    return nil
}

func (lcr *layoutComponentRepository) UpdatePosition(componentId uint, userId uint, position map[string]int) error {
    result := lcr.db.Model(&model.LayoutComponent{}).
        Where("id = ? AND user_id = ?", componentId, userId).
        Updates(map[string]interface{}{
            "x": position["x"],
            "y": position["y"],
            "width": position["width"],
            "height": position["height"],
        })
    
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected < 1 {
        return fmt.Errorf("layout component does not exist")
    }
    return nil
}
