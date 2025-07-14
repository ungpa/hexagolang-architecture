package core

import (
	"github.com/ungpa/goon-ah-lang/internal/core/ports"
	"github.com/ungpa/goon-ah-lang/internal/core/services"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(services.NewUrlService, fx.As(new(ports.UrlService))),
		fx.Annotate(services.NewScraper, fx.As(new(ports.ScraperService))),
	),
)
