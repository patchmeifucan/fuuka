/*
walker.go - Handles the filewalking capabilities of Fuuka
Author: patchmeifucan
*/

package walker

import (
	"os"
	"path/filepath"
)

/* Public */

func Walker_Start() {
	var (
		root string /* figure out how we wanna handle getting the path from user */
	)

	err := filepath.WalkDir(root, path_handler)
	if err != nil {
		return
	}
}

/* Private */

func path_handler(path string, entry os.DirEntry, err error) error {
	if err != nil {
		return err
	}

	/* write YARA logic */

	return nil
}
