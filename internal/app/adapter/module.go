package adapter

import (
	"github.com/ungpa/goon-ah-lang/internal/app/adapter/handlers"
	"github.com/ungpa/goon-ah-lang/internal/app/adapter/repositories"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
