package http

import (
	"github.com/mewway/go-laravel/contracts/cache"
	"github.com/mewway/go-laravel/contracts/config"
	consolecontract "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/contracts/http"
	"github.com/mewway/go-laravel/contracts/log"
	"github.com/mewway/go-laravel/contracts/validation"
	"github.com/mewway/go-laravel/http/console"
)

const Binding = "cicada.http"

var (
	ConfigFacade      config.Config
	CacheFacade       cache.Cache
	LogFacade         log.Log
	RateLimiterFacade http.RateLimiter
	ValidationFacade  validation.Validation
)

type ServiceProvider struct {
}

func (http *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewRateLimiter(), nil
	})
}

func (http *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	CacheFacade = app.MakeCache()
	LogFacade = app.MakeLog()
	RateLimiterFacade = app.MakeRateLimiter()
	ValidationFacade = app.MakeValidation()

	http.registerCommands(app)
}

func (http *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
