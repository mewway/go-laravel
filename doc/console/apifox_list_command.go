package console

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
	"github.com/mewway/go-laravel/doc/apifox"
)

type ApifoxListCommand struct {
	artisan console.Artisan
}

func NewApifoxListCommand(atrisan console.Artisan) *ApifoxListCommand {
	return &ApifoxListCommand{
		artisan: atrisan,
	}
}

func (y ApifoxListCommand) Signature() string {
	return "doc:apifox-list"
}

func (y ApifoxListCommand) Description() string {
	return "List all apifox api list"
}

func (y ApifoxListCommand) Extend() command.Extend {
	return command.Extend{
		Category: "doc",
	}
}

func (y ApifoxListCommand) Handle(ctx console.Context) error {
	c := apifox.NewClient()
	if c == nil {
		return errors.New("please complete the doc configuration first")
	}
	l := c.QueryApiList()
	color.Cyanln("Apifox api list:")
	for i, v := range l.Data {
		color.Grayln(fmt.Sprintf("%d.【%s】%s: %s", i+1, strings.ToUpper(v.Method), v.Path, v.Name))
	}
	return nil
}
