package config

type ServerConfig struct {
    PORT string
    HOST string

    AcceptRequestDomain[] string
}

var Server ServerConfig
