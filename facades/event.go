package facades

import "github.com/mewway/go-laravel/contracts/event"

func Event() event.Instance {
	return App().MakeEvent()
}
