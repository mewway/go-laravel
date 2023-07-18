package facades

import (
	"github.com/mewway/go-laravel/contracts/route"
)

func Route() route.Engine {
	return App().MakeRoute()
}
