package facades

import (
	"github.com/mewway/go-laravel/contracts/log"
)

func Log() log.Log {
	return App().MakeLog()
}
