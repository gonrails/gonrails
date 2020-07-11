package rgin

type routerConfiguration struct {
}

type httpConfiguration struct {
}

type Configuration struct {
	routerConfiguration *routerConfiguration
	httpConfiguration   *httpConfiguration
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		routerConfiguration: &routerConfiguration{},
		httpConfiguration:   &httpConfiguration{},
	}
}
