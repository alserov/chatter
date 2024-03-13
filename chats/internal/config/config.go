package config

type Config struct {
	Env  string `json:"env" yaml:"env"`
	Port int    `json:"port" yaml:"port"`
}

func MustLoad() *Config {
	var cfg Config

	return &cfg
}
