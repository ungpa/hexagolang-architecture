package mongodb

import (
	"context"
	"time"

	"github.com/ungpa/goon-ah-lang/internal/core/models"
	"github.com/ungpa/goon-ah-lang/internal/infrastructure/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlRepository struct {
	collection *mongo.Collection
}

func NewUrlRepository(client *mongo.Client, config *config.AppConfig) *UrlRepository {
	collection := client.Database("urlDB").Collection(config.Mongo.Collection)

	return &UrlRepository{collection: collection}
}

func (repo *UrlRepository) Create(data models.UrlAnalysisResult) error {
	doc := map[string]interface{}{
		"link":      data.Link,
		"metaTags":  data.MetaTags,
		"timestamp": time.Now(),
	}
	_, err := repo.collection.InsertOne(context.Background(), doc)

	return err
}
