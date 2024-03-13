package app

import (
	"context"
	"fmt"
	"github.com/alserov/chatter/messages/internal/config"
	"github.com/alserov/chatter/messages/internal/log"
	"github.com/alserov/chatter/messages/internal/repository/scylla"
	"github.com/alserov/chatter/messages/internal/server"
	"github.com/alserov/chatter/messages/internal/usecase"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
)

func MustStart() {
	fx.New(
		fx.Provide(
			config.MustLoad,
			log.MustSetup,
			scylla.NewRepository,
			usecase.NewChat,
			server.NewServer,
		),

		fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, log log.Logger, srvr server.Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Info("starting server")

					l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
					if err != nil {
						panic("failed to listen: " + err.Error())
					}

					go func() {
						if err = srvr.Serve(l); err != nil {
							panic("failed to start server")
						}
					}()

					log.Info("server is running", zap.Int("port", cfg.Port))
					return nil
				},
				OnStop: func(ctx context.Context) error {
					srvr.Stop()
					log.Info("server was stopped")
					return nil
				},
			})
		}),
	).Run()
}
