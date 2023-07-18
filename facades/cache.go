package facades

import (
	"github.com/mewway/go-laravel/contracts/cache"
)

func Cache() cache.Cache {
	return App().MakeCache()
}
