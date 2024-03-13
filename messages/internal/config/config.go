package config

type Config struct {
	Env  string `json:"env" yaml:"env"`
	Port int    `json:"port" yaml:"port"`
	DB   Scylla `json:"db" yaml:"db"`
}

type Scylla struct {
	Addr string
}

func MustLoad() *Config {
	var cfg Config

	return &cfg
}
