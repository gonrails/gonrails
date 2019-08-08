package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

// TouchConfig - generate config dir and files
func TouchConfig(moduleName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/config/config.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/config/config.go.template", helper.ProjectPath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/config/app.yml", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/config/app.yml", helper.ProjectPath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/config/config.yml", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/config/config.yml", helper.ProjectPath),
		nil,
	)
}
