package repository

import (
	"context"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConfig struct {
	ConnectionString string
	DatabaseName     string
}

func NewMongoClient(ctx context.Context, config MongoConfig) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed to establish mongo db connection: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to check (ping) mongo db connection: %s", err)
	}

	log.Print("connected to MongoDB")

	return client.Database(config.DatabaseName), nil
}
