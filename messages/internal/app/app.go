package app

import (
	"context"
	"fmt"
	"github.com/alserov/chatter/messages/internal/config"
	"github.com/alserov/chatter/messages/internal/db"
	"github.com/alserov/chatter/messages/internal/db/scylla"
	"github.com/alserov/chatter/messages/internal/log"
	"github.com/alserov/chatter/messages/internal/server"
	"github.com/alserov/chatter/messages/internal/server/grpc"
	"github.com/alserov/chatter/messages/internal/usecase"
	"github.com/scylladb/gocqlx/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
)

func MustStart() {
	fx.New(
		fx.Provide(
			config.MustLoad,
			log.MustSetup,
		),

		fx.Invoke(func(cfg *config.Config, log log.Logger) {
			log.Debug("app config", zap.Any("cfg", cfg))
		}),

		// db conn
		fx.Provide(func(cfg *config.Config) gocqlx.Session {
			return scylla.MustConnect(cfg.DB.Keyspace, cfg.DB.Addr)
		}),

		// db repo init
		fx.Provide(func(conn gocqlx.Session) db.Repository {
			return scylla.NewRepository(conn)
		}),

		// usecase init
		fx.Provide(func(repo db.Repository) *usecase.Chat {
			return usecase.NewChat(usecase.Param{
				Repo: repo,
			})
		}),

		// server init
		fx.Provide(func(ucase *usecase.Chat) server.Server {
			return grpc.NewServer(ucase)
		}),

		fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, log log.Logger, srvr server.Server, dbConn gocqlx.Session) {
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
					dbConn.Close()
					srvr.Stop()
					log.Info("server was stopped")
					return nil
				},
			})
		}),
	).Run()
}
