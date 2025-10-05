/*
 * args.go - Entry point for handling command line arguments
 * Author: patchmeifucan
 */

package args

import (
	"flag"
	"fmt"
	"os"

	"patchmeifucan/fuuka/consts"
)

/* Public */

/*
 * Handling any user flags provided, separate from main for sake of sanity
 * Returns a string and an error should flag.NFlag() be > 1
 * There is no reason for the user to use more than 1 flag with Fuuka
 * file_path is stored within main.go as a string to pass to Walker_Start()
 * It's much easier than just having this logic in the same directory as walker.go
 */
func Arg_Handler() (string, int, string, error) {
	var (
		dir_path  string
		err_flag  error
		max_jobs  int
		yara_path string
	)

	/*
	 * Option of either -p or -path for setting the path
	 */
	flag.StringVar(&dir_path, "p", "", "Set path for Fuuka to scan")
	flag.StringVar(&dir_path, "path", "", "Set path for Fuuka to scan")
	flag.IntVar(&max_jobs, "j", consts.FUUKA_DEFAULT_JOBS, "Set max jobs")
	flag.IntVar(&max_jobs, "jobs", consts.FUUKA_DEFAULT_JOBS, "Set max jobs")
	flag.StringVar(&yara_path, "yara", "", "Set path of YARA file")

	flag.Usage = help_menu
	flag.Parse()

	switch flag.NFlag() {
	case 0:
		flag.Usage()
		os.Exit(0)

	case 2, 3:
		if dir_path == "" {
			err_flag = fmt.Errorf("No path provided\n%s\n", consts.FUUKA_USAGE)
			return "", 0, "", err_flag
		} else if yara_path == "" {
			err_flag = fmt.Errorf("No yara file provided\n%s\n", consts.FUUKA_USAGE)
			return "", 0, "", err_flag
		} else {
			fmt.Printf("Performing scan on %s against ruleset %s\nMax jobs: %d\n",
				dir_path, yara_path, max_jobs)
		}

	default:
		err_flag = fmt.Errorf("%d flags provided\n%s\n",
				flag.NFlag(), consts.FUUKA_USAGE)

		return "", 0, "", err_flag
	}

	return dir_path, max_jobs, yara_path, nil
}

/* Private */

func help_menu() {
	fmt.Printf("\033[1mFuuka %s\033[0m\n", consts.FUUKA_VERSION)
	fmt.Printf("Developed by patchmeifucan and haven7\n")
	fmt.Printf("\n%s\n", consts.FUUKA_USAGE)
	fmt.Printf("\nFlags:\n")
	fmt.Printf("\n    -h, --help - Print out the help menu\n")
	fmt.Printf("    -p, --path - Select path to scan, use . for pwd\n")
	fmt.Printf("    --yara - Select YARA ruleset file to load\n")
}
