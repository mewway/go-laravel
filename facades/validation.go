package facades

import (
	"github.com/mewway/go-laravel/contracts/validation"
)

func Validation() validation.Validation {
	return App().MakeValidation()
}
