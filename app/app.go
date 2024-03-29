package app

import (
	"example.com/di-example/di"
)

// Application represents our main application that depends on a Logger.
type Application struct {
	logger di.Logger // Dependency injection of the logger.
}

// NewApplication creates a new instance of Application with the given logger dependency.
func NewApplication(logger di.Logger) *Application {
	return &Application{logger: logger}
}

// Run starts the application and logs a message.
func (app *Application) Run() {
	// Application logic goes here.
	app.logger.Log("Application is running.")
}
