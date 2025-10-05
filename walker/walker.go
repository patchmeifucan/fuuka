/*
 * walker.go - Handles the filewalking capabilities of Fuuka
 * Author: patchmeifucan
 */

package walker

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/hillu/go-yara/v4"
)

/* Public */

func Walker_Start(dir_path string, max_jobs int, yara_path string) {
	const matched_size int = 100

	var (
		err        error
		yara_rules *yara.Rules

		matched    chan string     = make(chan string, matched_size)
		semaphore  chan struct{}   = make(chan struct{}, max_jobs)
		wait_group *sync.WaitGroup = &sync.WaitGroup{}
	)

	yara_rules, err = yara.LoadRules(yara_path)
	error_handler(err)

	err = filepath.WalkDir(dir_path,
	func(path string, entry os.DirEntry, path_err error) error {
		semaphore <- struct{}{}
		wait_group.Add(1)
		go yara_match_handler(path, entry, path_err, semaphore, 
				wait_group, matched, yara_rules)
	
		return nil
	})

	wait_group.Wait()
	close(semaphore)
	close(matched)
	
	error_handler(err)

	yara_print_matches(dir_path, matched)
}

/* Private */

/*
 * Logic can become more complex if necessary
 * If such is the case, should be a separate package
 */
func error_handler(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func yara_match_handler(path string, entry os.DirEntry, path_err error,
			semaphore chan struct{}, wait_group *sync.WaitGroup,
			matched chan string, yara_rules *yara.Rules) {
	var (
		err          error
		match_entry  string
		yara_scanner *yara.Scanner
		yara_matches yara.MatchRules
	)

	/* Deferring emptying our semaphore and the waitgroup at the same time */
	defer func() {
		<-semaphore
		wait_group.Done()
	}()

	if path_err != nil {
		error_handler(path_err)
	} 
	
	yara_scanner, err = yara.NewScanner(yara_rules)
	error_handler(err)
  
	if !entry.IsDir() && entry.Type().IsRegular() {
		fmt.Printf("Scanning: %s\n", path)

		err = yara_scanner.SetCallback(&yara_matches).ScanFile(path)
		error_handler(err)
	
		/* 0 means we did not find a match on the file */
		if len(yara_matches) != 0 {
			match_entry = fmt.Sprintf("%s - %v", path, yara_matches)
			matched <- match_entry
		}
	}
}

func yara_print_matches(path string, matched chan string) {
	if len(matched) == 0 {
		fmt.Printf("No matches found in %s.\n", path)
		return
	}

	fmt.Println("Matches were found.")
	for range len(matched) {
		entry := <- matched
		fmt.Printf("%s\n", entry)
	}
}
