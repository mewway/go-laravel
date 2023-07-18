package facades

import (
	"github.com/mewway/go-laravel/contracts/crypt"
)

func Crypt() crypt.Crypt {
	return App().MakeCrypt()
}
