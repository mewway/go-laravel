package console

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/mewway/go-laravel/contracts/config"
	"github.com/mewway/go-laravel/llm"
	"github.com/mewway/go-laravel/llm/function"
	"github.com/mewway/go-laravel/llm/prompt"
	"github.com/mewway/go-laravel/support"

	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
	"github.com/mewway/go-laravel/support/file"
	"github.com/mewway/go-laravel/support/str"
)

type ModelMakeCommand struct {
	config config.Config
}

func NewModelMakeCommand(config config.Config) *ModelMakeCommand {
	return &ModelMakeCommand{
		config: config,
	}
}

// Signature The name and signature of the console command.
func (receiver *ModelMakeCommand) Signature() string {
	return "make:model"
}

// Description The console command description.
func (receiver *ModelMakeCommand) Description() string {
	return "Create a new model class"
}

// Extend The console command extend.
func (receiver *ModelMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
	}
}

// Handle Execute the console command.
func (receiver *ModelMakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	database := ctx.Option("database")
	if name == "" {
		color.Redln("Not enough arguments (missing: name)")

		return nil
	}

	if err := file.Create(receiver.getPath(name), receiver.populateStub(receiver.getStub(), name, database)); err != nil {
		return err
	}

	color.Greenln("Model created successfully")

	return nil
}

func (receiver *ModelMakeCommand) getStub() string {
	return Stubs{}.Model()
}

// populateStub Populate the place-holders in the command stub.
func (receiver *ModelMakeCommand) populateStub(stub, name, database string) string {
	modelName, packageName, _ := receiver.parseName(name)

	stub = strings.ReplaceAll(stub, "DummyModel", str.Case2Camel(modelName))
	stub = strings.ReplaceAll(stub, "DummyPackage", packageName)

	helper, err := NewStructHelper(receiver.config)
	if err != nil {
		return stub
	}
	db := helper.DefaultDb
	if database != "" {
		db = database
	}
	cols := helper.GetTableStruct(db, modelName)
	comments := ""
	columns := []string{}
	for _, col := range cols {
		comments += fmt.Sprintf("【%s】%s \n", col.ColumnName, col.ColumnComment)
		columns = append(columns, col.ColumnName)
	}
	// 查出表数据去重数量
	countMapString := function.NewDB(receiver.config).QueryTableDistinctCountMap(db, name, columns)
	// 查出建表语句
	ddl := function.NewDB(receiver.config).DDL(db, name)

	llm.NewGpt(receiver.config).Resolve(prompt.TableCheckResolver(ddl, countMapString)).ChatCompletion("检查表结构是否存在错误")
	// 检查并定义 const
	reply := llm.NewGpt(receiver.config).Resolve(prompt.ConstResolver(comments)).ChatCompletion("表结构体生成")
	dummyField := helper.StringColumns(cols)
	stub = strings.ReplaceAll(stub, "DummyConst", reply)
	stub = strings.ReplaceAll(stub, "DummyField", dummyField)
	source, err := format.Source([]byte(stub))
	if err != nil {
		return stub
	}
	return string(source)
}

// getPath Get the full path to the command.
func (receiver *ModelMakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	modelName, _, folderPath := receiver.parseName(name)

	return filepath.Join(pwd, support.DirApp, support.DirModel, folderPath, str.Camel2Case(modelName)+".go")
}

// parseName Parse the name to get the model name, package name and folder path.
func (receiver *ModelMakeCommand) parseName(name string) (string, string, string) {
	name = strings.TrimSuffix(name, ".go")

	segments := strings.Split(name, "/")

	modelName := segments[len(segments)-1]

	packageName := support.DirModel
	folderPath := ""

	if len(segments) > 1 {
		folderPath = filepath.Join(segments[:len(segments)-1]...)
		packageName = segments[len(segments)-2]
	}

	return modelName, packageName, folderPath
}
