package migration

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gonrails/gonrails/cmds/helper"
	golbalhelper "github.com/gonrails/gonrails/helper"
)

/**
 Generate Migration Files

 1. Create Migration Dir if Not existed
 2. Create Migration File

 TODO MigrateName Camp
**/

var (
	basePath string
)

func init() {
	if "true" == os.Getenv("DEV") {
		basePath = fmt.Sprintf("%s/watermelon", helper.ProjectPath)
	} else {
		basePath, _ = os.Getwd()
	}
}

func Exec(args []string) {
	err := validateArgs(args)

	if err != nil {
		log.Fatal(err)
	}

	err = touchDir()

	if err != nil {
		log.Fatal(err)
	}

	migrateID := timeStr()
	migrateName := fmt.Sprintf("%s%s", args[2], "_migration")

	fileName := fmt.Sprintf("%s_%s.go", migrateID, args[2])
	filePath := fmt.Sprintf("%s/models/migrations/%s", basePath, fileName)

	type migrate struct {
		MigrateID   string
		MigrateName string
	}

	helper.CreateFile(
		filePath,
		fmt.Sprintf("%s/templates/models/migration.go.template", helper.ProjectPath),
		migrate{
			MigrateID:   migrateID,
			MigrateName: migrateName,
		},
	)
}

func validateArgs(args []string) error {
	if 3 != len(args) {
		return errors.New("Error: generate migration params length error")
	}
	return nil
}

func touchDir() error {
	migrationsDir := fmt.Sprintf("%s/models/migrations", basePath)
	return helper.FindOrCreateDir(migrationsDir)
}

func timeStr() string {
	return strings.Replace(time.Now().Format(golbalhelper.StampLayout), ".", "", -1)
}
