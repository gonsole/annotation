package annotation

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
