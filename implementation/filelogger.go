package implementation

import "fmt"

// FileLogger is an implementation of the Logger interface that logs messages to a file.
type FileLogger struct {
	// Additional fields and configurations for the file logger can be added here.
}

// Log logs the given message to a file.
func (fl *FileLogger) Log(message string) {
	// Code to write the message to a file.
	fmt.Println("[File Logger] " + message)
}
