package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ILayoutValidator interface {
	LayoutValidate(layout model.Layout) error
}

type layoutValidator struct{}

func NewLayoutValidator() ILayoutValidator {
	return &layoutValidator{}
}

func (lv *layoutValidator) LayoutValidate(layout model.Layout) error {
	return validation.ValidateStruct(&layout,
		validation.Field(&layout.Title, validation.Required.Error("タイトルは必須です")),
	)
}
