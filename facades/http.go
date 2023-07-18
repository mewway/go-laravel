package facades

import (
	"github.com/mewway/go-laravel/contracts/http"
)

func RateLimiter() http.RateLimiter {
	return App().MakeRateLimiter()
}
