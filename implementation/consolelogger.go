package implementation

import "fmt"

// ConsoleLogger is an implementation of the Logger interface that logs messages to the console.
type ConsoleLogger struct {
	// Additional fields and configurations for the console logger can be added here.
}

// Log logs the given message to the console.
func (cl *ConsoleLogger) Log(message string) {
	// Code to print the message to the console.
	fmt.Println("[Console Logger] " + message)
}
