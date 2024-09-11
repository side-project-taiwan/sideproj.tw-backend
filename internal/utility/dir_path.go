package utility

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectRootDirAndEnvPath() (string, string, error) {
	dir, err := os.Getwd() // Get the current working directory
	if err != nil {
		return "", "", err
	}

	for {
		files, err := os.ReadDir(dir)
		if err != nil {
			return "", "", err
		}

		for _, file := range files {
			if file.Name() == "go.mod" {
				envPath := fmt.Sprintf("%s/.env", dir) // Combine the project root with ".env"
				return dir, envPath, nil               // Return the project root directory and the .env path
			}
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break // Reached the root directory of the filesystem, stop traversing
		}
		dir = parentDir
	}

	return "", "", fmt.Errorf("project root directory not found")
}
