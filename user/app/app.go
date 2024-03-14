package app

import (
	"github.com/MaksKazantsev/chatter/user/internal/config"
	"github.com/MaksKazantsev/chatter/user/internal/log"
)

func MustStart(config *config.Config) {
	l := log.GetLogger(config.Env)
	l.Info("To be continued")
}
