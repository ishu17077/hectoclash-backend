package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	mongoDb := "mongodb://localhost:27017"
	fmt.Print("Starting mongo db service if stopped")
	// cmd := exec.Command("sudo", "systemctl", "start", "mongod.service")
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error starting mongod service: %v\n", err)
	// }
	fmt.Print(mongoDb)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDb))

	if err != nil {
		msg := fmt.Sprintf("Connection to MongoDB has failed :%s", err)
		log.Fatal(msg);
	}

	fmt.Print("Connected to MongoDB Instance")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("hectoclash").Collection(collectionName)
	return collection
}
