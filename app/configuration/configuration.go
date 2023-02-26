package configuration

import "test/infra/web"

type ConfigurationWebServer struct {
	APIsPathCommon string
	Port           string `valid:"required"`
}

type Configuration struct {
	ConfigurationWebServer

	Mode int
}

func (c *Configuration) AsWebServerConfiguration() *web.ConfigurationWebServer {
	return &web.ConfigurationWebServer{
		APIsPathCommon: c.APIsPathCommon,
		Port:           c.Port,
	}
}
