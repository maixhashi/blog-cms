package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ILayoutComponentValidator interface {
	ValidateLayoutComponentRequest(component model.LayoutComponentRequest) error
	ValidateAssignLayoutRequest(request model.AssignLayoutRequest) error
	ValidatePositionRequest(position model.PositionRequest) error
}

type layoutComponentValidator struct{}

func NewLayoutComponentValidator() ILayoutComponentValidator {
	return &layoutComponentValidator{}
}

func (lcv *layoutComponentValidator) ValidateLayoutComponentRequest(component model.LayoutComponentRequest) error {
	return validation.ValidateStruct(&component,
		validation.Field(&component.Name, validation.Required.Error("名前は必須です")),
		validation.Field(&component.Type, validation.Required.Error("タイプは必須です")),
	)
}

func (lcv *layoutComponentValidator) ValidateAssignLayoutRequest(request model.AssignLayoutRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.LayoutId, validation.Required.Error("レイアウトIDは必須です")),
	)
}

func (lcv *layoutComponentValidator) ValidatePositionRequest(position model.PositionRequest) error {
	// 位置情報のバリデーションが必要な場合はここに追加
	return nil
}
