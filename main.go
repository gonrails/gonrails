package main

import (
	_ "github.com/one-hole/gonrails/config"
	"github.com/one-hole/gonrails/routers"
	_ "github.com/one-hole/gonrails/utils/sources"
)

func main() {
	routers.Start()
}
