package server

type ServiceInfo struct {
	Name string
}

type Server interface {
	Start() error
	Stop() error
	//Info() *ServiceInfo
}
