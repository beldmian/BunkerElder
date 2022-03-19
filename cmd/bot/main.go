// Package main provides main application executable
package main

import (
	"context"

	"github.com/beldmian/bunkerelder/pkg/config"
	"github.com/beldmian/bunkerelder/pkg/logger"
	"github.com/beldmian/bunkerelder/pkg/tg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func registerHooks(lc fx.Lifecycle, l *zap.Logger, conf *config.Config, tg *tg.TgBot) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go tg.Start()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}

func main() {
	app := fx.New(
		fx.Provide(
			config.ProvideConfig,
			logger.ProvideLogger,
			tg.ProvideTg,
		),
		fx.Invoke(registerHooks),
		fx.WithLogger(func(l *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: l,
			}
		}),
	)
	app.Run()
}
