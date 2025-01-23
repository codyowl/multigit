package utilities

import (
	"fmt"
	"os"
	"path/filepath"

)

// helper function to read yaml file
func readCommandsyml(){

}

// helper function to check if the cli has been triggered from a git repository 
func IsGitRepository() bool {
	
	currentWorkingDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while trying to get current working directory:", err)
		return false
	}

	// checking for .git file
	for {
		if _, err := os.Stat(filepath.Join(currentWorkingDir, ".git")); err == nil {
			return true
		}

		// stoping at parent directory
		parentDir := filepath.Dir(currentWorkingDir)
		if parentDir == currentWorkingDir{
			break
		}
	}
	return false

}