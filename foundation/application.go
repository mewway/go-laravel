package foundation

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/mewway/go-laravel/config"
	consolecontract "github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/foundation/console"
	"github.com/mewway/go-laravel/support"
	"github.com/mewway/go-laravel/support/carbon"
)

var (
	App foundation.Application
)

func init() {
	setEnv()

	app := &Application{
		Container:     NewContainer(),
		publishes:     make(map[string]map[string]string),
		publishGroups: make(map[string]map[string]string),
	}
	app.registerBaseServiceProviders()
	app.bootBaseServiceProviders()
	App = app
}

type Application struct {
	foundation.Container
	publishes     map[string]map[string]string
	publishGroups map[string]map[string]string
}

func NewApplication() foundation.Application {
	return App
}

// Boot Register and bootstrap configured service providers.
func (app *Application) Boot() {
	app.registerConfiguredServiceProviders()
	app.bootConfiguredServiceProviders()
	app.registerCommands([]consolecontract.Command{
		console.NewPackageMakeCommand(),
		console.NewVendorPublishCommand(app.publishes, app.publishGroups),
	})
	app.bootArtisan()
	app.setTimezone()
	setRootPath()
}

func (app *Application) Commands(commands []consolecontract.Command) {
	app.registerCommands(commands)
}

func (app *Application) Path(path string) string {
	return filepath.Join("app", path)
}

func (app *Application) BasePath(path string) string {
	return filepath.Join("", path)
}

func (app *Application) ConfigPath(path string) string {
	return filepath.Join("config", path)
}

func (app *Application) DatabasePath(path string) string {
	return filepath.Join("database", path)
}

func (app *Application) StoragePath(path string) string {
	return filepath.Join("storage", path)
}

func (app *Application) PublicPath(path string) string {
	return filepath.Join("public", path)
}

func (app *Application) Publishes(packageName string, paths map[string]string, groups ...string) {
	app.ensurePublishArrayInitialized(packageName)

	for key, value := range paths {
		app.publishes[packageName][key] = value
	}

	for _, group := range groups {
		app.addPublishGroup(group, paths)
	}
}

func (app *Application) ensurePublishArrayInitialized(packageName string) {
	if _, exist := app.publishes[packageName]; !exist {
		app.publishes[packageName] = make(map[string]string)
	}
}

func (app *Application) addPublishGroup(group string, paths map[string]string) {
	if _, exist := app.publishGroups[group]; !exist {
		app.publishGroups[group] = make(map[string]string)
	}

	for key, value := range paths {
		app.publishGroups[group][key] = value
	}
}

// bootArtisan Boot artisan command.
func (app *Application) bootArtisan() {
	app.MakeArtisan().Run(os.Args, true)
}

// getBaseServiceProviders Get base service providers.
func (app *Application) getBaseServiceProviders() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&config.ServiceProvider{},
	}
}

// getConfiguredServiceProviders Get configured service providers.
func (app *Application) getConfiguredServiceProviders() []foundation.ServiceProvider {
	return app.MakeConfig().Get("app.providers").([]foundation.ServiceProvider)
}

// registerBaseServiceProviders Register base service providers.
func (app *Application) registerBaseServiceProviders() {
	app.registerServiceProviders(app.getBaseServiceProviders())
}

// bootBaseServiceProviders Bootstrap base service providers.
func (app *Application) bootBaseServiceProviders() {
	app.bootServiceProviders(app.getBaseServiceProviders())
}

// registerConfiguredServiceProviders Register configured service providers.
func (app *Application) registerConfiguredServiceProviders() {
	app.registerServiceProviders(app.getConfiguredServiceProviders())
}

// bootConfiguredServiceProviders Bootstrap configured service providers.
func (app *Application) bootConfiguredServiceProviders() {
	app.bootServiceProviders(app.getConfiguredServiceProviders())
}

// registerServiceProviders Register service providers.
func (app *Application) registerServiceProviders(serviceProviders []foundation.ServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Register(app)
	}
}

// bootServiceProviders Bootstrap service providers.
func (app *Application) bootServiceProviders(serviceProviders []foundation.ServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot(app)
	}
}

func (app *Application) registerCommands(commands []consolecontract.Command) {
	app.MakeArtisan().Register(commands)
}

func (app *Application) setTimezone() {
	carbon.SetTimezone(app.MakeConfig().GetString("app.timezone", carbon.UTC))
}

func setEnv() {
	args := os.Args
	if strings.HasSuffix(os.Args[0], ".test") || strings.HasSuffix(os.Args[0], ".test.exe") {
		support.Env = support.EnvTest
	}
	if len(args) >= 2 {
		for _, arg := range args[1:] {
			if arg == "artisan" {
				support.Env = support.EnvArtisan
				break
			}
		}
	}

	var env *string
	if !flag.Parsed() && support.Env != support.EnvTest {
		env = flag.String("env", ".env", "custom .env path")
		flag.Parse()
	} else {
		testEnv := ".env"
		env = &testEnv
	}

	support.EnvPath = *env
}

func setRootPath() {
	rootPath := getCurrentAbsolutePath()

	// Hack air path
	airPath := "/storage/temp"
	if strings.HasSuffix(rootPath, airPath) {
		rootPath = strings.ReplaceAll(rootPath, airPath, "")
	}

	support.RootPath = rootPath
}
