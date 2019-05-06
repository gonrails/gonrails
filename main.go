package main

import (
	_ "kalista/config"
	"kalista/console"
	_ "kalista/utils/sources"
)

func main() {
	//routers.Start()
	console.Run()
}
