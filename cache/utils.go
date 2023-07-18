package cache

import (
	"github.com/mewway/go-laravel/contracts/config"
)

func prefix(config config.Config) string {
	return config.GetString("cache.prefix") + ":"
}
