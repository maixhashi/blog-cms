package repository

import (
	"fmt"
	"errors"
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
    // まずコンポーネントを取得（順序を変更）
    var component model.LayoutComponent
    if err := lcr.db.Where("id = ? AND user_id = ?", componentId, userId).First(&component).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("layout component does not exist")
        }
        return err
    }
    
    // 次にレイアウトが存在するか確認
    var layout model.Layout
    if err := lcr.db.Where("id = ? AND user_id = ?", layoutId, userId).First(&layout).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("layout does not exist")
        }
        return err
    }
    
    // レイアウトに割り当て
    component.LayoutId = &layoutId
    
    // 位置情報を設定
    component.X = position["x"]
    component.Y = position["y"]
    
    // 幅と高さを設定（もし提供されていれば）
    if width, ok := position["width"]; ok {
        component.Width = width
    }
    if height, ok := position["height"]; ok {
        component.Height = height
    }
    
    // 保存
    if err := lcr.db.Save(&component).Error; err != nil {
        return err
    }
    
    return nil
}
func (lcr *layoutComponentRepository) RemoveFromLayout(componentId uint, userId uint) error {
    var component model.LayoutComponent
    if err := lcr.db.Where("id = ? AND user_id = ?", componentId, userId).First(&component).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("layout component does not exist")
        }
        return err
    }
    
    // レイアウトIDをnilに設定
    component.LayoutId = nil
    
    // 位置情報をリセット
    component.X = 0
    component.Y = 0
    
    // 保存
    if err := lcr.db.Save(&component).Error; err != nil {
        return err
    }
    
    return nil
}

func (lcr *layoutComponentRepository) UpdatePosition(componentId uint, userId uint, position map[string]int) error {
    var component model.LayoutComponent
    if err := lcr.db.Where("id = ? AND user_id = ?", componentId, userId).First(&component).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("layout component does not exist")
        }
        return err
    }
    
    // レイアウトに割り当てられていない場合はエラー
    if component.LayoutId == nil {
        return errors.New("component is not assigned to any layout")
    }
    
    // 位置情報を更新
    component.X = position["x"]
    component.Y = position["y"]
    
    // 幅と高さを更新（もし提供されていれば）
    if width, ok := position["width"]; ok {
        component.Width = width
    }
    if height, ok := position["height"]; ok {
        component.Height = height
    }
    
    // 保存
    if err := lcr.db.Save(&component).Error; err != nil {
        return err
    }
    
    return nil
}
