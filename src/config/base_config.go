package config

type Config struct{}

func New() *Config {
	return &Config{}
}

type Interface interface{}
