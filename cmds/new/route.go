package new

import (
	"fmt"
	"os"

	"github.com/one-hole/gonrails/cmds/helper"
)

// TouchRoute -
/*
1. routes/base.go
2. routes/admin/base.go
*/
func TouchRoute(moduleName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/routes/base.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/routes/base.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)

	_ = os.Mkdir(fmt.Sprintf("%s/routes/%s", moduleName, "admin"), os.ModePerm)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/routes/admin/base.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/routes/admin/base.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)
}
