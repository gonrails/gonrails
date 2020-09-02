package cmds

import (
	"fmt"
	"log"
	"os"

	"github.com/gonrails/gonrails/cmds/helper"
	"github.com/gonrails/gonrails/cmds/new"
	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) error {
	name := ctx.Args().First()
	makeDirs(name)

	new.TouchMain(name)
	new.TouchConfig(name)

	return nil
}

func makeDirs(name string) {
	log.Printf(helper.Yellow("---> Generating gonrails project named %s ..."), helper.Green(name))
	_ = os.Mkdir(name, os.ModePerm)

	for _, subDir := range dirs {
		log.Println(fmt.Sprintf(helper.Blue("------> Making directory: %s"), fmt.Sprintf(helper.Green("%s/%s"), name, subDir)))
		_ = os.Mkdir(fmt.Sprintf("%s/%s", name, subDir), os.ModePerm)
	}
}

// New - Creates a new gonrails project
//func New(projectName string) {
//
//	log.Printf("Creating gonrails project named %s ...", projectName)
//	_ = os.Mkdir(projectName, os.ModePerm)
//
//	for _, name := range dirs {
//		log.Println(fmt.Sprintf("Making Dir: %s/%s", projectName, name))
//		_ = os.Mkdir(fmt.Sprintf("%s/%s", projectName, name), os.ModePerm)
//	}
//
//	new.TouchMain(projectName)
//	new.TouchConfig(projectName)
//	new.TouchRoute(projectName)
//	new.TouchController(projectName)
//	new.TouchSerializer(projectName)
//	new.TouchDockerfile(projectName)
//	new.InitMod(projectName)
//	new.InitGit(projectName)
//}
