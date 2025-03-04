package validator

import (
	"fmt"
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(model.UserEmailMinLength, model.UserEmailMaxLength).Error("limited min "+
				 fmt.Sprintf("%d", model.UserEmailMinLength) + 
				 " max " + fmt.Sprintf("%d", model.UserEmailMaxLength) + " char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(model.UserPasswordMinLength, model.UserPasswordMaxLength).Error("limited min " + 
				fmt.Sprintf("%d", model.UserPasswordMinLength) + 
				" max " + fmt.Sprintf("%d", model.UserPasswordMaxLength) + " char"),
		),
	)
}