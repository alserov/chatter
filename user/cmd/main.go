package main

import (
	"github.com/MaksKazantsev/chatter/user/app"
	"github.com/MaksKazantsev/chatter/user/internal/config"
)

func main() {
	cfg := config.ReadConfig()
	app.MustStart(cfg)
}
