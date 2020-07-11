package rgin

import "time"

type routerConfiguration struct {
	mode           string
	enableRecovery bool
}

type httpConfiguration struct {
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Configuration struct {
	routerConfiguration *routerConfiguration
	httpConfiguration   *httpConfiguration
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		httpConfiguration: &httpConfiguration{
			readTimeout:  10 * time.Second,
			writeTimeout: 10 * time.Second,
			addr:         ":8080",
		},
		routerConfiguration: &routerConfiguration{
			mode:           "debug",
			enableRecovery: true,
		},
	}
}
