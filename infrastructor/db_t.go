package infrastructor

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func SetupTestDatabase(t *testing.T) (*mongo.Client, *mongo.Database) {
	clientOpts := options.Client().ApplyURI("mongodb+srv://andrew2611:11062001Phong@andrew.8ulkv.mongodb.net/?retryWrites=true&w=majority&appName=Andrew")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	database := client.Database("test_db")
	return client, database
}

func TearDownTestDatabase(client *mongo.Client, t *testing.T) {
	err := client.Disconnect(context.Background())
	if err != nil {
		t.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}
