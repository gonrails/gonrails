package main

import (
	_ "kalista/config"
	"kalista/controllers"
	_ "kalista/utils/sources"
)

var (
	forever = make(chan bool)
)

func main() {
	controllers.Start()
}
