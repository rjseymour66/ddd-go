package purchase

import (
	"context"
	"fmt"
)

// Repository provides flexibility about which data store we use.
type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{
		purchases: purchases,
	}, nil
}
