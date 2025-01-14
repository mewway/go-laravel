package console

import (
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
)

type ListCommand struct {
	artisan console.Artisan
}

func NewListCommand(artisan console.Artisan) *ListCommand {
	return &ListCommand{
		artisan: artisan,
	}
}

// Signature The name and signature of the console command.
func (receiver *ListCommand) Signature() string {
	return "list"
}

// Description The console command description.
func (receiver *ListCommand) Description() string {
	return "List commands"
}

// Extend The console command extend.
func (receiver *ListCommand) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *ListCommand) Handle(ctx console.Context) error {
	receiver.artisan.Call("--help")

	return nil
}
