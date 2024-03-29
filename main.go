package main

import (
	"fmt"

	"github.com/rashedkvm/di-example/app"
	"github.com/rashedkvm/di-example/implementation"
	"github.com/rashedkvm/di-example/nestedfolders"
)

const filepath = `nf/test/cnb/spdx.json`

func main() {
	println("hello")
	fmt.Printf("hello %q \n", "test")

	// Create a FileLogger instance for dependency injection.
	fileLogger := &implementation.FileLogger{}
	// Create a ConsoleLogger instance for dependency injection.
	consoleLogger := &implementation.ConsoleLogger{}
	// Create a new Application instance with FileLogger dependency.
	appWithFileLogger := app.NewApplication(fileLogger)
	appWithFileLogger.Run()
	// Create a new Application instance with ConsoleLogger dependency.
	appWithConsoleLogger := app.NewApplication(consoleLogger)
	appWithConsoleLogger.Run()

	if err := nestedfolders.CreateNestedDirectoriesAndFile(filepath); err != nil {
		fmt.Println(err)
	}
}
