package helper

import (
	"fmt"
	"os"
)

var ProjectPath string  // 在哪里生成项目代码
var TemplatePath string // 模版的目录在哪里

func init() {
	if "true" == os.Getenv("DEV") {
		ProjectPath, _ = os.Getwd()
		TemplatePath = fmt.Sprintf("%s/templates", ProjectPath)
	} else {
		ProjectPath, _ = os.Getwd()
		TemplatePath = fmt.Sprintf("%s/src/github.com/gonrails/gonrails/templates", os.Getenv("GOPATH"))
	}

	fmt.Println(ProjectPath)
	fmt.Println(TemplatePath)
}
