package new

import (
	"fmt"
	"os"

	"github.com/one-hole/gonrails/cmds/helper"
)

// TouchController - generate controller example
func TouchController(moduleName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/base.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/base.go.template", helper.ProjectPath),
		nil,
	)

	_ = os.Mkdir(fmt.Sprintf("%s/controllers/%s", moduleName, "home"), os.ModePerm)
	_ = os.Mkdir(fmt.Sprintf("%s/controllers/%s", moduleName, "books"), os.ModePerm)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/home/index.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/home/index.go.template", helper.ProjectPath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/books/index.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/books/index.go", helper.ProjectPath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/books/show.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/books/show.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)
}
