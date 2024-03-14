package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env  string   `yaml:"env" env:"ENV" env-default:"local"`
	Port int      `yaml:"port" env:"PORT" env-default:"9991"`
	DB   Database `yaml:"db"`
}

type Database struct {
	Host string `yaml:"host" env:"DB_HOST"`
	Port int    `yaml:"port" env:"DB_PORT"`
}

func (d Database) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%d", d.Host, d.Port)
}

func ReadConfig() *Config {
	path := fetchPath()

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config file:" + err.Error())
	}
	return &cfg
}

func fetchPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
