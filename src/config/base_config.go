package config

type Interface interface {
	GetBBLUrl() string
}

type Config struct{}

func (c *Config) GetBBLUrl() string {
	return "https://psipay.bangkokbank.com/b2c/eng/merchant/go-api/"
}

func New() Interface {
	return &Config{}
}
