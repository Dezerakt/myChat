package config

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig() *Config {
	return &Config{}
}
