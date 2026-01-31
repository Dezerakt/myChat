package config

type (
	Config struct {
		Server   Server
		Otp      Otp
		Postgres Postgres
	}

	Postgres struct {
		Port     string
		Host     string
		User     string
		Password string
		Database string
	}

	Server struct {
		Port string
	}

	Otp struct {
		Issuer string
		Secret string
		Period int
	}
)

func NewConfig() *Config {
	return &Config{}
}
