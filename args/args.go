/*
args.go - Entry point for handling command line arguments
Author: patchmeifucan
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
Handling any user flags provided, separate from main for sake of sanity
Returns a string and an error should flag.NFlag() be > 1
There is no reason for the user to use more than 1 flag with Fuuka
file_path is stored within main.go as a string to pass to Walker_Start()
It's much easier than just having this logic in the same directory as walker.go
*/
func Arg_Handler() (string, error) {
	var (
		err_flag 	error
		file_path string
	)

	/*
	Option of either -p or -path for setting the path
	Defaults to FUUKA_DEFAULT_PATH if no value is provided
	*/
	flag.StringVar(&file_path, "p",
									consts.FUUKA_DEFAULT_PATH,
									"Set path for Fuuka to scan")

	flag.StringVar(&file_path, "path",
									consts.FUUKA_DEFAULT_PATH,
									"Set path for Fuuka to scan")

	flag.Usage = help_menu
	flag.Parse()

	switch(flag.NFlag()) {
	case 0:
		help_menu()
		os.Exit(0)

	case 1:
		fmt.Printf("Performing scan on %s\n", file_path)

	default:
		err_flag = fmt.Errorf(
			"Flag amount of %d too large, maximum amount of flags should be 1\n",
			flag.NFlag())

		return "", err_flag
	}

	return file_path, nil
}

/* Private */

func help_menu() {
	fmt.Printf("\033[1mFuuka %s\033[0m\n",
							consts.FUUKA_VERSION)

	fmt.Printf("\nUsage: fuuka [args]\n")
	fmt.Printf("\nFlags:\n")
	fmt.Printf("\n    -h, --help - Print out the help menu\n")
	fmt.Printf("    -p, --path - Select path to scan, blank defaults to pwd\n")
}
