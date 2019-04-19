package main

import (
	"fmt"
	_ "kalista/config"
	"kalista/models"
	"kalista/utils/common"
	_ "kalista/utils/sources"
)

var (
	forever = make(chan bool)
)

func main() {
	//controllers.Start()
	fmt.Println(common.RandStr(63))

	models.NewTenant()
}
