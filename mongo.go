package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	mongoClient        *mongo.Client
	chatsCollection    *mongo.Collection
	messagesCollection *mongo.Collection
)

func InitMongo() {
	// mongodb connection string
	connectionString := "mongodb+srv://ant_joshua:joshuakeren123@cluster0.rjmmzo6.mongodb.net/?retryWrites=true&w=majority"

	// set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	mongoClient = client

	chatsCollection = mongoClient.Database("dating_go").Collection("chats")
	messagesCollection = mongoClient.Database("dating_go").Collection("chat_messages")

}
