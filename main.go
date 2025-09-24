/*
main.go - Entry point of Fuuka
Author: patchmeifucan
*/

package main

import (
	"fmt"
	"patchmeifucan/fuuka/args"
	"patchmeifucan/fuuka/walker"
	"time"
)

func main() {
	var (
		dir_path   string
		err_argerr error
		yara_path  string
	)

	dir_path, yara_path, err_argerr = args.Arg_Handler()
	if err_argerr != nil {
		panic(err_argerr)
	}
	time_now := time.Now()
	walker.Walker_Start(dir_path, yara_path)
	time_elapsed := time.Since(time_now).Milliseconds()

	fmt.Printf("Directory traversal took %d ms!\n", time_elapsed)
}
