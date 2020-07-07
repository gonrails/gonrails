package new

import (
	"fmt"
	"os"

	"github.com/gonrails/gonrails/cmds/helper"
)

// TouchRoute -
/*
1. routes/base.go
2. routes/admin/base.go
*/
func TouchRoute(projectName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/routes/base.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/routes/base.go.template", helper.TemplatePath),
		info{
			ModuleName: projectName,
		},
	)

	_ = os.Mkdir(fmt.Sprintf("%s/routes/%s", projectName, "admin"), os.ModePerm)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/routes/admin/base.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/routes/admin/base.go.template", helper.TemplatePath),
		info{
			ModuleName: projectName,
		},
	)
}
