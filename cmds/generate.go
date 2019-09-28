package cmds

import (
	"github.com/gonrails/gonrails/cmds/generate/controller"
	"github.com/gonrails/gonrails/cmds/generate/migration"
)

// Generate - for command generate
func Generate(args []string) {
	if "controller" == args[1] {
		controller.Exec(args)
	}

	if "migration" == args[1] {
		migration.Exec(args)
	}

	if "model" == args[1] {

	}
}
