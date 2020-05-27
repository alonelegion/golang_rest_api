package database

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
