package mongodb

import (
	"github.com/ungpa/goon-ah-lang/internal/core/ports"
	"github.com/ungpa/goon-ah-lang/internal/infrastructure/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(config *config.AppConfig) MongoDBConfig {
			conn := config.Mongo.ConnectionString
			if conn == "" && config.Environment == "Development" { // for local development
				conn = "mongodb://localhost:27017"
			}

			return MongoDBConfig{URI: conn}
		},
		func(cfg MongoDBConfig) (*mongo.Client, error) {
			return ConnectToMongo(cfg)
		},
	),
	fx.Provide(
		fx.Annotate(NewUrlRepository, fx.As(new(ports.UrlRepository))),
	),
)
