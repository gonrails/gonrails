package main

import (
	"fmt"
	"os"

	"github.com/gonrails/gonrails/cmds"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "gonrails"
	app.Usage = "Gonrails 命令行工具"
	app.Commands = []*cli.Command{
		{
			Name:   "new",
			Usage:  "创建新项目 | Create a new project",
			Action: cmds.New,
		},
		{
			Name:  "version",
			Usage: "当前版本 | gonrails version",
			Action: func(c *cli.Context) error {
				fmt.Println(cmds.Version())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
