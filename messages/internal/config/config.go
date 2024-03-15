package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env  string `json:"env" yaml:"env"`
	Port int    `json:"port" yaml:"port"`
	DB   DB     `json:"db" yaml:"db"`
}

type DB struct {
	Scylla `yaml:"scylla"`
}

type Scylla struct {
	Addr string `yaml:"addr"`
}

func MustLoad() *Config {
	path := fetchPath()

	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}

	var cfg Config
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to parse fconfig file: " + err.Error())
	}

	return &cfg
}

func fetchPath() string {
	var path string
	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
