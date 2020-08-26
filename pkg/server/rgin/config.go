package rgin

import "time"

type RouterConfiguration struct {
	Mode           string
	EnableRecovery bool
}

type HttpConfiguration struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Configuration struct {
	routerConfiguration *RouterConfiguration
	httpConfiguration   *HttpConfiguration
}

func Config(httpConfig *HttpConfiguration, ginConfig *RouterConfiguration) *Configuration {
	return &Configuration{
		httpConfiguration:   httpConfig,
		routerConfiguration: ginConfig,
	}
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		httpConfiguration: &HttpConfiguration{
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
			Addr:         ":8080",
		},
		routerConfiguration: &RouterConfiguration{
			Mode:           "debug",
			EnableRecovery: true,
		},
	}
}
