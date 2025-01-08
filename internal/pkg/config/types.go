package config

type Config struct {
	Server Server
}

type Server struct {
	Port     int    `env:"PORT" envDefault:"3000"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}
