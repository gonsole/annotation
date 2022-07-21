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
