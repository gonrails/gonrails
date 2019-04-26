package main

import (
	_ "kalista/config"
	"kalista/routers"
	_ "kalista/utils/sources"
)

func main() {
	routers.Start()
}
