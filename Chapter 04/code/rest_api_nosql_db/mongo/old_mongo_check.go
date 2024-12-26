package main

import (
	"context"
	"fmt"
	"log"
	"time"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	mongoURI := "mongodb+srv://Cluster72877:RX1mZlt9bUZB@cluster72877.rx6yjqr.mongodb.net/?ssl=true&retryWrites=true"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("go")
	collection := database.Collection("atlas")

	fmt.Println("database", database)

	fmt.Println("collection", collection)
}
