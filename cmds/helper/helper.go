package helper

import (
	"fmt"
	"os"
)

// ProjectPath for template files
var ProjectPath string

// Pwd for current dir
var Pwd, _ = os.Getwd()

func init() {
	if "true" == os.Getenv("DEV") {
		ProjectPath, _ = os.Getwd()
	} else {
		ProjectPath = fmt.Sprintf("%s/src/github.com/one-hole/gonrails-cli", os.Getenv("GOPATH"))
	}
}
