package new

import (
	"fmt"
	"github.com/gonrails/gonrails/cmds/helper"
)

func TouchDockerfile(moduleName string) {

	helper.CreateFile(
		fmt.Sprintf("%s/%s/Dockerfile", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/Dockerfile.template", helper.ProjectPath),
		ventory{
			ModuleName:moduleName,
		},
		)
}