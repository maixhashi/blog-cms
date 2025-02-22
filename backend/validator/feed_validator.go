package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IFeedValidator interface {
	FeedValidate(feed model.Feed) error
}

type FeedValidator struct{}

func NewFeedValidator() IFeedValidator {
	return &FeedValidator{}
}

func (tv *FeedValidator) FeedValidate(feed model.Feed) error {
	return validation.ValidateStruct(&feed,
		validation.Field(
			&feed.Title,
			validation.Required.Error("title is required"),
		),
	)
}