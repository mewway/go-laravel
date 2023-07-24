package support

const Version string = "v1.0.0"

const (
	EnvRuntime = "runtime"
	EnvArtisan = "july"
	CliAppName = "Cicada July Skeleton"
	EnvTest    = "test"
)

const (
	DirApp        = "app"
	DirPolicy     = "policy"
	DirModel      = "model"
	DirDatabase   = "database"
	DirMigration  = "migration"
	DirObserver   = "observer"
	DirController = "controller"
	DirCommand    = "console"
	DirListener   = "listener"
	DirEvent      = "event"
	DirJob        = "job"
	DirRule       = "rule"
	DirSeeder     = "seeder"
	DirGrpc       = "grpc"
	DirConfig     = "config"
	DirMiddleware = "middleware"
	DirRequest    = "request"
)

const TabSpace = "    "

var (
	Env      = EnvRuntime
	EnvPath  = ".env"
	RootPath string
)
