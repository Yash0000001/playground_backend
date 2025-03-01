package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(fileName, code string) (string, error) {
	folderPath := "files"                           // Specify the target folder
	fullPath := filepath.Join(folderPath, fileName) // Create the full path

	// Ensure the directory exists
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return "", fmt.Errorf("failed to create folder")
	}
	err = os.WriteFile(fullPath, []byte(code), 0644)

	if err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}
	return fullPath, nil
}
