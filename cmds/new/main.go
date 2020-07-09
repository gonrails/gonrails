package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

type info struct {
	ModuleName string
}

func TouchMain(projectName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/main.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/main.go.tmpl", helper.TemplatePath),
		info{
			ModuleName: projectName,
		},
	)
}
