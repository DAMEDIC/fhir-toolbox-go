package config

import (
	"cmp"
	"fhir-toolbox/rest"
	"os"
)

type serverConfig struct {
	rest.Config
	Port string `json:"port"`
}

var defaultServerConfig = serverConfig{
	Config: rest.DefaultConfig,
	Port:   cmp.Or(os.Getenv("SERVER_PORT"), "80"),
}
