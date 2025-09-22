/*
walker.go - Handles the filewalking capabilities of Fuuka
Author: patchmeifucan
*/

package walker

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hillu/go-yara/v4"
)

/* Public */

func Walker_Start(dir_path string, yara_path string) {
	var (
		err 				error
		yara_rules 	*yara.Rules
	)

	yara_rules, err = yara.LoadRules(yara_path)
	error_handler(err)

	yara_scanner, err = yara.NewScanner(yara_rules)
	error_handler(err)

	err = filepath.WalkDir(dir_path, path_handler)
	error_handler(err)

	yara_print_matches(dir_path)
}

/* Private */

/*
Private scope variables for YARA log to limit use to walker.go
It's cleaner to read than wrapping path_handler in another func
Did not see need to call yara_scanner with ptr_ prefix, no need in my opinion
*/
var (
	/*
	The size variable really doesn't need to be 500 I don't think
	But these slices absolutely need a big size
	We don't know how many files may get matched here
	*/
	matched_size 	int = 100
	matched_files	[]string = make([]string, 0, matched_size)
	matched_rules	[]yara.MatchRule = make([]yara.MatchRule, 0, matched_size)

	yara_scanner 	*yara.Scanner
)

/*
Logic can become more complex if necessary
If such is the case, should be a separate package
*/
func error_handler(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func path_handler(path string, entry os.DirEntry, path_err error) error {
	var (
		err_yara 			error
		yara_matches 	yara.MatchRules
	)

	if path_err != nil {
		return path_err
	}

	if !entry.IsDir() {
		fmt.Printf("Scanning: %s\n", path)
		err_yara = yara_scanner.SetCallback(&yara_matches).ScanFile(path)
		yara_match_handler(path, yara_matches, err_yara)
	}

	return nil
}

func yara_match_handler(file string, matches []yara.MatchRule, err error) {
	error_handler(err)

	/* 0 means that we did not find a match on the file */
	if len(matches) == 0 {
		return
	}

	/* Append to slices for printing out when walk is finished to show matches */
	matched_files = append(matched_files, file)
	matched_rules = append(matched_rules, matches...)
}

func yara_print_matches(path string) {
	if len(matched_rules) == 0 {
		fmt.Printf("No matches found in %s.\n", path)
		return
	}

	fmt.Println("Matches were found.")
	for index, entry := range matched_files {
		fmt.Printf("File: %s\nRule Match: %v\n", entry, matched_rules[index])
	}
}
