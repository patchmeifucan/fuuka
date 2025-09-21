/*
walker.go - Handles the filewalking capabilities of Fuuka
Author: patchmeifucan
*/

package walker

import (
	"os"
	"path/filepath"
)

// Public

func Start() {
	var (
		root string // figure out how we wanna handle getting the path from user
	)

	err := filepath.WalkDir(root, path_handler)
	if err != nil {
		// some error logic, we'll figure this out eventually
	}
}

// Private

func path_handler(path string, entry os.DirEntry, err error) error {
	// write some logic and handle YARA with an external package

	return nil
}
