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

	err := filepath.WalkDir(root, pathHandler)
	if err != nil {
		// some error logic, we'll figure this out eventually
	}
}

// Private

func pathHandler(path string, entry os.DirEntry, err error) error {
	// write some logic and handle YARA with an external package

	return nil
}
