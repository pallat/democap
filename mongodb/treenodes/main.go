package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	mongo2 = 2
	mongo3 = 3
)

func main() {
	primary()
	secondary(mongo2)
	secondary(mongo3)
}

func primary() {
	clientOpts := options.Client().ApplyURI(
		// "mongodb://mongo1:27017/?connect=direct")
		// "mongodb://localhost:27017/?connect=direct")
		"mongodb://mongo1:27017,mongo2:27017,mongo3:27017/?replicaSet=rs0")
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		slog.Error("ping:" + err.Error())
		return
	}

	collection := client.Database("db").Collection("coll")
	result, err := collection.InsertOne(ctx, bson.M{"x": 2})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("inserted ID: %v\n", result.InsertedID)
}

func secondary(num int) {
	uri := fmt.Sprintf("mongodb://mongo%d:27017/?connect=direct", num)

	clientOpts := options.Client().ApplyURI(
		uri)
	// "mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs0")
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		slog.Error("ping:" + err.Error())
		return
	}

	collection := client.Database("db").Collection("coll")

	var result bson.M
	if err := collection.FindOne(ctx, bson.M{"x": 1}).Decode(&result); err != nil {
		log.Panic(err)
	}
	fmt.Printf("inserted ID: %v\n", result)
}
