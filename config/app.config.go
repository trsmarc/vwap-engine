package config

type AppConfig struct {
	MaxWindowSize int
}

var App = AppConfig{
	MaxWindowSize: 200,
}
