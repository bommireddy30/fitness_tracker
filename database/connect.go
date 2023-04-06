package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Ctx context.Context

	UsersCollection    *mongo.Collection
	UserCollectionName string = "users"
)

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://durgareddyalt:Abcd1234@messages01.bw6knds.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println("Failed to establish the client!! with Error:", err.Error())
	}

	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(Ctx)
	if err != nil {
		log.Println("Failed to connect to database:", err.Error())
	}

	err = client.Ping(Ctx, readpref.Primary())
	if err != nil {
		log.Println("Ping Failed with error:", err.Error())
	}

	FitnessTracker := client.Database("fitness_tracker")
	UsersCollection = FitnessTracker.Collection(UserCollectionName)

	log.Println("Database is connected")
}
