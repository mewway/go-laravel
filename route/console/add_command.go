package console

import (
	"github.com/gookit/color"
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
)

type AddCommand struct {
	artisan  console.Artisan
	callback func() []string
}

func NewAddCommand(artisan console.Artisan, callback func() []string) *AddCommand {
	return &AddCommand{
		artisan:  artisan,
		callback: callback,
	}
}

// Signature The name and signature of the console command.
func (receiver *AddCommand) Signature() string {
	return "route:add"
}

// Description The console command description.
func (receiver *AddCommand) Description() string {
	return "Add a new route to system"
}

// Extend The console command extend.
func (receiver *AddCommand) Extend() command.Extend {
	return command.Extend{
		Category: "route",
		Args: []command.Arg{
			&command.StringArg{
				Name:     "method",
				Usage:    "request method",
				Required: true,
			},
			&command.StringArg{
				Name:     "api",
				Usage:    "Request specific uri",
				Required: true,
			},
			&command.StringArg{
				Name:     "desc",
				Usage:    "Api specific description or usage",
				Required: false,
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *AddCommand) Handle(ctx console.Context) error {
	routes := receiver.callback()
	color.Greenln("Routes List:")
	for i, v := range routes {
		if i%2 == 0 {
			color.Cyanln(v)
		} else {
			color.Grayln(v)
		}

	}
	return nil
}
