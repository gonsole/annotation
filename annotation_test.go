package annotation_test

import (
	"testing"

	"github.com/gonsole/annotation"
)

func TestNew(t *testing.T) {
	app := new(annotation.App)
	if app == nil {
		t.Error("New() returned nil")
	}
}
func TestSetConfig(t *testing.T) {
	app := new(annotation.App)
	app.SetConfig(annotation.Config{
		Directory: ".",
		Output:    "routes.go",
	})
	if app.Config.Directory != "." {
		t.Error("SetConfig() failed")
	}
	if app.Config.Output != "routes.go" {
		t.Error("SetConfig() failed")
	}
}
func TestParseDirectory(t *testing.T) {
	app := new(annotation.App)
	app.SetConfig(annotation.Config{
		Directory: "./example",
		Output:    "routes.go",
	})
	app.ParseDirectory()
	if len(app.Routes) == 0 {
		t.Error("ParseDirectory() failed")
	}
}
func TestParseAnnotationRoute(t *testing.T) {
	app := new(annotation.App)
	app.SetConfig(annotation.Config{
		Directory: ".",
		Output:    "routes.go",
	})
	annotation := `@Route("/foo", name="foo", methods={"GET"})`
	route, err := app.ParseAnnotationRoute(annotation)
	if err != nil {
		t.Error("ParseAnnotationRoute() failed")
	}
	if route.Method != "GET" {
		t.Error("ParseAnnotationRoute() failed")
	}
	if route.Path != "/foo" {
		t.Error("ParseAnnotationRoute() failed")
	}
	if route.Name != "foo" {
		t.Error("ParseAnnotationRoute() failed")
	}
}
func TestGenerate(t *testing.T) {
	app := new(annotation.App)
	app.SetConfig(annotation.Config{
		Directory: "./handler",
		Output:    "routes.go",
	})
	app.ParseDirectory()
	app.Generate()

}
