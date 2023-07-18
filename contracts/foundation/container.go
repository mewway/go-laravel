package foundation

import (
	"github.com/mewway/go-laravel/contracts/auth"
	"github.com/mewway/go-laravel/contracts/auth/access"
	"github.com/mewway/go-laravel/contracts/cache"
	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/crypt"
	"github.com/mewway/go-laravel/contracts/database/orm"
	"github.com/mewway/go-laravel/contracts/database/seeder"
	"github.com/mewway/go-laravel/contracts/event"
	"github.com/mewway/go-laravel/contracts/filesystem"
	"github.com/mewway/go-laravel/contracts/grpc"
	"github.com/mewway/go-laravel/contracts/hash"
	"github.com/mewway/go-laravel/contracts/http"
	"github.com/mewway/go-laravel/contracts/log"
	"github.com/mewway/go-laravel/contracts/mail"
	"github.com/mewway/go-laravel/contracts/queue"
	"github.com/mewway/go-laravel/contracts/route"
	"github.com/mewway/go-laravel/contracts/schedule"
	"github.com/mewway/go-laravel/contracts/validation"
)

type Container interface {
	Bind(key any, callback func(app Application) (any, error))
	BindWith(key any, callback func(app Application, parameters map[string]any) (any, error))
	Instance(key, instance any)
	Make(key any) (any, error)
	MakeArtisan() console.Artisan
	MakeAuth() auth.Auth
	MakeCache() cache.Cache
	MakeConfig() config.Config
	MakeCrypt() crypt.Crypt
	MakeEvent() event.Instance
	MakeGate() access.Gate
	MakeGrpc() grpc.Grpc
	MakeHash() hash.Hash
	MakeLog() log.Log
	MakeMail() mail.Mail
	MakeOrm() orm.Orm
	MakeQueue() queue.Queue
	MakeRateLimiter() http.RateLimiter
	MakeRoute() route.Engine
	MakeSchedule() schedule.Schedule
	MakeStorage() filesystem.Storage
	MakeValidation() validation.Validation
	MakeSeeder() seeder.Facade
	MakeWith(key any, parameters map[string]any) (any, error)
	Singleton(key any, callback func(app Application) (any, error))
}
