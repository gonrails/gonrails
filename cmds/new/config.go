package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

// TouchConfig - generate config dir and files
func TouchConfig(projectName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/config/config.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/config/config.go.tmpl", helper.TemplatePath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/config/app.yml", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/config/app.yml", helper.TemplatePath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/config/config.yml", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/config/config.yml", helper.TemplatePath),
		nil,
	)
}
