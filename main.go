/*
main.go - Entry point of Fuuka
Author: patchmeifucan
*/

package main

import (
	"patchmeifucan/fuuka/args"
	"patchmeifucan/fuuka/walker"
)

func main() {
	var (
		dir_path 		string
		err_argerr 	error
		yara_path 	string
	)

	dir_path, yara_path, err_argerr = args.Arg_Handler()
	if err_argerr != nil {
		panic(err_argerr)
	}

	walker.Walker_Start(dir_path, yara_path)
}
