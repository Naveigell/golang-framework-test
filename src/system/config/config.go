package config

type ServerConfig struct {
    PORT string
    HOST string

    AcceptableHostRequest[] string
}

type CorsConfig struct {
    GlobalCorsEnable bool
}

var (
    Cors CorsConfig
    Server ServerConfig
)
