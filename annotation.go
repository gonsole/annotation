package annotation

import (
	_ "embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"text/template"
)

type Config struct {
	Directory string // Handler or Controller directory path.
	Output    string // Output file path; routes.go
}
type Route struct {
	Method  string // HTTP method
	Path    string // Path
	Handler string // Handler name
	Name    string // Route name
}
type App struct {
	Config
	Routes []Route
}

// Sets the configuration.
func (app *App) SetConfig(config Config) {
	app.Config = config
}

// Parses the handler directory.
func (app *App) ParseDirectory() {
	fileSet := token.NewFileSet()
	node, err := parser.ParseDir(fileSet, app.Directory, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, v := range node {
		for _, f := range v.Files {
			for _, d := range f.Decls {
				switch d.(type) {
				case *ast.FuncDecl:
					route, err := app.ParseAnnotationRoute(d.(*ast.FuncDecl).Doc.Text())
					if err != nil {
						log.Fatal(err)
					}
					route.Handler = d.(*ast.FuncDecl).Name.Name
					app.Routes = append(app.Routes, route)
				}
			}
		}
	}
}
func (app *App) ParseAnnotationRoute(annotation string) (Route, error) {
	regex := regexp.MustCompile(`@Route\(\"(.*?)\", name=\"(.*?)\", methods=\{\"(.*?)\"\}\)`)
	match := regex.FindStringSubmatch(annotation)
	if len(match) == 0 {
		err := errors.New("Invalid annotation")
		return Route{}, err
	}
	return Route{
		Method: match[3],
		Path:   match[1],
		Name:   match[2],
	}, nil
}

//go:embed template/fiber.tmpl
var fiberTemplate string

// Generates the routes.go file.
func (app *App) Generate() {
	tmpl, err := template.New("fiber").Parse(fiberTemplate)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(app.Output)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tmpl.Execute(file, map[string]interface{}{
		"Routes": app.Routes,
	})

}
