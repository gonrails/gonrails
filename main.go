package main

import (
	_ "kalista/config"
	"kalista/routers"
	_ "kalista/utils/sources"
)

var (
	forever = make(chan bool)
)

func main() {
	routers.Start()
}
