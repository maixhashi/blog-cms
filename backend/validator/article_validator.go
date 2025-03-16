package validator

import (
	"go-react-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IArticleValidator interface {
	ValidateArticleRequest(article model.ArticleRequest) error
}

type articleValidator struct{}

func NewArticleValidator() IArticleValidator {
	return &articleValidator{}
}

func (av *articleValidator) ValidateArticleRequest(article model.ArticleRequest) error {
	return validation.ValidateStruct(&article,
		validation.Field(&article.Title, validation.Required.Error("タイトルは必須です")),
	)
}
