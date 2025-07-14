package handlers

import (
	"github.com/ungpa/goon-ah-lang/internal/app/adapter/handlers/api"
	"github.com/ungpa/goon-ah-lang/internal/app/adapter/handlers/kafka"
	"go.uber.org/fx"
)

var Module = fx.Options(
	kafka.Module,
	api.Module,
)
