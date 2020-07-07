package new

import (
	"fmt"
	"os"

	"github.com/gonrails/gonrails/cmds/helper"
)

// TouchController - generate controller example
func TouchController(projectName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/base.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/controllers/base.go.template", helper.TemplatePath),
		nil,
	)

	_ = os.Mkdir(fmt.Sprintf("%s/controllers/%s", projectName, "home"), os.ModePerm)
	_ = os.Mkdir(fmt.Sprintf("%s/controllers/%s", projectName, "books"), os.ModePerm)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/home/index.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/controllers/home/index.go.template", helper.TemplatePath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/books/index.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/controllers/books/index.go", helper.TemplatePath),
		nil,
	)

	helper.CreateFile(
		fmt.Sprintf("%s/%s/controllers/books/show.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/controllers/books/show.go.template", helper.TemplatePath),
		info{
			ModuleName: projectName,
		},
	)
}
