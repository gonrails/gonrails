package new

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// InitMod - execute go mod commands
func InitMod(moduleName string) {
	os.Chdir(moduleName)
	pwd, _ := os.Getwd()
	log.Printf("Now in directory: %s", pwd)

	cmd := fmt.Sprintf("go mod init %s", moduleName)
	log.Printf("Excuting command: %s", cmd)
	ans := exec.Command("/bin/bash", "-c", cmd)
	_, err := ans.Output()
	if err != nil {
		log.Println(err)
	}

	cmd = "go mod tidy"
	log.Printf("Excuting command: %s", cmd)
	ans = exec.Command("/bin/bash", "-c", cmd)
	_, err = ans.Output()
	if err != nil {
		log.Println(err)
	}
}
