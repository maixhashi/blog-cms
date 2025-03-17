package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ILayoutValidator interface {
	ValidateLayoutRequest(layout model.LayoutRequest) error
}

type layoutValidator struct{}

func NewLayoutValidator() ILayoutValidator {
	return &layoutValidator{}
}

func (lv *layoutValidator) ValidateLayoutRequest(layout model.LayoutRequest) error {
	return validation.ValidateStruct(&layout,
		validation.Field(&layout.Title, validation.Required.Error("タイトルは必須です")),
	)
}
