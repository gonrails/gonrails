package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

func TouchDockerfile(projectName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/Dockerfile", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/Dockerfile.template", helper.TemplatePath),
		info{
			ModuleName: projectName,
		},
	)
}
