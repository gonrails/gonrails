package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

type ventory struct {
	ModuleName string
}

// TouchMain - generate main.go
func TouchMain(projectName string) {

	helper.CreateFile(
		fmt.Sprintf("%s/%s/main.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/main.go.template", helper.TemplatePath),
		ventory{
			ModuleName: projectName,
		},
	)
}
