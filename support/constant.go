package support

const Version string = "v1.0.0"

const (
	EnvRuntime = "runtime"
	EnvArtisan = "july"
	CliAppName = "Cicada July Skeleton"
	EnvTest    = "test"
)

var (
	Env      = EnvRuntime
	EnvPath  = ".env"
	RootPath string
)
