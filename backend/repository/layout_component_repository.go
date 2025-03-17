package repository

import (
	"errors"
	"fmt"
	"go-react-app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ILayoutComponentRepository interface {
	GetAllLayoutComponents(userId uint) ([]model.LayoutComponent, error)
	GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponent, error)
	CreateLayoutComponent(component *model.LayoutComponent) error
	UpdateLayoutComponent(component *model.LayoutComponent, userId uint, componentId uint) error
	DeleteLayoutComponent(userId uint, componentId uint) error
	AssignToLayout(componentId uint, layoutId uint, userId uint, position model.PositionRequest) error
	RemoveFromLayout(componentId uint, userId uint) error
	UpdatePosition(componentId uint, userId uint, position model.PositionRequest) error
}

type layoutComponentRepository struct {
	db *gorm.DB
}

func NewLayoutComponentRepository(db *gorm.DB) ILayoutComponentRepository {
	return &layoutComponentRepository{db}
}

func (lcr *layoutComponentRepository) GetAllLayoutComponents(userId uint) ([]model.LayoutComponent, error) {
	var components []model.LayoutComponent
	if err := lcr.db.Where("user_id=?", userId).Order("created_at DESC").Find(&components).Error; err != nil {
		return nil, err
	}
	return components, nil
}

func (lcr *layoutComponentRepository) GetLayoutComponentById(userId uint, componentId uint) (model.LayoutComponent, error) {
	var component model.LayoutComponent
	if err := lcr.db.Where("user_id=?", userId).First(&component, componentId).Error; err != nil {
		return model.LayoutComponent{}, err
	}
	return component, nil
}

func (lcr *layoutComponentRepository) CreateLayoutComponent(component *model.LayoutComponent) error {
	if err := lcr.db.Create(component).Error; err != nil {
		return err
	}
	return nil
}

func (lcr *layoutComponentRepository) UpdateLayoutComponent(component *model.LayoutComponent, userId uint, componentId uint) error {
	result := lcr.db.Model(&model.LayoutComponent{}).Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", componentId, userId).
		Updates(map[string]interface{}{
			"name":    component.Name,
			"type":    component.Type,
			"content": component.Content,
		}).First(component)
	
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

func (lcr *layoutComponentRepository) AssignToLayout(componentId uint, layoutId uint, userId uint, position model.PositionRequest) error {
    // まずコンポーネントを取得
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
    component.X = position.X
    component.Y = position.Y
    
    // 幅と高さを設定（もし提供されていれば）
    if position.Width != 0 {
        component.Width = position.Width
    }
    if position.Height != 0 {
        component.Height = position.Height
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

func (lcr *layoutComponentRepository) UpdatePosition(componentId uint, userId uint, position model.PositionRequest) error {
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
    component.X = position.X
    component.Y = position.Y
    
    // 幅と高さを更新（もし提供されていれば）
    if position.Width != 0 {
        component.Width = position.Width
    }
    if position.Height != 0 {
        component.Height = position.Height
    }
    
    // 保存
    if err := lcr.db.Save(&component).Error; err != nil {
        return err
    }
    
    return nil
}
