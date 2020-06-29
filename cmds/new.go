package cmds

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) error {
	fmt.Println(ctx.Args())
	return nil
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
