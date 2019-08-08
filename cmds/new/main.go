package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

type ventory struct {
	ModuleName string
}

// TouchMain - generate main.go
func TouchMain(moduleName string) {

	helper.CreateFile(
		fmt.Sprintf("%s/%s/main.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/main.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)
}
