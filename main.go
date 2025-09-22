/*
main.go - Entry point of Fuuka
Author: patchmeifucan
*/

package main

import (
	"fmt"

	"patchmeifucan/fuuka/args"
)

func main() {
	var (
		err_argerr 	error
		file_path 	string
	)

	file_path, err_argerr = args.Arg_Handler()
	if err_argerr != nil {
		panic(err_argerr)
	}

	fmt.Printf("File Path: %s\n", file_path)
}
