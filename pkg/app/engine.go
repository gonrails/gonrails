package app

import "gonrails/pkg/server"

type Engine struct {
	servers []server.Server
}

func (engine *Engine) initialize() {

}

func (engine *Engine) startServers() {
	for _, server := range engine.servers {
		server.Start()
	}
}

func (engine *Engine) clean() {}

func (engine *Engine) Run() {
	defer engine.clean()
	engine.startServers()
}
