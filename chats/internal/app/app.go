package app

import (
	"context"
	"fmt"
	"github.com/alserov/chatter/chats/internal/config"
	"github.com/alserov/chatter/chats/internal/log"
	"github.com/alserov/chatter/chats/internal/repository/scylla"
	"github.com/alserov/chatter/chats/internal/usecase"
	"go.uber.org/fx"
	"net"
)

func MustStart() {
	fx.New(
		fx.Provide(
			config.MustLoad,
			log.MustSetup,
			scylla.NewRepository,
			usecase.NewChat,
		),

		fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, ctrl) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
					if err != nil {
						panic("failed to listen: " + err.Error())
					}

				},
			})
		}),
	).Run()
}
