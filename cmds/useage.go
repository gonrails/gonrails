package cmds

import (
	"github.com/one-hole/gonrails/cmds/helper"
)

// Help - gonrails-cli help

var useageTemplate = `{{ "Gonrails" | bold }} is the command line tools for your Gonrails Web Application.

{{"USEAGE" | bold | red}}

	gonrails command [arguments]

{{"COMMANDS" | bold | red}}

	new       Create a new project
	help      Show the usage of commands
	generate  Generate code by commands

{{"EXAMPLES" | bold | red}}

	new

		'gonrails new kalista'

	generate

		'gonrails generate model user'
		'gonrails generate controller users index show update delete create'
`

func Useage() {
	helper.Tmpl(useageTemplate, nil)
}

/*
 1. gonrails
 2. gonrails help
 3. gonrails help new
 4. gonrails help generate
*/
