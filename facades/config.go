package facades

import (
	"github.com/mewway/go-laravel/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
