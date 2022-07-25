package main

import (
	"os"

	"go-alfred/cmd_parse"
)

func main() {
	cmd := ""
	if len(os.Args) == 2 {
		cmd = os.Args[1]
	}
	cmd_parse.Execute(cmd)
}
