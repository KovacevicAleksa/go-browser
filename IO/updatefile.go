package IO

import (
	"fmt"
	"os"
	"path/filepath"
)

func UpdateFile(name string, text string) {
	// Define the folder name
	folderName := "user_files"

	// Build the full file path
	filePath := filepath.Join(folderName, name)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", filePath)
		return
	}

	// Open the file in write-only mode, truncating it if it exists
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file %s: %v\n", filePath, cerr)
		}
	}()

	// Write the provided text to the file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("File updated successfully: %s\n", file.Name())
}
