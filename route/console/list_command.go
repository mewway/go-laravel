package console

import (
	"github.com/gookit/color"
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
)

type ListCommand struct {
	artisan  console.Artisan
	callback func() []string
}

func NewListCommand(artisan console.Artisan, callback func() []string) *ListCommand {
	return &ListCommand{
		artisan:  artisan,
		callback: callback,
	}
}

// Signature The name and signature of the console command.
func (receiver *ListCommand) Signature() string {
	return "route:list"
}

// Description The console command description.
func (receiver *ListCommand) Description() string {
	return "List all routes"
}

// Extend The console command extend.
func (receiver *ListCommand) Extend() command.Extend {
	return command.Extend{
		Category: "route",
	}
}

// Handle Execute the console command.
func (receiver *ListCommand) Handle(ctx console.Context) error {
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
