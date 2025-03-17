package layout_validator_test

import (
	"go-react-app/validator"
)

var (
	layoutValidator validator.ILayoutValidator
)

func setupLayoutValidatorTest() {
	layoutValidator = validator.NewLayoutValidator()
}
