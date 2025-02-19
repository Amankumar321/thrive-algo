package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URI from the environment
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is not set in the environment")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("holiday_calendar")
	// Ensure the collection exists
	ensureCollectionExists(ctx, DB, "holidays")
}

func ensureCollectionExists(ctx context.Context, db *mongo.Database, collectionName string) {
	collections, err := db.ListCollectionNames(ctx, bson.M{"name": collectionName})
	if err != nil {
		log.Fatal("Error listing collections: ", err)
	}

	if len(collections) == 0 {
		err := db.CreateCollection(ctx, collectionName)
		if err != nil {
			log.Fatal("Failed to create collection: ", err)
		}
	}
}
