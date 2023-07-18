package facades

import (
	"github.com/mewway/go-laravel/contracts/console"
)

func Artisan() console.Artisan {
	return App().MakeArtisan()
}
