package main

import (
	_ "kalista/config"
	"kalista/console"
	_ "kalista/utils/sources"
)

var (
	forever = make(chan bool)
)

func main() {
	console.Run()
	//routers.Start()

	<-forever
}
