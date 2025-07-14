package app

import (
	"github.com/ungpa/goon-ah-lang/internal/app/adapter"
	"github.com/ungpa/goon-ah-lang/internal/core"
	"github.com/ungpa/goon-ah-lang/internal/infrastructure/config"
	"github.com/ungpa/goon-ah-lang/internal/infrastructure/server"
	"go.uber.org/fx"
)

func NewApp(cfg *config.AppConfig) *fx.App {
	return fx.New(
		fx.Provide(func() *config.AppConfig {
			return cfg
		}),
		adapter.Module,
		core.Module,
		server.Module,
		fx.Invoke(
			server.StartServer,
		),
	)
}
