package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/one-hole/gonrails/cmds/helper"
)

var validActions = map[string]bool{
	"index":  true,
	"create": true,
	"delete": true,
	"update": true,
	"show":   true,
}

var basePath string

type ventory struct {
	ModuleName string
	ActionName string
}

func init() {
	if "true" == os.Getenv("DEV") {
		basePath = fmt.Sprintf("%s/watermelon", helper.Pwd)
	} else {
		basePath, _ = os.Getwd() // 非开发环境、那么这里应该是项目的根目录
	}
}

func Exec(args []string) {
	if len(args) == 2 {
		log.Fatal("Missing controller name")
	}

	fmt.Println(args[2])

	mkdir(args[2])
	touchActions(args)
}

func mkdir(path string) {
	var dirPath string
	if os.Getenv("DEV") == "true" {
		dirPath = fmt.Sprintf("%s%s/%s", helper.Pwd, "/watermelon/controllers", path)

		log.Println(dirPath)
	} else {
		tmpPath, _ := os.Getwd() // 这个命令我希望是在项目的目录里面执行
		dirPath = fmt.Sprintf("%s%s/%s", tmpPath, "/controllers", path)
	}

	err := helper.FindOrCreateDir(dirPath)
	helper.PanicError(err)
}

func touchActions(args []string) {
	actionArgs := args[3:]

	for _, action := range actionArgs {
		touchAction(action, args[2])
	}
}

func touchAction(action, path string) {
	if _, ok := validActions[action]; ok {
		helper.CreateFile(
			fmt.Sprintf("%s/controllers/%s/%s.go", basePath, path, action),
			fmt.Sprintf("%s/templates/controllers/action.go.template", helper.ProjectPath),
			ventory{
				ModuleName: path,
				ActionName: action,
			},
		)
	}
}
