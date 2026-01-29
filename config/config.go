package config

type Config struct {
	Server struct {
		Port string
	}

	Auth struct {
		Issuer string
		Secret string
	}
}

func NewConfig() *Config {
	return &Config{}
}
