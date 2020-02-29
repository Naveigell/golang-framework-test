package config

type ServerConfig struct {
    PORT string
    HOST string

    AcceptableRequestDomain[] string
}

var Server ServerConfig
