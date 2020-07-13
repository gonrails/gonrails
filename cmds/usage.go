package cmds

import (
	"github.com/gonrails/gonrails/cmds/helper"
)

// Help - gonrails-cli help

var usageTemplate = `{{ "Gonrails" | bold }} is the command line tools for your Gonrails Web Application.

{{"USAGE" | bold | red}}

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
		'gonrails generate migration createUsers'
`

func Usage() {
	helper.Tmpl(usageTemplate, nil)
}

/*
 1. gonrails
 2. gonrails help
 3. gonrails help new
 4. gonrails help generate
*/
