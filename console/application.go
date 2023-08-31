package console

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
	"github.com/mewway/go-laravel/support"
)

type Application struct {
	instance   *cli.App
	commandMap map[string]console.Command
}

func NewApplication() console.Artisan {
	instance := cli.NewApp()
	instance.Name = support.CliAppName
	instance.Usage = support.Version
	instance.UsageText = support.EnvArtisan + " [global options] command [options] [arguments...]"

	return &Application{instance, make(map[string]console.Command)}
}

func (c *Application) Register(commands []console.Command) {
	for _, item := range commands {
		item := item
		cliCommand := cli.Command{
			Name:      item.Signature(),
			ArgsUsage: item.Extend().ArgsUsage(),
			Usage:     item.Description(),
			Action: func(ctx *cli.Context) error {
				return item.Handle(&CliContext{ctx})
			},
		}

		cliCommand.Category = item.Extend().Category
		cliCommand.Flags = flagsToCliFlags(item.Extend().Flags)
		c.instance.Commands = append(c.instance.Commands, &cliCommand)
		c.commandMap[item.Signature()] = item
	}
}

// Call Run an Artisan console command by name.
func (c *Application) Call(command string) {
	c.Run(append([]string{os.Args[0], support.EnvArtisan}, strings.Split(command, " ")...), false)
}

// CallAndExit Run an Artisan console command by name and exit.
func (c *Application) CallAndExit(command string) {
	c.Run(append([]string{os.Args[0], support.EnvArtisan}, strings.Split(command, " ")...), true)
}

// Run a command. Args come from os.Args.
func (c *Application) Run(args []string, exitIfArtisan bool) {
	artisanIndex := -1
	for i, arg := range args {
		if arg == support.EnvArtisan {
			artisanIndex = i
			break
		}
	}

	if artisanIndex != -1 {
		// Add --help if no command argument is provided.
		if artisanIndex+1 == len(args) {
			args = append(args, "--help")
		}

		if args[artisanIndex+1] != "-V" && args[artisanIndex+1] != "--version" {
			fmt.Println(args)
			cliArgs := append([]string{args[0]}, args[artisanIndex+1:]...)
			cmd := c.instance.Command(cliArgs[1])
			if cmd != nil {
				signature := cliArgs[1]
				if cslCmd, ok := c.commandMap[signature]; ok == true {
					// TODO
					fmt.Println("holy shit:", signature, cslCmd.Description())
				}
			}
			if err := c.instance.Run(cliArgs); err != nil {
				color.Errorln(err.Error())
				if exitIfArtisan {
					os.Exit(2)
				}
			}
		}

		printResult(args[artisanIndex+1])

		if exitIfArtisan {
			os.Exit(0)
		}
	}
}

func flagsToCliFlags(flags []command.Flag) []cli.Flag {
	var cliFlags []cli.Flag
	for _, flag := range flags {
		switch flag.Type() {
		case command.ArgTypeBool:
			flag := flag.(*command.BoolFlag)
			cliFlags = append(cliFlags, &cli.BoolFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    flag.Value,
			})
		case command.ArgTypeFloat64:
			flag := flag.(*command.Float64Flag)
			cliFlags = append(cliFlags, &cli.Float64Flag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    flag.Value,
			})
		case command.ArgTypeFloat64Slice:
			flag := flag.(*command.Float64SliceFlag)
			cliFlags = append(cliFlags, &cli.Float64SliceFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    cli.NewFloat64Slice(flag.Value...),
			})
		case command.ArgTypeInt:
			flag := flag.(*command.IntFlag)
			cliFlags = append(cliFlags, &cli.IntFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    flag.Value,
			})
		case command.ArgTypeIntSlice:
			flag := flag.(*command.IntSliceFlag)
			cliFlags = append(cliFlags, &cli.IntSliceFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    cli.NewIntSlice(flag.Value...),
			})
		case command.ArgTypeInt64:
			flag := flag.(*command.Int64Flag)
			cliFlags = append(cliFlags, &cli.Int64Flag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    flag.Value,
			})
		case command.ArgTypeInt64Slice:
			flag := flag.(*command.Int64SliceFlag)
			cliFlags = append(cliFlags, &cli.Int64SliceFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    cli.NewInt64Slice(flag.Value...),
			})
		case command.ArgTypeString:
			flag := flag.(*command.StringFlag)
			cliFlags = append(cliFlags, &cli.StringFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    flag.Value,
			})
		case command.ArgTypeStringSlice:
			flag := flag.(*command.StringSliceFlag)
			cliFlags = append(cliFlags, &cli.StringSliceFlag{
				Name:     flag.Name,
				Aliases:  flag.Aliases,
				Usage:    flag.Usage,
				Required: flag.Required,
				Value:    cli.NewStringSlice(flag.Value...),
			})
		}
	}

	return cliFlags
}

func printResult(command string) {
	switch command {
	case "make:command":
		color.Greenln("Console command created successfully")
	case "-V", "--version":
		color.Greenln("Cicada Framework " + support.Version)
	}
}
