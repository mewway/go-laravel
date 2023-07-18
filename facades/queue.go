package facades

import (
	"github.com/mewway/go-laravel/contracts/queue"
)

func Queue() queue.Queue {
	return App().MakeQueue()
}
