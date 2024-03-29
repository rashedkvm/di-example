package nestedfolders

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateNestedDirectoriesAndFile(path string) error {
	// Split the path into directory path and file name
	lastSlashIndex := strings.LastIndex(path, "/")
	dirPath := path[:lastSlashIndex]
	fileName := path[lastSlashIndex+1:]

	// Create nested directories
	if dirPath != "" {
		dirs := strings.Split(dirPath, "/")
		currentPath := "."
		for _, dir := range dirs {
			currentPath = filepath.Join(currentPath, dir)
			if err := os.Mkdir(currentPath, 0755); err != nil {
				if !os.IsExist(err) {
					return err
				}
			}
		}
	}

	// Create a file inside the deepest directory
	filePath := filepath.Join(dirPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write some content to the file
	_, err = file.WriteString("This is a test file.")
	if err != nil {
		return err
	}

	fmt.Printf("File created at %s\n", filePath)
	return nil
}

func CreateNestedDirectoriesAndFileOld(path string) error {
	// Split the path into individual directory names
	dirs := strings.Split(path, "/")

	// Create nested directories
	currentPath := "."
	for _, dir := range dirs {
		currentPath = filepath.Join(currentPath, dir)
		if err := os.Mkdir(currentPath, 0755); err != nil {
			if !os.IsExist(err) {
				return err
			}
		}
	}

	// Create a file inside the deepest directory
	filePath := filepath.Join(currentPath, "example.txt")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write some content to the file
	_, err = file.WriteString("This is a test file.")
	if err != nil {
		return err
	}

	fmt.Printf("File created at %s\n", filePath)
	return nil
}
