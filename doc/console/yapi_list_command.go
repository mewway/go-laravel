package console

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
	"github.com/mewway/go-laravel/doc/yapi"
)

type YapiListCommand struct {
	artisan console.Artisan
}

func NewYapiListCommand(atrisan console.Artisan) *YapiListCommand {
	return &YapiListCommand{
		artisan: atrisan,
	}
}

func (y YapiListCommand) Signature() string {
	return "doc:yapi-list"
}

func (y YapiListCommand) Description() string {
	return "List all yapi api list"
}

func (y YapiListCommand) Extend() command.Extend {
	return command.Extend{
		Category: "doc",
	}
}

func (y YapiListCommand) Handle(ctx console.Context) error {
	c := yapi.NewClient()
	if c == nil {
		return errors.New("please complete the doc configuration first")
	}
	l := c.QueryApiList()
	color.Cyanln("Yapi api list:")
	for i, v := range l.List {
		color.Grayln(fmt.Sprintf("%d.【%s】%s: %s", i+1, strings.ToUpper(v.Method), v.Path, v.Title))
	}
	return nil
}
