package repositories

import (
	"github.com/ungpa/goon-ah-lang/internal/app/adapter/repositories/mongodb"
	"go.uber.org/fx"
)

var Module = fx.Options(
	mongodb.Module,
)
