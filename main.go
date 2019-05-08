package main

import (
	_ "kalista/config"
	"kalista/console"
	_ "kalista/utils/sources"
)

func main() {
	console.Run()
	//routers.Start()
}
