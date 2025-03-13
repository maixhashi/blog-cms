package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ILayoutComponentValidator interface {
	LayoutComponentValidate(layoutComponent model.LayoutComponent) error
}

type layoutComponentValidator struct{}

func NewLayoutComponentValidator() ILayoutComponentValidator {
	return &layoutComponentValidator{}
}

func (lcv *layoutComponentValidator) LayoutComponentValidate(layoutComponent model.LayoutComponent) error {
	return validation.ValidateStruct(&layoutComponent,
		validation.Field(&layoutComponent.Name, validation.Required.Error("名前は必須です")),
		validation.Field(&layoutComponent.Type, validation.Required.Error("タイプは必須です")),
	)
}
