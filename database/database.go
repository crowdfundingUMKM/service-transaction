package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnectionDB() *mongo.Database {
	// Set your MongoDB connection parameters
	var (
		dbUser = os.Getenv("DB_USER") // e.g. 'my-db-user'
		dbPwd  = os.Getenv("DB_PASS") // e.g. 'my-db-password'
		dbName = os.Getenv("DB_NAME") // e.g. 'my-database'
		dbHost = os.Getenv("DB_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
		dbPort = os.Getenv("DB_PORT") // e.g. '3306'
	)

	// Create a MongoDB connection string
	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPwd, dbHost, dbPort)

	// Create a MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Access the database
	db := client.Database(dbName)

	return db
}
