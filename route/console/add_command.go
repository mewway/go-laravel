package console

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/console/command"
	"github.com/mewway/go-laravel/support"
	"github.com/mewway/go-laravel/support/str"
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
	method := strings.ToUpper(ctx.Argument(0))
	api := ctx.Argument(1)
	desc := ctx.Argument(2)

	if !strings.HasPrefix(api, "/") {
		api = "/" + api
	}
	if s := getIllegalChars(api); len(s) > 0 {
		return errors.New(fmt.Sprintf("Route【%s】contains illegal chars: %s", api, strings.Join(s, "、")))
	}
	routes := receiver.callback()
	for _, r := range routes {
		n := fmt.Sprintf("【%s】%s", method, api)
		if r == n {
			return errors.New(fmt.Sprintf("%s is already exists!", n))
		}
	}
	fmt.Println(method, api, desc)
	version, file, funcName := parseRoute(api)
	fmt.Println(version)
	pwd, _ := os.Getwd()
	filePath := filepath.Join(pwd, support.DirRoute, file+".go")
	_, err := os.Stat(filePath)

	// ast解析处理
	var f *ast.File
	t := token.NewFileSet()

	// 文件不存在的时候
	if err != nil {
		f = &ast.File{
			Name:    &ast.Ident{Name: "route"},
			Imports: []*ast.ImportSpec{},
			Decls:   []ast.Decl{},
		}
		f.Imports = append(f.Imports, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: `"github.com/mewway/go-laravel/contracts/http"`,
			},
		}, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: `"fmt"`,
			},
		})
	} else {
		f, err = parser.ParseFile(t, filePath, nil, parser.ParseComments)
		if err != nil {
			return errors.New(fmt.Sprintf("File[%s] cant't be parsed by ast", filePath))
		}
	}

	var targetFunc *ast.FuncDecl
	for _, decl := range f.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok && fd.Name != nil && fd.Name.Name == funcName {
			targetFunc = fd
			break
		}
	}
	// 函数已经存在目标文件中 追加路由
	if targetFunc != nil {

	} else {
		targetFunc = &ast.FuncDecl{
			Doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: fmt.Sprintf("// %s", funcName),
					},
				},
			},
			Name: ast.NewIdent(funcName),
			Type: &ast.FuncType{
				Func:       0,
				TypeParams: nil,
				Params: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{
								ast.NewIdent("route"),
							},
							Type: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "route"},
								Sel: &ast.Ident{Name: "Route"},
							},
						},
					},
				},
				Results: nil,
			},
			Body: &ast.BlockStmt{},
		}
		f.Decls = append(f.Decls, targetFunc)
	}

	output, err := os.Create(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("File create or update fail: %s", err))
	}

	defer output.Close()

	err = printer.Fprint(output, t, f)
	if err != nil {
		return errors.New(fmt.Sprintf("File ouput error occured: %s", err))
	}
	// 新建函数并创建路由
	return nil
}

func parseRoute(route string) (version string, file string, funcName string) {
	if route == "/" || strings.HasPrefix(version, "/:") {
		version = "default"
		file = support.FileRouteDefault
		funcName = file + "Init"
	} else if m, _ := regexp.MatchString(`/(v\d+)(/([a-zA-Z0-9\:_\-]+))+`, route); m == true {
		groups := strings.Split(route, "/")
		version = groups[1]
		file = str.Camel2Case(groups[2])
		funcName = str.Case2Camel(file) + "Init"
		funcName = strings.ToLower(funcName[:1]) + funcName[1:]
	} else {
		groups := strings.Split(route, "/")
		version = groups[1]
		file = str.Camel2Case(version)
		funcName = str.Case2Camel(file) + "Init"
		funcName = strings.ToLower(funcName[:1]) + funcName[1:]
	}
	fmt.Println(version, file, funcName)
	return
}

func getIllegalChars(route string) []string {
	regex := regexp.MustCompile(`[^\/a-zA-Z0-9\:_\-]`)
	ill := regex.FindAllString(route, -1)
	fmt.Println(ill, "holy")
	return ill
}
