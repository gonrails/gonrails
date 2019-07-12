package new

import (
	"log"
	"os/exec"
)

// TODO .gitignore

/*InitGit -
1. .gitignore
2. git init
*/
func InitGit(moduleName string) {
	log.Printf("Excuting command: %s", "git init")
	ans := exec.Command("/bin/bash", "-c", "git init")
	_, err := ans.Output()
	if err != nil {
		log.Println(err)
	}
}
