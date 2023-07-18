package foundation

import (
	"fmt"
	"sync"

	"github.com/gookit/color"

	"github.com/mewway/go-laravel/auth"
	"github.com/mewway/go-laravel/cache"
	"github.com/mewway/go-laravel/config"
	"github.com/mewway/go-laravel/console"
	authcontract "github.com/mewway/go-laravel/contracts/auth"
	accesscontract "github.com/mewway/go-laravel/contracts/auth/access"
	cachecontract "github.com/mewway/go-laravel/contracts/cache"
	configcontract "github.com/mewway/go-laravel/contracts/config"
	consolecontract "github.com/mewway/go-laravel/contracts/console"
	cryptcontract "github.com/mewway/go-laravel/contracts/crypt"
	ormcontract "github.com/mewway/go-laravel/contracts/database/orm"
	seerdercontract "github.com/mewway/go-laravel/contracts/database/seeder"
	eventcontract "github.com/mewway/go-laravel/contracts/event"
	filesystemcontract "github.com/mewway/go-laravel/contracts/filesystem"
	foundationcontract "github.com/mewway/go-laravel/contracts/foundation"
	grpccontract "github.com/mewway/go-laravel/contracts/grpc"
	hashcontract "github.com/mewway/go-laravel/contracts/hash"
	httpcontract "github.com/mewway/go-laravel/contracts/http"
	logcontract "github.com/mewway/go-laravel/contracts/log"
	mailcontract "github.com/mewway/go-laravel/contracts/mail"
	queuecontract "github.com/mewway/go-laravel/contracts/queue"
	routecontract "github.com/mewway/go-laravel/contracts/route"
	schedulecontract "github.com/mewway/go-laravel/contracts/schedule"
	validationcontract "github.com/mewway/go-laravel/contracts/validation"
	"github.com/mewway/go-laravel/crypt"
	"github.com/mewway/go-laravel/database"
	"github.com/mewway/go-laravel/event"
	"github.com/mewway/go-laravel/filesystem"
	"github.com/mewway/go-laravel/grpc"
	"github.com/mewway/go-laravel/hash"
	"github.com/mewway/go-laravel/http"
	goravellog "github.com/mewway/go-laravel/log"
	"github.com/mewway/go-laravel/mail"
	"github.com/mewway/go-laravel/queue"
	"github.com/mewway/go-laravel/route"
	"github.com/mewway/go-laravel/schedule"
	"github.com/mewway/go-laravel/validation"
)

type instance struct {
	concrete any
	shared   bool
}

type Container struct {
	bindings  sync.Map
	instances sync.Map
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Bind(key any, callback func(app foundationcontract.Application) (any, error)) {
	c.bindings.Store(key, instance{concrete: callback, shared: false})
}

func (c *Container) BindWith(key any, callback func(app foundationcontract.Application, parameters map[string]any) (any, error)) {
	c.bindings.Store(key, instance{concrete: callback, shared: false})
}

func (c *Container) Instance(key any, ins any) {
	c.bindings.Store(key, instance{concrete: ins, shared: true})
}

func (c *Container) Make(key any) (any, error) {
	return c.make(key, nil)
}

func (c *Container) MakeArtisan() consolecontract.Artisan {
	instance, err := c.Make(console.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(consolecontract.Artisan)
}

func (c *Container) MakeAuth() authcontract.Auth {
	instance, err := c.Make(auth.BindingAuth)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(authcontract.Auth)
}

func (c *Container) MakeCache() cachecontract.Cache {
	instance, err := c.Make(cache.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(cachecontract.Cache)
}

func (c *Container) MakeConfig() configcontract.Config {
	instance, err := c.Make(config.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(configcontract.Config)
}

func (c *Container) MakeCrypt() cryptcontract.Crypt {
	instance, err := c.Make(crypt.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(cryptcontract.Crypt)
}

func (c *Container) MakeEvent() eventcontract.Instance {
	instance, err := c.Make(event.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(eventcontract.Instance)
}

func (c *Container) MakeGate() accesscontract.Gate {
	instance, err := c.Make(auth.BindingGate)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(accesscontract.Gate)
}

func (c *Container) MakeGrpc() grpccontract.Grpc {
	instance, err := c.Make(grpc.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(grpccontract.Grpc)
}

func (c *Container) MakeHash() hashcontract.Hash {
	instance, err := c.Make(hash.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(hashcontract.Hash)
}

func (c *Container) MakeLog() logcontract.Log {
	instance, err := c.Make(goravellog.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(logcontract.Log)
}

func (c *Container) MakeMail() mailcontract.Mail {
	instance, err := c.Make(mail.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(mailcontract.Mail)
}

func (c *Container) MakeOrm() ormcontract.Orm {
	instance, err := c.Make(database.BindingOrm)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(ormcontract.Orm)
}

func (c *Container) MakeQueue() queuecontract.Queue {
	instance, err := c.Make(queue.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(queuecontract.Queue)
}

func (c *Container) MakeRateLimiter() httpcontract.RateLimiter {
	instance, err := c.Make(http.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(httpcontract.RateLimiter)
}

func (c *Container) MakeRoute() routecontract.Engine {
	instance, err := c.Make(route.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(routecontract.Engine)
}

func (c *Container) MakeSchedule() schedulecontract.Schedule {
	instance, err := c.Make(schedule.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(schedulecontract.Schedule)
}

func (c *Container) MakeStorage() filesystemcontract.Storage {
	instance, err := c.Make(filesystem.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(filesystemcontract.Storage)
}

func (c *Container) MakeValidation() validationcontract.Validation {
	instance, err := c.Make(validation.Binding)
	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(validationcontract.Validation)
}
func (c *Container) MakeSeeder() seerdercontract.Facade {
	instance, err := c.Make(database.BindingSeeder)

	if err != nil {
		color.Redln(err)
		return nil
	}

	return instance.(seerdercontract.Facade)
}

func (c *Container) MakeWith(key any, parameters map[string]any) (any, error) {
	return c.make(key, parameters)
}

func (c *Container) Singleton(key any, callback func(app foundationcontract.Application) (any, error)) {
	c.bindings.Store(key, instance{concrete: callback, shared: true})
}

func (c *Container) make(key any, parameters map[string]any) (any, error) {
	binding, ok := c.bindings.Load(key)
	if !ok {
		return nil, fmt.Errorf("binding not found: %+v", key)
	}

	if parameters == nil {
		instance, ok := c.instances.Load(key)
		if ok {
			return instance, nil
		}
	}

	bindingImpl := binding.(instance)
	switch concrete := bindingImpl.concrete.(type) {
	case func(app foundationcontract.Application) (any, error):
		concreteImpl, err := concrete(App)
		if err != nil {
			return nil, err
		}
		if bindingImpl.shared {
			c.instances.Store(key, concreteImpl)
		}

		return concreteImpl, nil
	case func(app foundationcontract.Application, parameters map[string]any) (any, error):
		concreteImpl, err := concrete(App, parameters)
		if err != nil {
			return nil, err
		}

		return concreteImpl, nil
	default:
		c.instances.Store(key, concrete)

		return concrete, nil
	}
}
