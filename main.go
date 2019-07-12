package main

import (
	"flag"
	"os"

	"github.com/one-hole/gonrails/cmds"
)

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		cmds.Useage()
		os.Exit(2)
	}

	if "new" == args[0] {
		cmds.New(args[1])
	}

	if "help" == args[0] {
		cmds.Useage()
		return
	}

	if "generate" == args[0] {
		cmds.Generate(args)
	}
}
