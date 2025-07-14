package api

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHandler),
	fx.Provide(NewRouter),
	fx.Provide(SetupMiddleware),
	fx.Invoke(RegisterRoutes),
)
