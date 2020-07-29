package app

import "github.com/gonrails/gonrails/pkg/server"

type Engine struct {
	servers []server.Server
}

func DefaultEngine() *Engine {
	engine := &Engine{}
	engine.initialize()
	return engine
}

func (engine *Engine) startServers() {
	for _, server := range engine.servers {
		server.Start()
	}
}

func (engine *Engine) clean() {
	for _, server := range engine.servers {
		server.Stop()
	}
}

func (engine *Engine) initialize() {
	engine.servers = make([]server.Server, 0)
}

func (engine *Engine) AddServer(s server.Server) {
	engine.servers = append(engine.servers, s)
}

func (engine *Engine) Serve() {
	defer engine.clean()
	engine.startServers()
}
