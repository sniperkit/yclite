package config

import "flag"

// GlobalConfig -
type GlobalConfig struct {
	Port string // listen port
}

// DefaultGlobalConfig -
var DefaultGlobalConfig GlobalConfig

func init() {
	flag.StringVar(&DefaultGlobalConfig.Port, "port", ":8080", "listen port")
}
