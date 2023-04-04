package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://durgareddyalt:aYuvqzOuHqwF2Rg6@messages01.bw6knds.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println("Failed to establish the client!! with Error:", err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Println("Failed to connect to database:", err.Error())
	}

	defer client.Disconnect(ctx)

	ChatBoxDatabase := client.Database("chat_box")
	UsersCollection := ChatBoxDatabase.Collection("users")

	userResult, err := UsersCollection.InsertOne(ctx, bson.D{{Key: "UserId", Value: "DurgareddyAlt"}, {Key: "Password", Value: "aYuvqzOuHqwF2Rg6"}})
	if err != nil {
		log.Println("UsersCollection.InsertOne failed with an error:", err.Error())
	}
	log.Println("--->", userResult)
}
