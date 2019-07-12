package cmds

import (
	"log"

	"github.com/one-hole/gonrails-cli/cmds/generate/controller"
)

// Generate - for command generate
func Generate(args []string) {
	log.Println("Generate ...")

	if "controller" == args[1] {
		controller.Exec(args)
	}

	if "model" == args[1] {

	}
}
