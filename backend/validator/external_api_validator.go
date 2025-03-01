package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IExternalAPIValidator interface {
	ExternalAPIValidate(api model.ExternalAPI) error
}

type externalAPIValidator struct{}

func NewExternalAPIValidator() IExternalAPIValidator {
	return &externalAPIValidator{}
}

func (av *externalAPIValidator) ExternalAPIValidate(api model.ExternalAPI) error {
	return validation.ValidateStruct(&api,
		validation.Field(
			&api.Name,
			validation.Required.Error("name is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
	)
}
